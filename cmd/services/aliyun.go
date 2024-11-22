package services

import (
	"Sugar-Roasted-Chestnuts/services/aliyun"
	"context"
	"fmt"
	"github.com/bpcoder16/Chestnut/contrib/aliyun/oss"
	"github.com/bpcoder16/Chestnut/modules/cmd"
)

type AliyunOSSImageTransfer struct {
	cmd.Base
}

func (a *AliyunOSSImageTransfer) Description() string {
	return "测试 Chestnut 基础库对应的 AliyunOSSImageTransfer 模块"
}

func (a *AliyunOSSImageTransfer) Run(ctx context.Context, _ []string) {
	aliyun.ImageTransferTest(ctx)
}

type AliyunOSSSignURL struct {
	cmd.Base
}

func (a *AliyunOSSSignURL) Description() string {
	return "测试 Chestnut 基础库对应的 AliyunOSSSignURL 模块"
}

func (a *AliyunOSSSignURL) Run(ctx context.Context, _ []string) {
	signURL, err := oss.SignURL("test/01/01/baidu_cover.jpg", 10)
	fmt.Println(signURL)
	if err != nil {
		fmt.Println(err.Error())
	}
}
