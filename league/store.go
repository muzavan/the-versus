package league

import "context"

type CompetitionStore interface {
	LoadByID(ctx context.Context, competitionID string) (*Competition, error)
}

type ReportStore interface {
	Save(ctx context.Context, reportID string, competitionID string, matches []*Match) error
}
