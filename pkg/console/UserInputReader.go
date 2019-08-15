package console

import (
	"bufio"
	"fmt"
	"os"
)

// UserInputReader ユーザーからの標準入力を読み込む
type UserInputReader struct {
	promptSymbol string
}

// NewUserInputReader return pointer of type UserInputReader struct
func NewUserInputReader(promptSymbol string) *UserInputReader {
	return &UserInputReader{
		promptSymbol: promptSymbol,
	}
}

// ReadUserInput from console
func (userInputReader *UserInputReader) ReadUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	userInput, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(userInput)
}

// Prompt prints some short string. usually use before ReadUserInput func
func (userInputReader *UserInputReader) Prompt() {
	fmt.Print(userInputReader.promptSymbol)
}
