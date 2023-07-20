package api

import (
	"context"
	"github.com/daijun4you/golang-gpt-course/configs"
	openai "github.com/sashabaranov/go-openai"
)

func Completion() {
	// 初始化client，并设置api key
	client := openai.NewClient(configs.Instance.Get("api_key", "openai.ini"))
	req := openai.CompletionRequest{
		// 最新GPT-3.5 16K模型
		Model: openai.GPT3TextDavinci003,
		// 限制上下文最大的Token数量
		MaxTokens: 4000,
		// 提示词
		Prompt: "请介绍下自己",
		// 在GPT答复信息中需要插入的信息
		Suffix: "菠菜GPT技术课程",
		// 这里调节系数，让模型的答复更稳定
		Temperature: 0.2,
		// 不采用流式响应
		Stream: false,
		// 期望GPT每次答复两条（这里只是为了演示，正常情况取值为1）
		N:    1,
		Echo: false,
	}

	completion, err := client.CreateCompletion(context.Background(), req)
	if err != nil {
		println(err.Error())
		return
	}

	// GPT答复
	println(completion.Choices[0].Text)
}
