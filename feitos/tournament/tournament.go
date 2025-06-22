package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Team holds the stats for a single team
type Team struct {
	Name string
	MP   int // Matches Played
	W    int // Wins
	D    int // Draws
	L    int // Losses
	P    int // Points
}

// Tally reads match results from a reader and writes the tournament table to a writer
func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	teams := make(map[string]*Team)

	// Process each line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comment lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse the line
		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return errors.New("invalid format: each line must have exactly 3 parts separated by semicolons")
		}

		team1 := parts[0]
		team2 := parts[1]
		result := parts[2]

		// Validate the result
		if result != "win" && result != "loss" && result != "draw" {
			return errors.New("invalid result: must be 'win', 'loss', or 'draw'")
		}

		// Validate team names (no special characters)
		if strings.ContainsAny(team1, "_@") || strings.ContainsAny(team2, "_@") {
			return errors.New("invalid team name: must not contain special characters")
		}

		// Create teams if they don't exist yet
		if _, exists := teams[team1]; !exists {
			teams[team1] = &Team{Name: team1}
		}
		if _, exists := teams[team2]; !exists {
			teams[team2] = &Team{Name: team2}
		}

		// Update team stats based on the result
		teams[team1].MP++
		teams[team2].MP++

		switch result {
		case "win":
			teams[team1].W++
			teams[team1].P += 3
			teams[team2].L++
		case "loss":
			teams[team1].L++
			teams[team2].W++
			teams[team2].P += 3
		case "draw":
			teams[team1].D++
			teams[team1].P++
			teams[team2].D++
			teams[team2].P++
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Sort teams by points (descending) and then by name (ascending)
	sortedTeams := make([]*Team, 0, len(teams))
	for _, team := range teams {
		sortedTeams = append(sortedTeams, team)
	}
	sort.Slice(sortedTeams, func(i, j int) bool {
		if sortedTeams[i].P != sortedTeams[j].P {
			return sortedTeams[i].P > sortedTeams[j].P
		}
		return sortedTeams[i].Name < sortedTeams[j].Name
	})

	// Write the header
	fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")

	// Write team stats
	for _, team := range sortedTeams {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			team.Name, team.MP, team.W, team.D, team.L, team.P)
	}

	return nil
}
