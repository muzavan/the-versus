package league

type Competitor struct {
	// Name is the ID of Competitor, so should be unique
	Name string `json:"name,omitempty"`
	// DisplayName is the name that shown to any FE
	// Name is basically DisplayName.toLowerCase()
	DisplayName   string            `json:"display_name,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
	CompetitionID string            `json:"competition_id,omitempty"`
}

type Competition struct {
	ID          string            `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	Competitors []Competitor      `json:"competitors,omitempty"`
}

// GetMatches return all matches
// which are all combinations of Competitor
func (c *Competition) Matches() []*Match {
	matches := []*Match{}

	allCompetitors := c.Competitors
	for ai := 0; ai < len(allCompetitors)-1; ai++ {
		for bi := ai + 1; bi < len(allCompetitors); bi++ {
			cpA := allCompetitors[ai]
			cpB := allCompetitors[bi]

			if cpA.Name > cpB.Name {
				cpA, cpB = cpB, cpA
			}

			matches = append(matches, &Match{
				A:      &cpA,
				B:      &cpB,
				Result: MatchResultDraw,
			})
		}
	}

	return matches
}

type MatchResult uint32

const (
	MatchResultDraw = MatchResult(0)
	MatchResultAWin = MatchResult(1)
	MatchResultBWin = MatchResult(2)
)

// Match is a representation of A vs B
// To simplify, A.Name always <= B.Name
type Match struct {
	A      *Competitor `json:"a,omitempty"`
	B      *Competitor `json:"b,omitempty"`
	Result MatchResult `json:"result"`
}
