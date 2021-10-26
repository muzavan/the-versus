package league

import (
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func (s *Server) Matches(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	competitionID := params["id"]

	ctx := req.Context()
	pageParams := MatchesPageParam{}
	competition, err := s.compStore.LoadByID(ctx, competitionID)
	if err != nil {
		s.logger.Error().Err(err).Msg("compSore.LoadByID fails")
		pageParams.ErrMessage = "Oops! We could not server this page at the moment."
	}
	pageParams.Competition = competition
	pageParams.Matches = competition.Matches()
	err = s.renderMatchesPage(w, pageParams)
	if err != nil {
		s.logger.Error().Err(err).Msg("renderMatchesPage fails")
	}
}

type MatchesPageParam struct {
	Competition *Competition
	ErrMessage  string
	Matches     []*Match
}

var (
	matchesTemplate *template.Template
)

func (s *Server) loadMatchesPageTemplate() error {
	content, err := ioutil.ReadFile("league/html/matches.html")
	if err != nil {
		return err
	}

	s.matchesTemplate, err = template.New("matches").Parse(string(content))
	return err
}

func (s *Server) renderMatchesPage(w http.ResponseWriter, param MatchesPageParam) error {
	if err := s.loadMatchesPageTemplate(); err != nil {
		return err
	}

	return s.matchesTemplate.Execute(w, param)
}
