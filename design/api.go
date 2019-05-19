package design

import (
	. "goa.design/goa/dsl"
)

var _ = API("BbM", func(){
	Title("Baseball Matching")
	Description("野球チームのマッチングアプリケーション")
	Version("1.0")


	Server("BbMatching",func(){
		Services("User", "MatchRecruit")
		Host("localhost", func(){
			Description("development host")
			URI("http://localhost:8080")
		})
	})
})


var JWTAuth = JWTSecurity("jwt", func(){
	Description("Secures endpoint by requiring a valid JWT token retrieved via the firebase.")
})

