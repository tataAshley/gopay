package apple

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/xlog"
)

var (
	ctx    = context.Background()
	client *Client
	err    error

	iss = "44e94c4f-4685-4515-88a5-5ba4f4d0a900"
	bid = "com.ashley.product.miao"
	kid = "46CCA69F3M"
)

// Iss: "44e94c4f-4685-4515-88a5-5ba4f4d0a900"
// Bid: "com.ashley.product.miao"
// Kid: "46CCA69F3M"
// IsProd: false
// PrivateKey: "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgd1DJnHHzwfde6Uj2\nPhqDPROWPPkggRjVc/h26TJNscygCgYIKoZIzj0DAQehRANCAAQX27G/0KviDKd7\nb8pBdXRsIhNGTYhCQGMqOz6pN7e3lPYN4N6xn0deb77d14BBxwuPXUZAVb9I+c/0\nFAFh4Mbx\n-----END PRIVATE KEY-----"
func TestMain(m *testing.M) {
	xlog.Level = xlog.DebugLevel
	// 初始化客户端
	// iss：issuer ID
	// bid：bundle ID
	// kid：private key ID
	// privateKey：私钥文件读取后的字符串内容
	// isProd：是否是正式环境
	privateKey := ""
	client, err = NewClient(iss, bid, kid, privateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}

	os.Exit(m.Run())
}
