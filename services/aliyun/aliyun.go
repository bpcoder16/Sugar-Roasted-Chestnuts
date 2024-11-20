package aliyun

import (
	"context"
	"github.com/bpcoder16/Chestnut/logit"
	"github.com/bpcoder16/Chestnut/modules/image/aliyunoss"
)

func ImageTransferTest(ctx context.Context) {
	originURL := "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"
	err := aliyunoss.ImageTransfer(ctx, originURL, "test/01/01/baidu_cover.jpg")
	if err != nil {
		logit.Context(ctx).WarnW("aliyunImageTransfer.Err", err)
	}

	targetOSSPath := aliyunoss.BuildTargetOSSPath("test/01/02", originURL)
	err = aliyunoss.ImageTransfer(ctx, originURL, targetOSSPath)
	if err != nil {
		logit.Context(ctx).WarnW("aliyunImageTransfer.Err", err)
	}

	//var signURL string
	//signURL, err = oss.SignURL(targetOSSPath, 10)
	//logit.Context(ctx).InfoW("signURL", signURL, "err", err)
}
