// Code generated by goctl. DO NOT EDIT.
package types

type SearchReq struct {
	Name string `json:"name"`
}

type SearchReply struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
