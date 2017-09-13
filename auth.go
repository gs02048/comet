package main

import "encoding/json"

type Auth struct{
	UserId int64 `json:"uid"`
	RoomId string `json:"rid"`
	Token string `json:"token"`
}

func UserAuth(arg json.RawMessage) (uid int64,rid string,err error){
	a := &Auth{}
	if err = json.Unmarshal(arg,a);err != nil{
		return uid,rid,err
	}
	return a.UserId,a.RoomId,nil
}
