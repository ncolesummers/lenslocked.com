package models

import (
	"fmt"
	"testing"
	"time"
)

func testingUserService() (*UserService, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "lenslocked_test"
	)
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d  sslmode=disable", user, password, dbname, host, port)
	us, err := NewUserService(psqlInfo)
	if err != nil {
		return nil, err
	}
	// Clear the User Table between tests
	us.DestructiveReset()
	return us, nil
}

func TestCreateUser(t *testing.T) {
	us, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}
	user := User{
		Name: "Micheal Scott",
		Email: "michael@dundermifflin.com",
	}
	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}
	if user.ID == 0 {
		t.Errorf("Expected ID > 0.  Received %d", user.ID)
	}
	if time.Since(user.CreatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected CreatedAt to be recent.  Received %s", user.CreatedAt)
	}
	if time.Since(user.UpdatedAt) > time.Duration(5*time.Second) {
		t.Errorf("Expected UpdatedAt to be recent.  Received %s", user.UpdatedAt)
	}


}