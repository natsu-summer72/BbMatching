package design

import(
	. "goa.design/goa/dsl"
)

var phoneNumberPattern = `^\+?[\d]{10,}$`

var JWT = Type("JWT Token", func(){
	Token("token", String, func(){
		Description("JWT used for Authentication")
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
})

var SessionTokenPayload = Type("SessionTokenPayload", func(){
	Reference(JWT)
	Token("token")
})


var GetJWTPayload = Type("GetJWTPayload", func(){
	Reference(UserProfile)
	Attribute("user_id")
	Required("user_id")
})

// User
var UserProfile = Type("UserProfile", func(){
	Attribute("user_id", String, func(){
		Description("firebaseのユーザーID")
		Example("XRQ85mtXnINISH25zfM0m5RlC6L2")
		MaxLength(28)
		MinLength(28)
	})
	Attribute("email", func(){
		Description("チームのプライマリメールアドレス")
		Format("email")
		Example("chunichi@example.com")
	})
	Attribute("phoneNumber", func(){
		Description("チームのメイン電話番号")
		Pattern(phoneNumberPattern)
		Example("09012345678")
	})
	Attribute("photoURL", func(){
		Description("チームの写真URL")
		Example("http://img.com")
	})
	Attribute("UserName",func(){
		Description("チームの表示名")
		Example("Chunichi Dragons")
	})
	Attribute("email_verified", Boolean, func(){
		Description("ユーザーのプライマリメールアドレスが確認されているか。")
	})
})


// Payload for User
var UpdateUserPayload = Type("UpdateUserPayload", func(){
	Reference(JWT)
	Reference(UserProfile)
	Token("token")
	Attribute("email")
	Attribute("phoneNumber")
	Attribute("photoURL")
	Attribute("UserName")
})

var GetUserPayload = Type("GetUserPayload", func(){
	Reference(JWT)
	Token("token")
	Attribute("user_id", String, func() {
		Description("firebaseのユーザーID")
		Example("XRQ85mtXnINISH25zfM0m5RlC6L2")
		MaxLength(28)
		MinLength(28)
	})
	Required("user_id")
})