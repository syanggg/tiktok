package cache

import (
	"fmt"
	"testing"
)

func TestInitRedis(t *testing.T) {
	err := InitRedis()
	if err != nil{
		fmt.Println("failed to connection :",err)
	}

	op := NewRedisOperation()

	////op.AddFavorite(1,1)
	////op.AddFavorite(1,2)
	//
	//op.CancelFavorite(1,1)
	//op.CancelFavorite(1,2)
	//
	list,err := op.GetFavoriteVideoIdList(2)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(list)



	//result,err := op.GetFavoriteState(1,1)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(result)
}
