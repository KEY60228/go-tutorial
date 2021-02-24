package greetings

import (
	"errors"
	"fmt"
)

// 引数nameへの挨拶を返す
// 頭文字が大文字の関数は外部から使える(exportされた)関数
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// ':=' は変数宣言と同時に代入する記法
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}