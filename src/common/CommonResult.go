package common

type ConmmonResult struct {
	Code int	//200-成功	777-参数有问题		416-出错
	Message string
	Data interface{}
}