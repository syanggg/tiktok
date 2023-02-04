package util

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

type Bind struct {
	FFmpeg string  //ffmpeg path
}

func NewBind(ffmpeg string) *Bind{
	return &Bind{
		FFmpeg: ffmpeg,
	}
}

//Thumbnail 视频缩略图 src输入路径 dst输出路径 duration起始时间 overwrite是否覆盖
func (b *Bind) Thumbnail(src string,dst string,duration time.Duration,overwrite bool) error{
	args := []string{"-i",src,"-ss",fmt.Sprintf("%f",duration.Seconds()),"-vframes","1",dst}
	if overwrite{
		args = append([]string{"-y"},args...)
	}

	fmt.Println(args)
	//使用参数name执行指定的命令
	cmd := exec.Command(b.FFmpeg,args...)


	var out bytes.Buffer
	//指定标准输出
	cmd.Stdout = &out

	//指定标准错误输出
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	//执行
	err := cmd.Run()
	if err != nil{
		return errors.New(fmt.Sprintf("err: %s",stderr.String()))
	}

	return nil
}
