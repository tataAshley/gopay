package apple

import (
	"github.com/go-pay/gopay"
	"testing"
)

func Test_SubscriptionExtend(t *testing.T) {
	transactionId := "2000000184445477"
	bm := make(gopay.BodyMap)
	bm.Set("extendByDays", 30).
		Set("extendReasonCode", 2).
		Set("requestIdentifier", "1234532")

	rsp, err := client.SubscriptionExtend(ctx, transactionId, bm)
	t.Errorf("SubscriptionExtend res:%+v,err:%+v", rsp, err)
}
