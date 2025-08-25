package service

import (
	"testing"
	"time"

	"github.com/skip2/go-qrcode"
)

func TestLoginQRService(t *testing.T) {
	service := &LoginQRService{}
	code, _, qrcodeUrl, err := service.GetKey()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	t.Log(code, string(qrcodeUrl))
	if code != 200 {
		t.Fatalf("code error: %f", code)
	}
	if qrcodeUrl == "" {
		t.Fatalf("url empty: %s", qrcodeUrl)
	}
	// 将返回的信息制作成二维码，打印至终端
	t.Log("请使用网易云音乐APP扫码登录")
	qr, err := qrcode.New(qrcodeUrl, qrcode.Medium)
	if err != nil {
		t.Fatalf("Failed to generate QR code: %v", err)
	}
	qrString := qr.ToSmallString(false)
	t.Logf("\n%s", qrString)

	// 开始轮询扫码状态
	for {
		code, resp, err := service.CheckQR()
		if err != nil {
			t.Fatalf("Failed to check qrcode: %s", err.Error())
		}
		switch {
		case code == 803:
			t.Logf("扫码成功：%s", string(resp))
			goto Success
		case code == 800:
			t.Log("二维码已失效，请重新获取")
		case code == 801:
			t.Log("等待扫码")
		case code == 802:
			t.Log("已扫码，请在网易云app上确认")
		case code == 8821:
			t.Fatalf("触发风控：%s", string(resp))
		default:
			t.Fatalf("错误，无法识别的状态：%s", string(resp))
		}
		time.Sleep(1 * time.Second)
	}
Success:
	// 扫码成功获取用户信息
	accountService := &UserAccountService{}
	code, _ = accountService.AccountInfo()
	if code != 200 {
		t.Fatalf("code error: %f", code)
	}
}
