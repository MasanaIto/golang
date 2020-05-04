package main

import "fmt"

/*
UserNotFound ...
*/
type UserNotFound struct {
	Username string
}

// エラーを作る時は、ポインタで*, ＆を付けるのが慣習
func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	// 他のエラー
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "Masana"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err) // User not found: Masana
	}
}
