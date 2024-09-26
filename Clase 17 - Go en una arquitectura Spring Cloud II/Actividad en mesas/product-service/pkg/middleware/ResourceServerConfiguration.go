package middleware

import (
	"context"
	"crypto/tls"
	"net/http"
	"strings"
	"time"

	oidc "github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)

type Res401Struct struct {
	Status   string `json:"status" example:"FAILED"`
	HTTPCode int    `json:"httpCode" example:"401"`
	Message  string `json:"message" example:"authorisation failed"`
}

//"resource_access":{"Gateway":{"roles":["USER","EDITOR"]}},
type Claims struct {
	ResourceAccess client `json:"resource_access,omitempty"`
	JTI            string `json:"jti,omitempty"`
}

type client struct {
	Gateway clientRoles `json:"Gateway,omitempty"`
}

type clientRoles struct {
	Roles []string `json:"roles,omitempty"`
}

var RealmConfigURL string = "http://localhost:8082/realms/digital-house"
var clientID string = "gateway"

func IsAuthorizedJWT(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		rawAccessToken := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{
			Timeout:   time.Duration(6000) * time.Second,
			Transport: tr,
		}
		ctx := oidc.ClientContext(context.Background(), client)
		provider, err := oidc.NewProvider(ctx, RealmConfigURL)
		if err != nil {
			authorizationFailed("authorization failed while getting the provider: "+err.Error(), c)
			return
		}

		oidcConfig := &oidc.Config{
			ClientID: clientID,
		}
		verifier := provider.Verifier(oidcConfig)
		idToken, err := verifier.Verify(ctx, rawAccessToken)
		if err != nil {
			authorizationFailed("authorization failed while verifying the token: "+err.Error(), c)
			return
		}

		var IDTokenClaims Claims
		if err := idToken.Claims(&IDTokenClaims); err != nil {
			authorizationFailed("claims : "+err.Error(), c)
			return
		}

		user_access_roles := IDTokenClaims.ResourceAccess.Gateway.Roles

		//roles[USER,ADMIN]
		for _, userRole := range user_access_roles {
			if userRole == roles[0] && c.FullPath() == "/products/:id" {
				c.Next()
				return
			}
			if userRole == roles[1] {
				c.Next()
				return
			}

		}
		authorizationFailed("user not allowed to access this api", c)
	}
}

func authorizationFailed(message string, c *gin.Context) {
	data := Res401Struct{
		Status:   "FAILED",
		HTTPCode: http.StatusUnauthorized,
		Message:  message,
	}

	c.AbortWithStatusJSON(200, gin.H{"response": data})

}
