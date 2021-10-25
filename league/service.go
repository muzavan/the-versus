package league

import (
	"context"
	"fmt"
	"strings"
)

type Service struct {
	compStore   CompetitionStore
	reportStore ReportStore
}

func NewService(compStore CompetitionStore, reportStore ReportStore) (*Service, error) {
	return &Service{
		compStore:   compStore,
		reportStore: reportStore,
	}, nil
}

var (
	ErrInvalidParams = fmt.Errorf("invalid params")
	ErrInternal      = fmt.Errorf("internal error")
)

func (r *Service) Report(ctx context.Context, ID string, matches []*Match) error {
	competition, err := r.validate(ctx, ID, matches)
	if err != nil {
		return err
	}

	return r.reportStore.Save(ctx, ID, competition.ID, matches)

}

func (r *Service) validate(ctx context.Context, ID string, matches []*Match) (*Competition, error) {
	// Valid if:
	// Has ID
	// All competitor belongs to the same Competition
	// All competitor has the same number of match participated

	if len(strings.TrimSpace(ID)) == 0 {
		return nil, fmt.Errorf("did not have any ID")
	}

	competitionID := ""
	participationMap := map[string]int{} //Map[Competitor.Name] -> No. Occurence

	for _, m := range matches {
		if m.A.CompetitionID != m.B.CompetitionID {
			return nil, fmt.Errorf("%w - invalid report: multiple competition", ErrInvalidParams)
		}

		if len(m.A.CompetitionID) == 0 {
			return nil, fmt.Errorf("%w - invalid report: no competition", ErrInvalidParams)
		}

		if competitionID == "" {
			competitionID = m.A.CompetitionID
		} else if competitionID != m.A.CompetitionID {
			return nil, fmt.Errorf("%w - invalid report: multiple competition", ErrInvalidParams)
		}

		if m.A.Name == m.B.Name {
			return nil, fmt.Errorf("%w - invalid report: same participant", ErrInvalidParams)
		}

		participationMap[m.A.Name]++
		participationMap[m.B.Name]++
	}

	expCount := 0
	for _, count := range participationMap {
		if expCount == 0 {
			expCount = count
		} else if expCount != count {
			return nil, fmt.Errorf("%w - invalid report: unbalanced participant", ErrInvalidParams)
		}
	}

	competition, err := r.compStore.LoadByID(ctx, competitionID)
	if err != nil {
		return nil, fmt.Errorf("%w - loadCompetitionByID fails: %s", ErrInternal, err.Error())
	}

	for _, cp := range competition.Competitors {
		if _, ok := participationMap[cp.Name]; !ok {
			return nil, fmt.Errorf("%w - invalid report: missing participant(s)", ErrInvalidParams)
		}
		delete(participationMap, cp.Name)
	}

	if len(participationMap) > 0 {
		return nil, fmt.Errorf("%w - invalid report: unrecognized participant", ErrInvalidParams)
	}

	return competition, nil
}
