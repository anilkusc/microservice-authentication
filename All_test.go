package main

import "testing"

type TestItems struct {
	user   User
	output bool //expected value
}

var Users = []TestItems{
	{User{Username: "rasim", Password: "rasim", Role: "admin"}, true},
	{User{Username: "anil", Password: "anil", Role: "no"}, false},
}

func TestCheckUser(t *testing.T) {

	for _, item := range Users {
		isAuth, _ := CheckUser(item.user)
		if item.output {
			// expected an error
			if isAuth == false {
				t.Errorf("CheckUser() with username %v , password %v : FAILED, expected an true but got value '%v'", item.user.Username, item.user.Password, isAuth)
			} else {
				t.Logf("CheckUser() with username %v , password %v  : PASSED, expected an true and got value '%v'", item.user.Username, item.user.Password, isAuth)
			}
		} else {
			// expected a value
			if isAuth == true {
				t.Errorf("CheckUser() with username %v , password %v : FAILED, expected an false but got value '%v'", item.user.Username, item.user.Password, isAuth)
			} else {
				t.Logf("CheckUser() with username %v , password %v  : PASSED, expected an false and got value '%v'", item.user.Username, item.user.Password, isAuth)
			}
		}
	}
}
