package utils

import(
	"fmt"
	"os"
	"time"
	"log"
	"github.com/golang-jwt/jwt/v4"
)

type RegisterDetails struct{
	ID 			string `json:"ID"`
	Email 		string	`json:"email`

	jwt.RegisteredClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenToken(ID string, Email string)(signedToken string, signedRefreshToken string, err error){
	claims:= &RegisterDetails{
		ID : ID,
		Email : Email,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	refreshClaims := &RegisterDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)),
		},
	}

	token ,err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return 
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return 
	}

	return token, refreshToken, err
}

func AuthToken(signedToken string) (claims *RegisterDetails, msg string){
		token, err := jwt.ParseWithClaims(
			signedToken,
			&RegisterDetails{},
			func(token *jwt.Token)(interface{}, error){
				return []byte(SECRET_KEY), nil
			},
		)
		
		if err != nil {
			msg=err.Error()
			return
		}
	
		claims, ok:= token.Claims.(*RegisterDetails)
		if !ok{
			msg = fmt.Sprintf("the token is invalid")
			msg = err.Error()
			return
		}
	
		if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Unix() < time.Now().Unix(){
			msg = fmt.Sprint("token is expired")
			msg = err.Error()
			return
		}
		return claims, msg
	}
