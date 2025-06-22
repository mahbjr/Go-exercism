package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	// Caso especial: nenhum registro fornecido
	if len(records) == 0 {
		return nil, nil
	}

	// Verificar registros duplicados e criar um mapa de ID para Record
	recordMap := make(map[int]Record)
	for _, r := range records {
		// Verificar se o ID já existe (duplicado)
		if _, exists := recordMap[r.ID]; exists {
			return nil, errors.New("duplicate node")
		}
		recordMap[r.ID] = r
	}

	// Verificar se o nó raiz (ID 0) existe
	if _, exists := recordMap[0]; !exists {
		return nil, errors.New("no root node")
	}

	// Verificar se o nó raiz tem um pai
	if recordMap[0].Parent != 0 {
		return nil, errors.New("root node has parent")
	}

	// Verificar se todos os IDs são contínuos de 0 até N-1
	if len(recordMap) != len(records) {
		return nil, errors.New("duplicate node")
	}
	for i := 0; i < len(records); i++ {
		if _, exists := recordMap[i]; !exists {
			return nil, errors.New("non-continuous")
		}
	}

	// Criar um mapa de nós
	nodes := make(map[int]*Node)
	for id := range recordMap {
		nodes[id] = &Node{ID: id}
	}

	// Construir a árvore adicionando filhos aos seus pais
	for id, record := range recordMap {
		// Pular o nó raiz para verificações de ciclo
		if id == 0 {
			continue
		}

		// Verificar ciclo direto
		if id == record.Parent {
			return nil, errors.New("cycle directly")
		}

		// Verificar se o ID do pai existe
		if _, exists := recordMap[record.Parent]; !exists {
			return nil, errors.New("parent does not exist")
		}

		// Verificar se o filho tem ID maior que o pai
		if id < record.Parent {
			return nil, errors.New("higher id parent of lower id")
		}

		// Verificar ciclo indireto
		current := record.Parent
		for current != 0 {
			if current == id {
				return nil, errors.New("cycle indirectly")
			}
			current = recordMap[current].Parent
		}

		// Adicionar o nó atual como filho do seu pai
		parentNode := nodes[record.Parent]
		parentNode.Children = append(parentNode.Children, nodes[id])
	}

	// Ordenar os filhos de cada nó por ID
	for _, node := range nodes {
		if len(node.Children) > 0 {
			sort.Slice(node.Children, func(i, j int) bool {
				return node.Children[i].ID < node.Children[j].ID
			})
		}
	}

	return nodes[0], nil
}
