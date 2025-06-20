package forth

import (
	"errors"
	"strconv"
	"strings"
)

// Forth evaluates a series of Forth expressions and returns the resulting stack
func Forth(input []string) ([]int, error) {
	interpreter := newInterpreter()

	for _, program := range input {
		err := interpreter.eval(program)
		if err != nil {
			return nil, err
		}
	}

	return interpreter.stack, nil
}

type word struct {
	tokens  []string
	context map[string]word // Store a snapshot of the definitions at the time this word was defined
}

type interpreter struct {
	stack       []int
	definitions map[string]word
}

func newInterpreter() *interpreter {
	return &interpreter{
		stack:       []int{},
		definitions: make(map[string]word),
	}
}

func (i *interpreter) eval(program string) error {
	program = strings.ToLower(program)

	// Check if this is a definition
	if strings.HasPrefix(program, ":") && strings.HasSuffix(program, ";") {
		return i.parseDefinition(program)
	}

	tokens := strings.Fields(program)
	return i.evalTokens(tokens, i.definitions)
}

func (i *interpreter) parseDefinition(program string) error {
	// Remove ":" and ";" from the definition
	definition := strings.TrimPrefix(program, ":")
	definition = strings.TrimSuffix(definition, ";")
	tokens := strings.Fields(definition)

	if len(tokens) < 2 {
		return errors.New("illegal operation")
	}

	name := tokens[0]
	bodyTokens := tokens[1:]

	// Check if name is a number
	if _, err := strconv.Atoi(name); err == nil {
		return errors.New("illegal operation")
	}

	// Create a snapshot of the current definitions for this word's context
	i.definitions[name] = word{
		tokens:  bodyTokens,
		context: copyDefinitions(i.definitions),
	}

	return nil
}

func (i *interpreter) evalTokens(tokens []string, context map[string]word) error {
	for j := 0; j < len(tokens); j++ {
		token := tokens[j]

		// Try to parse as number
		num, err := strconv.Atoi(token)
		if err == nil {
			i.stack = append(i.stack, num)
			continue
		}

		// Check if token is a user-defined word
		if w, exists := context[token]; exists {
			// Execute the definition using its own context when it was defined
			if err := i.evalTokens(w.tokens, w.context); err != nil {
				return err
			}
			continue
		}

		// Handle built-in operations
		switch token {
		case "+":
			if len(i.stack) < 2 {
				if len(i.stack) == 0 {
					return errors.New("empty stack")
				}
				return errors.New("only one value on the stack")
			}
			a, b := i.pop(), i.pop()
			i.stack = append(i.stack, b+a)
		case "-":
			if len(i.stack) < 2 {
				if len(i.stack) == 0 {
					return errors.New("empty stack")
				}
				return errors.New("only one value on the stack")
			}
			a, b := i.pop(), i.pop()
			i.stack = append(i.stack, b-a)
		case "*":
			if len(i.stack) < 2 {
				if len(i.stack) == 0 {
					return errors.New("empty stack")
				}
				return errors.New("only one value on the stack")
			}
			a, b := i.pop(), i.pop()
			i.stack = append(i.stack, b*a)
		case "/":
			if len(i.stack) < 2 {
				if len(i.stack) == 0 {
					return errors.New("empty stack")
				}
				return errors.New("only one value on the stack")
			}
			a, b := i.pop(), i.pop()
			if a == 0 {
				return errors.New("divide by zero")
			}
			i.stack = append(i.stack, b/a)
		case "dup":
			if len(i.stack) < 1 {
				return errors.New("empty stack")
			}
			i.stack = append(i.stack, i.stack[len(i.stack)-1])
		case "drop":
			if len(i.stack) < 1 {
				return errors.New("empty stack")
			}
			i.stack = i.stack[:len(i.stack)-1]
		case "swap":
			if len(i.stack) < 2 {
				if len(i.stack) == 0 {
					return errors.New("empty stack")
				}
				return errors.New("only one value on the stack")
			}
			i.stack[len(i.stack)-1], i.stack[len(i.stack)-2] = i.stack[len(i.stack)-2], i.stack[len(i.stack)-1]
		case "over":
			if len(i.stack) < 2 {
				if len(i.stack) == 0 {
					return errors.New("empty stack")
				}
				return errors.New("only one value on the stack")
			}
			i.stack = append(i.stack, i.stack[len(i.stack)-2])
		default:
			return errors.New("undefined operation")
		}
	}

	return nil
}

func (i *interpreter) pop() int {
	val := i.stack[len(i.stack)-1]
	i.stack = i.stack[:len(i.stack)-1]
	return val
}

// Helper function to create a copy of the definitions map
func copyDefinitions(defs map[string]word) map[string]word {
	copy := make(map[string]word)
	for name, w := range defs {
		copy[name] = w
	}
	return copy
}
