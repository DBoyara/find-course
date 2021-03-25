package util

import (
	"testing"

	"github.com/DBoyara/find-course/pkg/models"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_isEmpty(t *testing.T) {
	Convey("Given some username", t, func() {
		userName := "admin"
		userNameEmpty := " "
		Convey("When username valid", func() {
			isStringEmty, errString := isEmpty(userName)
			Convey("isStringEmty shuld be false", func() {
				So(isStringEmty, ShouldBeFalse)
				So(errString, ShouldBeBlank)
			})
		})
		Convey("When username is empty", func() {
			isStringEmty, errString := isEmpty(userNameEmpty)
			Convey("isStringEmty shuld be True", func() {
				So(isStringEmty, ShouldBeTrue)
				So(errString, ShouldNotBeBlank)
			})
		})
	})
}

func TestValidateRegister(t *testing.T) {
	Convey("Set Users", t, func() {
		validUser := &models.User{Email: "test@test.com", Username: "admin", Password: "Qwer1234"}
		notValidEmail := &models.User{Email: "test", Username: "admin", Password: "Qwer1234"}
		notValidPassword := &models.User{Email: "test@test.com", Username: "admin", Password: "qwer1234"}

		Convey("When user's credentionals is valid", func() {
			userError := ValidateRegister(validUser)
			Convey("userError.Err shuld be false", func() {
				So(userError.Err, ShouldBeFalse)
				So(userError, ShouldHaveSameTypeAs, &models.UserErrors{})
			})
		})

		Convey("When user's email is not valid", func() {
			userError := ValidateRegister(notValidEmail)
			Convey("errString shuld be not blank", func() {
				So(userError.Err, ShouldBeTrue)
				So(userError.Email, ShouldContainSubstring, "valid email")
			})
		})

		Convey("When user's password is not valid", func() {
			userError := ValidateRegister(notValidPassword)
			Convey("errString shuld be not blank", func() {
				So(userError.Err, ShouldBeTrue)
				So(userError.Password, ShouldContainSubstring, "Length of password")
			})
		})
	})
}
