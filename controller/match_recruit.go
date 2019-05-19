package bbm

import (
	"context"
	"database/sql"
	"firebase.google.com/go/auth"
	"log"
	"strconv"

	matchrecruit "github.com/natsu-summer72/BbMatching/gen/match_recruit"
)

// MatchRecruit service example implementation.
// The example methods log the requests and return zero values.
type matchRecruitsrvc struct {
	logger *log.Logger
	authClient *auth.Client
	Database *sql.DB
}

type Recruit struct {
	id int
	user_id string
	location string
	date string
	comment string
	disabled bool
	created_at string
}


// NewMatchRecruit returns the MatchRecruit service implementation.
func NewMatchRecruit(logger *log.Logger, client *auth.Client, db *sql.DB) matchrecruit.Service {
	return &matchRecruitsrvc{logger, client, db}
}

// 募集中である試合の一覧を返します。
func (s *matchRecruitsrvc) ListMatchRecruitment(ctx context.Context, p *matchrecruit.SessionTokenPayload) (res matchrecruit.BbmatchingMatchRecruitCollection, err error) {
	s.logger.Print("matchRecruit.List Match Recruitment")

	rows, err := s.Database.Query("SELECT * FROM match_recruit")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		recruit:=Recruit{}
		err = rows.Scan(&recruit.id, &recruit.user_id, &recruit.location, &recruit.date, &recruit.comment, &recruit.disabled, &recruit.created_at)
		if err!=nil {
			return
		}
		r := &matchrecruit.BbmatchingMatchRecruit{
			ID: recruit.id,
			UserID: recruit.user_id,
			Location: recruit.location,
			Date: recruit.date,
			Comment: recruit.comment,
			Disabled: recruit.disabled,
		}
		if recruit.disabled == false{
			res = append(res, r)
		}
	}

	return
}

// 指定したIDの募集試合を取得
func (s *matchRecruitsrvc) GetMatchRecruit(ctx context.Context, p *matchrecruit.GetMatchRecruitPayload) (res *matchrecruit.BbmatchingMatchRecruit, err error) {
	res = &matchrecruit.BbmatchingMatchRecruit{}
	s.logger.Print("matchRecruit.Get Match Recruit")

	recruit := Recruit{}
	query := "SELECT * FROM match_recruit where(id="+ strconv.Itoa(p.ID) + ");"
	rows, err := s.Database.Query(query)
	defer rows.Close()

	rows.Next()
	if err!=nil {
		return
	}
	err = rows.Scan(&recruit.id, &recruit.user_id, &recruit.location, &recruit.date, &recruit.comment, &recruit.disabled, &recruit.created_at)

	res = &matchrecruit.BbmatchingMatchRecruit{
		ID: recruit.id,
		UserID: recruit.user_id,
		Location: recruit.location,
		Date: recruit.date,
		Comment: recruit.comment,
		Disabled: recruit.disabled,
	}
	return
}
