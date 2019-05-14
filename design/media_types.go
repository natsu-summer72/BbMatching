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
		Attribute("email")
		Attribute("phoneNumber")
		Attribute("photoURL")
		Attribute("UserName")
		Required("user_id", "email", "phoneNumber", "photoURL", "UserName")
	})

	View("default", func(){
		Attribute("user_id")
		Attribute("email")
		Attribute("phoneNumber")
		Attribute("photoURL")
		Attribute("UserName")
	})

	View("tiny", func(){
		Attribute("UserName")
		Attribute("email")
		Attribute("photoURL")
	})
})