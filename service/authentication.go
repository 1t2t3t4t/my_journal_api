package service

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/1t2t3t4t/my_journal_api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

const UserClaim = "User_Claim"
const tokenAge time.Duration = time.Hour * 72

var hmacSecretDefault string = "SECRET"
var hmacSecretByte []byte

func init() {
	hmacSecret := os.Getenv("AUTH_SECRET")
	if hmacSecret == "" {
		hmacSecret = hmacSecretDefault
	}
	hmacSecretByte = []byte(hmacSecret)
}

type AuthClaim struct {
	Uid string
}

func CreateAuthToken(claim AuthClaim) (string, error) {
	expiryUnix := time.Now().UTC().Add(tokenAge).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": claim.Uid,
		"exp": expiryUnix,
	})
	return token.SignedString(hmacSecretByte)
}

func ValidateAuthToken(tokenStr string) (AuthClaim, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecretByte, nil
	})
	if err != nil {
		return AuthClaim{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return AuthClaim{Uid: claims["uid"].(string)}, nil
	} else {
		return AuthClaim{}, fmt.Errorf("invalid token")
	}
}

func AuthMiddleware(userRepository database.UserRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := ctx.Get(fiber.HeaderAuthorization)
		if split := strings.Split(auth, " "); len(split) == 2 && strings.ToLower(split[0]) == "bearer" {
			if claim, err := ValidateAuthToken(split[1]); err == nil {
				user := userRepository.FindOne(claim.Uid)
				if user != nil {
					ctx.Locals(UserClaim, claim)
				} else {
					log.Printf("User not found for token %v", auth)
				}
			} else {
				log.Printf("Invalid token %v", auth)
			}
		}
		return ctx.Next()
	}
}
