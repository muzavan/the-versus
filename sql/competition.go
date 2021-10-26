package sql

import (
	"context"
	"strings"

	"github.com/muzavan/the-versus/league"
)

func generateCompetitor(competitionID string, displayName string) league.Competitor {
	return league.Competitor{
		Name:          strings.ToLower(displayName),
		DisplayName:   displayName,
		CompetitionID: competitionID,
	}
}

func (st *Store) LoadByID(ctx context.Context, competitionID string) (*league.Competition, error) {
	return &league.Competition{
		ID:          competitionID,
		Name:        "The Competition",
		Description: "A sample competition",
		Competitors: []league.Competitor{
			generateCompetitor(competitionID, "Team A"),
			generateCompetitor(competitionID, "Team B"),
			generateCompetitor(competitionID, "Team C"),
			generateCompetitor(competitionID, "Team D"),
			generateCompetitor(competitionID, "Team E"),
		},
	}, nil
}
