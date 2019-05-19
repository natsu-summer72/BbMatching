package design

import (
	. "goa.design/goa/dsl"
)

var _ = Service("MatchRecruit", func(){
	Description("試合の募集に関するエンドポイントです")

	Security(JWTAuth)
	Error("unauthorized", String)

	HTTP(func(){
		Path("/v1/match_recruit")
		Response("unauthorized", StatusUnauthorized)
	})

	Method("List Match Recruitment", func(){
		Description("募集中である試合の一覧を返します。")

		Payload(SessionTokenPayload)
		Result(CollectionOf(MatchRecruitResponse))

		HTTP(func(){
			GET("/match")
			Response(StatusOK)
		})
	})

})