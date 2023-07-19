package main

import (
	openaiApi "github.com/daijun4you/golang-gpt-course/internal/openai_api"
	"os"
)

var funcList = map[string]map[string]func(){
	"openai_api": {
		"chat": openaiApi.Chat,
	},
}

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}

	module, exists := funcList[args[1]]
	if !exists {
		println("Module: '" + args[1] + "' not found")
		return
	}

	moduleFunc, exists := module[args[2]]
	if !exists {
		println("Module Demo: '" + args[1] + "." + args[2] + "' not found")
		return
	}

	moduleFunc()
}
