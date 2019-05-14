package design

import (
	. "goa.design/goa/dsl"
)


// サービスの定義
var _ = Service("User",func(){
	Description("ユーザー(チーム)に関するエンドポイント")

	Security(JWTAuth)
	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func(){
		Path("/v1")
		Response("unauthorized", StatusUnauthorized)
	})

	Method("Get current user", func(){
		Description("現在のエンドポイントに紐づくユーザーの情報を返します。")

		Payload(SessionTokenPayload)
		Result(UserResponse)

		HTTP(func(){
			GET("/users/me")
			Response(StatusOK)
		})
	})

	Method("Get User", func(){
		Description("指定したIDのユーザーの情報を取得します。")

		Payload(GetUserPayload)
		Result(UserResponse)
		HTTP(func(){
			GET("/users/{user_id}")
			Response(StatusOK)
		})
	})

	Method("List User", func(){
		Description("ユーザーの一覧を取得します。")

		Payload(SessionTokenPayload)
		Result(CollectionOf(UserResponse))
		HTTP(func(){
			GET("/users")
			Response(StatusOK)
		})
	})

	Method("Update current user", func(){
		Description("現在のセッションに紐づくユーザーの情報を更新します。")

		Payload(UpdateUserPayload)
		Result(UserResponse)

		HTTP(func(){
			PUT("/users")
			Response(StatusOK)
		})
	})

	Method("Delete current user", func(){
		Description("現在のセッションに紐づくユーザーを削除します。")

		Payload(SessionTokenPayload)
		HTTP(func(){
			DELETE("/users")
			Response(StatusNoContent)
		})
	})
})

