package serializer

// Response 基础出参格式
type Response struct {
	Code   int    `json:"code"`
	Result any    `json:"result"`
	Msg    string `json:"msg"`
}
