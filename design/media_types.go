package design

import (
	. "goa.design/goa/dsl"
)


var UserResponse = ResultType("application/vnd.bbmatching.user+json", func(){
	Description("User Response")
	ContentType("application/json; charset=utf-8")

	Reference(UserProfile)
	Attributes(func(){
		Attribute("user_id")
		Attribute("UserName")
		Attribute("email")
		Attribute("phoneNumber")
		Attribute("photoURL")
		Attribute("email_verified")
		Required("user_id", "UserName", "email", "phoneNumber", "photoURL", "email_verified")
	})

	View("default", func(){
		Attribute("user_id")
		Attribute("email")
		Attribute("phoneNumber")
		Attribute("photoURL")
		Attribute("UserName")
		Attribute("email_verified")
	})

	View("tiny", func(){
		Attribute("UserName")
		Attribute("email")
		Attribute("photoURL")
	})
})


var JWTResponse = ResultType("application/vnd.bbmatching.jwt+json", func(){
	Description("JWT Response")

	Attributes(func(){
			Attribute("JWT", String, func(){
				Description("Json Web Token")
				Example("eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ1c2VyaG9nZSIsImF1ZCI6ImF1ZGhvZ2UiLCJpc3MiOiJodHRwczpcL1wvZXhhbXBsZS5jb21cLyIsImV4cCI6MTQ1MjU2NTYyOCwiaWF0IjoxNDUyNTY1NTY4fQ.BfW2a1SMY1a8cjb7ATcJPwYSmB1P6l4oN2QRNtod-xCyochsB1ZurxNLqPOGFr7_Abqpk-_lOUaPOdL2jQ23T1DS1lmQgaEZgUXaJPAqGSJygANpv8ds07Q6pbX_XbFpJdoVCQHzP8MjbjW_ft2ZAAJzjZZEC6Hm0WUxKS6V0yRNmEUyV-Ljh6n337rtAoSTks2APdgW1hffyZMrptKKazG2m0V0LQRnu5lmWLTYwTZC8NdNJZNepPwQiGNZS0IcrdZhguGAL75ZGTMw9O_EC9gv_I_9I5NUwZk6LG1feEy3MawT0QaTEsF5n6yUKJ8ziuMXnUEsymGdKC-VYEbPyw")
			})
	})
})


var MatchRecruitResponse = ResultType("application/vnd.bbmatching.match_recruit+json", func(){
	Description("Match Recruitment Response")
	ContentType("application/json; charset=utf-8")

	Reference(MatchRecruitProfile)
	Attributes(func(){
		Attribute("id")
		Attribute("user_id")
		Attribute("location")
		Attribute("date")
		Attribute("comment")
		Attribute("disabled")
		Required("id", "user_id", "location", "date", "comment", "disabled")
	})

	View("default", func(){
		Attribute("user_id")
		Attribute("location")
		Attribute("date")
		Attribute("comment")
		Attribute("disabled")
	})
})