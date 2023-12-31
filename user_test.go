package template_test

import (
	"testing"

	"github.com/Soapstone-Services/go-template-2024"
)

func TestChangePassword(t *testing.T) {
	user := &template.User{
		FirstName: "TestGuy",
	}

	hashedPassword := "h4$h3D"

	user.ChangePassword(hashedPassword)
	if user.LastPasswordChange.IsZero() {
		t.Errorf("Last password change was not changed")
	}

	if user.Password != hashedPassword {
		t.Errorf("Password was not changed")

	}
}

func TestUpdateLastLogin(t *testing.T) {
	user := &template.User{
		FirstName: "TestGuy",
	}

	token := "helloWorld"

	user.UpdateLastLogin(token)
	if user.LastLogin.IsZero() {
		t.Errorf("Last login time was not changed")
	}

	if user.Token != token {
		t.Errorf("Token was not changed")

	}
}
