package serializer

// Response 基础出参格式
type Response struct {
	Code   int    `json:"code"`
	Result any    `json:"result"`
	Msg    string `json:"msg"`
}

func NewResponse(code int, result any, msg string) Response {
	return Response{code, result, msg}
}

func SuccessResponse(result any) Response {
	return NewResponse(0, result, "ok")
}

//DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

//TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

//TrackedErrorResponse 有追踪信息的错误反应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

//SuccessListResponse 带有总数的列表构建器
func SuccessListResponse(items interface{}, total uint) Response {
	return Response{
		Code: 200,
		Result: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}
