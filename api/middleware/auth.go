package middleware

import (
	"errors"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/kconde2/vote-app/api/db"
	"github.com/kconde2/vote-app/api/models"
)

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "uuid"

// AuthMiddleware : The authmiddlare handling jwt token security
func AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("$2a$10$GCiJQcAqSaPV8.bU/mvGiOgdHV8GuMOdmW6.nUpCRisfUx9b.VGqy"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		LoginResponse:   LoginResponse,
	})
}

// Callback function that should perform the authentication of the user based on login info.
// Must return user data as user identifier, it will be stored in Claim Array. Required.
// Check error (e) to determine the appropriate error message.
func authenticator(c *gin.Context) (interface{}, error) {
	var loginVals login
	var user models.User
	var db = db.GetDB()

	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	// request fields
	email := loginVals.Email
	password := loginVals.Password

	err := db.Where("email = ?", email).First(&user)

	if err.RecordNotFound() {
		return nil, errors.New("User not found")
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println(err)
		// If the two passwords don't match, return a 401 status
		return nil, jwt.ErrFailedAuthentication
	}

	return &models.User{
		Email:       user.Email,
		AccessLevel: user.AccessLevel,
		UUID:        user.UUID,
	}, nil
}

// LoginResponse can define own LoginResponse func.
func LoginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, gin.H{
		"jwt": token,
	})
}

// Callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
func payloadFunc(data interface{}) jwt.MapClaims {
	log.Println("payloadFunc")
	if v, ok := data.(*models.User); ok {
		return jwt.MapClaims{
			"uuid":        v.UUID,
			"accessLevel": v.AccessLevel,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)

	return &models.User{
		Email: claims[identityKey].(string),
	}
}

// Callback function that should perform the authorization of the authenticated user. Called
// only after an authentication success. Must return true on success, false on failure.
func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.User); ok && v.Email == "admin" {
		return true
	}

	return false
}