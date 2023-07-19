package api

import (
	"context"
	"github.com/daijun4you/golang-gpt-course/configs"
	openai "github.com/sashabaranov/go-openai"
)

var embeddingClient *openai.Client
var embeddingReq openai.EmbeddingRequest

func Embedding() {
	embeddingClient = openai.NewClient(configs.Instance.Get("api_key", "openai.ini"))

	embeddingReq = openai.EmbeddingRequest{
		// 选择 text-embedding-ada-002 模型
		Model: openai.AdaEmbeddingV2,
		Input: "今天天气如何",
	}

	embedding, err := embeddingClient.CreateEmbeddings(context.Background(), embeddingReq)
	if err != nil {
		println(err.Error())
		return
	}

	println(embedding.Data[0].Embedding)
}
