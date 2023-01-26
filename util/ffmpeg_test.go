package util

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"
)

func TestBind_Thumbnail(t *testing.T) {
	bind := NewBind(Info.Path.FfmpegPath)
	input := filepath.Join(Info.Path.StaticPath,"haha.mp4")
	output := filepath.Join(Info.Path.StaticPath,"haha.jpg")

	//input := "../static/haha.mp4"
	//output := "../static/haha.jpg"

	//file,err := os.Open(input)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(file.Name())
	//
	//fmt.Println(input)
	//fmt.Println(output)
	err := bind.Thumbnail(input,output,1*time.Second,false)
	if err != nil{
		fmt.Println(err)
	}
}
