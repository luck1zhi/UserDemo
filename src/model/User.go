package model

type User struct {
	
	UserId string `json:"userid" binding:"required,checkUserId"`
	NickName string `json:"nickname" binding:"required,checkUserName"`
	Role uint32 `json:"role" binding:"eq=1|eq=2"` //1 anchor,2 audience
	Age int64 `json:"age" binding:"gte=0,lt=140"`
}
