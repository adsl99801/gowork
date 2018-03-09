package main

//User o
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//BaseResponse o
type BaseResponse struct {
	Sus    bool   `json:"sus"`
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// //New o
// func NewBaseResponse() BaseResponse {
// 	res := BaseResponse{
// 		Sus:    false,
// 		Status: 0,
// 		Msg:    "",
// 	}
// 	return res
// }
