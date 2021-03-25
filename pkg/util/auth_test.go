package util

import (
	"testing"
	"time"

	"github.com/DBoyara/find-course/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateAccessClaims(t *testing.T) {
	Convey("Given some uuid", t, func() {
		uuid := uuid.NewString()

		claim, accessToken := GenerateAccessClaims(uuid)
		Convey("Tokens shuld not be blank", func() {
			So(accessToken, ShouldNotBeBlank)
		})
		Convey("Claim shuld be valid", func() {
			So(claim, ShouldNotBeEmpty)
			So(claim.Issuer, ShouldEqual, uuid)
			So(claim.Subject, ShouldEqual, "access_token")
			So(time.Unix(claim.ExpiresAt, 0), ShouldHappenAfter, time.Now())
			So(time.Unix(claim.IssuedAt, 0), ShouldHappenOnOrBefore, time.Now())
		})
	})
}

func TestGenerateRefreshClaims(t *testing.T) {
	SkipConvey("Given some claim", t, func() {
		uuid := uuid.NewString()
		cl := &models.Claims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
				IssuedAt:  time.Now().Unix(),
				Issuer:    uuid,
				NotBefore: 0,
				Subject:   "refresh_token",
			},
			ID: 0,
		}
		refreshTokenString := GenerateRefreshClaims(cl)
		Convey("Tokens shuld not be blank", func() {
			So(refreshTokenString, ShouldNotBeBlank)
		})
	})
}

func TestGetAuthCookies(t *testing.T) {
	SkipConvey("Set app", t, func() {

	})
}
