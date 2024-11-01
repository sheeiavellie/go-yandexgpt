package yandexgpt

type YandexGPTMessage struct {
	Role         yandexGPTRole `json:"role"`
	Text         string        `json:"text"`
	ToolCallList *ToolCallList  `json:"toolCallList"`
}

type YandexToolResultList struct {
	ToolResults []YandexFunctionResult `json:"toolResults"`
}

type YandexFunctionResult struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type ToolCallList struct {
	ToolCalls []FunctionCall `json:"toolCalls"`
}
type FunctionCall struct {
	Name      string `json:"name"`
	Arguments any    `json:"arguments"`
}
