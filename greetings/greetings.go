package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 引数nameへの挨拶を返す
// 頭文字が大文字の関数は外部から使える(exportされた)関数
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// ':=' は変数宣言と同時に代入する記法
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// package内の変数の初期化が終わったら呼ばれる？
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// []stringでstring型のsliceの宣言
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// formatsの中からrandomなstringを返す
	return formats[rand.Intn(len(formats))]
}
