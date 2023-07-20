package api

import (
	"context"
	"github.com/daijun4you/golang-gpt-course/configs"
	openai "github.com/sashabaranov/go-openai"
)

func Embedding() {
	// 初始化client，并设置api key
	client := openai.NewClient(configs.Instance.Get("api_key", "openai.ini"))
	req := openai.EmbeddingRequest{
		// 选择 text-embedding-ada-002 模型
		Model: openai.AdaEmbeddingV2,
		Input: "今天天气如何",
	}

	embedding, err := client.CreateEmbeddings(context.Background(), req)
	if err != nil {
		println(err.Error())
		return
	}

	println(embedding.Data[0].Embedding)
}
