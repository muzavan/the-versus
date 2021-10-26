package league

import (
	"fmt"
	"text/template"

	"github.com/rs/zerolog"
)

// Server, flatten the controller and service layer
// Since the controller logic is quite thin
type Server struct {
	compStore   CompetitionStore
	reportStore ReportStore
	logger      zerolog.Logger

	matchesTemplate *template.Template
}

func NewServer(compStore CompetitionStore, reportStore ReportStore, baseLogger zerolog.Logger) (*Server, error) {
	return &Server{
		compStore:   compStore,
		reportStore: reportStore,
		logger:      baseLogger.With().Str("component", "league-server").Logger(),
	}, nil
}

var (
	ErrInvalidParams = fmt.Errorf("invalid params")
	ErrInternal      = fmt.Errorf("internal error")
)
