package shortmsg

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	models "hollow/internal/models"
)

type AliyunShortMsg struct {
	requestInfo *dysmsapi.SendSmsRequest
	client      *dysmsapi.Client
}

func NewAliyunShortMsg() *AliyunShortMsg {
	return &AliyunShortMsg{}
}

func (asm *AliyunShortMsg) Init(accessKeyId, accessKeySecret, regionId, scheme, signname, templatecode string) error {
	asm.requestInfo = dysmsapi.CreateSendSmsRequest()
	asm.requestInfo.Scheme = scheme
	asm.requestInfo.SignName = signname
	asm.requestInfo.TemplateCode = templatecode

	config := sdk.NewConfig()

	credential := credentials.NewAccessKeyCredential(accessKeyId, accessKeySecret)

	var err error

	asm.client, err = dysmsapi.NewClientWithOptions(regionId, config, credential)

	return err
}

func (asm *AliyunShortMsg) getTemplateParam(code string) string {
	return "{\"code\":\"" + code + "\"}"
}

func (asm *AliyunShortMsg) PostMsg(PhoneNumbers, params string) (data *models.ShortMsg, err error) {
	asm.requestInfo.TemplateParam = asm.getTemplateParam(params) //更新参数
	asm.requestInfo.PhoneNumbers = PhoneNumbers

	response, err := asm.client.SendSms(asm.requestInfo)

	if err != nil {
		return nil, err
	}

	// 项目较小，不需要实现短信回执

	return &models.ShortMsg{
		Code:      response.Code,
		Message:   response.Message,
		BizId:     response.BizId,
		RequestId: response.RequestId,
	}, nil
}
