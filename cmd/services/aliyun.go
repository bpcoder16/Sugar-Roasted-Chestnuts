package services

import (
	"Sugar-Roasted-Chestnuts/services/aliyun"
	"context"
	"github.com/bpcoder16/Chestnut/core/cmd"
)

type ImageTransfer struct {
	cmd.Base
}

func (d *ImageTransfer) Description() string {
	return "测试 Chestnut 基础库对应的 ImageTransfer 模块"
}

func (d *ImageTransfer) Run(ctx context.Context, _ []string) {
	aliyun.ImageTransferTest(ctx)
}
