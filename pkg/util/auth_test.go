package util

import (
	"log"
	"testing"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateTokens(t *testing.T) {
	SkipConvey("Given some uuid", t, func() {  // TODO create tests for auth
		uuid := uuid.NewString()
		log.Println(uuid)
		accessToken, refreshToken := GenerateTokens(uuid)
		Convey("Tokens shuld not be blank", func() {
			So(accessToken, ShouldPanic)
			So(refreshToken, ShouldPanic)
		})
	})
}
