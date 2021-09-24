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

	// テスト失敗用
	// message := fmt.Sprintf(randomFormat())

	return message, nil
}

// map[string]stringはstring => stringの連想配列？
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	// for ~ range ~ で繰り返し
	// forの変数は_で省略可能
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
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
