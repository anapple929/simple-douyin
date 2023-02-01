package utils

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gofrs/uuid"
	"os"
	"time"
)

//将视频上传到阿里云,返回视频地址
func UploadVideo(videobytes []byte) string {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com", "LTAI5t8K2KtWjUGnwek7BmSn", "yMhZfvo165xVG0ILQMLQboU1G0iYZl")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("simple-douyin-1122233")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	uuid, _ := uuid.NewV4()
	fileDir := time.Now().Format("2006-01-02") + "/" + uuid.String() + ".mp4"
	// 将Byte数组上传至exampledir目录下的exampleobject.txt文件。
	//fmt.Println(videobytes)
	err = bucket.PutObject(fileDir, bytes.NewReader(videobytes))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return "https://" + "simple-douyin-1122233" + ".oss-cn-hangzhou.aliyuncs.com/" + fileDir
}

//将视频封面上传到阿里云,返回视频封面地址
func UploadPicture(picturebytes []byte) string {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com", "LTAI5t8K2KtWjUGnwek7BmSn", "yMhZfvo165xVG0ILQMLQboU1G0iYZl")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("simple-douyin-1122233")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	uuid, _ := uuid.NewV4()
	fileDir := time.Now().Format("2006-01-02") + "/" + uuid.String() + ".jpg"
	// 将Byte数组上传至exampledir目录下的exampleobject.txt文件。
	//fmt.Println(videobytes)
	err = bucket.PutObject(fileDir, bytes.NewReader(picturebytes))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return "https://" + "simple-douyin-1122233" + ".oss-cn-hangzhou.aliyuncs.com/" + fileDir
}
