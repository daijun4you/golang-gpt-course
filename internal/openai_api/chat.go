package api

import (
	"context"
	"github.com/daijun4you/golang-gpt-course/configs"
	openai "github.com/sashabaranov/go-openai"
)

var chatClient *openai.Client
var chatReq openai.ChatCompletionRequest

func Chat() {
	chatClient = openai.NewClient(configs.Instance.Get("api_key", "openai.ini"))

	chatReq = openai.ChatCompletionRequest{
		// 最新GPT-3.5 16K模型
		Model: openai.GPT3Dot5Turbo16K0613,
		// 限制上下文最大的Token数量
		MaxTokens: 16000,
		// GPT上下文
		Messages: []openai.ChatCompletionMessage{
			// GPT角色设定
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "你是一个资深的心理咨询师",
			},
			// 模拟用户输入信息
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "我觉得GPT很酷！",
			},
		},
		// 这里调节系数，让模型的答复更稳定
		Temperature: 0.2,
		// 不采用流式响应
		Stream: false,
		// 期望GPT每次答复两条（这里只是为了演示，正常情况取值为1）
		N: 2,
	}

	chatCompletion, err := chatClient.CreateChatCompletion(context.Background(), chatReq)
	if err != nil {
		println(err.Error())
		return
	}

	// 第一个答复
	println(chatCompletion.Choices[0].Message.Content)
	// 第二个答复，上边的n>=2时，才会有该条回复
	println(chatCompletion.Choices[1].Message.Content)

	// 加入到上下文中，
	chatReq.Messages = append(chatReq.Messages, chatCompletion.Choices[0].Message)
}
