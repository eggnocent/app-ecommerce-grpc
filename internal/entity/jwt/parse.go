package jwt

import (
	"context"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"strings"

	"google.golang.org/grpc/metadata"
)

func ParseTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", utils.UnauthenticatedResponse()
	}

	bearerToken, ok := md["authorization"]
	if !ok || len(bearerToken) == 0 {
		return "", utils.UnauthenticatedResponse()
	}

	tokenSplit := strings.Split(bearerToken[0], " ")
	if len(tokenSplit) != 2 || tokenSplit[0] != "Bearer" {
		return "", utils.UnauthenticatedResponse()
	}

	return tokenSplit[1], nil
}
