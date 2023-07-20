package context_optimize

import (
	openai "github.com/sashabaranov/go-openai"
)

func CleanUpOlderContext(
	contextMessages []openai.ChatCompletionMessage) []openai.ChatCompletionMessage {

	// 当上下文超过4K则进行清理，这里并不是精准的Token数量，只是近似
	cleanUpDataMoreThan := 1024 * 4

	totalDataSize := 0
	// 倒序遍历上下文数据，既req.Messages
	for i := len(contextMessages) - 1; i >= 0; i-- {
		// 计算截止当前的上下文总量
		totalDataSize += len(contextMessages[i].Content)
		// 若超过，则抛弃之前的数据
		if totalDataSize > cleanUpDataMoreThan {
			contextMessages = contextMessages[i:]
			break
		}
	}

	return contextMessages
}
