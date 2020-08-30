package String

import (
	"regexp"
	"strings"
)

/*
*字节切片根据分隔符进行转换为字符串数组
*/

func ByteArrayToStringArrayBySep(b []byte,sep string)([]string){
	str := string(b)
	reg := regexp.MustCompile(sep)
	resList := reg.Split(str,-1)
	return  resList
}

/*
*字节切片根据分隔符替换后进行转换字符串数组
*/

func ByteArrayToStringArrayByNewSeq(b []byte,sep string,newseq string,f string)([]string){
	str := string(b)
	reg := regexp.MustCompile(sep)
	strtmp := reg.ReplaceAllString(str,newseq)
	resList := strings.Split(strtmp,f)
	return resList
}