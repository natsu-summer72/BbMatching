package bbm

import (
	"context"
	"github.com/natsu-summer72/BbMatching/gen/user"
	"log"

	"goa.design/goa/security"
)

// JWTAuth implements the authorization logic for service "User" for the "jwt"
// security scheme.
func (s *usersrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	verifiedToken, err := s.authClient.VerifyIDToken(ctx, token)

	if err!=nil{
		panic(err)
		return ctx, user.Unauthorized("invalid token")
	}

	log.Printf("Verified ID Token: %v\n", verifiedToken)
	return ctx, nil
}
