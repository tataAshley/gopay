package apple

import (
	"github.com/go-pay/gopay"
	"testing"
)

//
//func TestClient_SubscriptionExtend(t *testing.T) {
//	type fields struct {
//		iss        string
//		bid        string
//		kid        string
//		isProd     bool
//		privateKey *ecdsa.PrivateKey
//		hc         *xhttp.Client
//	}
//	type args struct {
//		ctx           context.Context
//		transactionId string
//		bm            gopay.BodyMap
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantRsp *SubscriptionsExtendRsp
//		wantErr bool
//	}{
//		{
//			name: "subscription extend ",
//			fields: fields{
//				iss:        "",
//				bid:        "",
//				kid:        "",
//				isProd:     false,
//				privateKey: nil,
//				hc:         nil,
//			},
//			args: args{
//				ctx:           nil,
//				transactionId: "2000000184445477",
//				bm:            nil,
//			},
//			wantRsp: &SubscriptionsExtendRsp{
//				StatusCodeErr: StatusCodeErr{
//					ErrorCode:    0,
//					ErrorMessage: "",
//				},
//			},
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &Client{
//				iss:        tt.fields.iss,
//				bid:        tt.fields.bid,
//				kid:        tt.fields.kid,
//				isProd:     tt.fields.isProd,
//				privateKey: tt.fields.privateKey,
//				hc:         tt.fields.hc,
//			}
//			gotRsp, err := c.SubscriptionExtend(tt.args.ctx, tt.args.transactionId, tt.args.bm)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("SubscriptionExtend() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(gotRsp, tt.wantRsp) {
//				t.Errorf("SubscriptionExtend() gotRsp = %v, want %v", gotRsp, tt.wantRsp)
//			}
//		})
//	}
//}

func TestClient_SubscriptionExtend(t *testing.T) {
	transactionId := "2000000184445477"
	bm := make(gopay.BodyMap)
	bm.Set("extendByDays", 30).
		Set("extendReasonCode", 2).
		Set("requestIdentifier", "1234532")

	rsp, err := client.SubscriptionExtend(ctx, transactionId, bm)
	t.Errorf("SubscriptionExtend res:%+v,err:%+v", rsp, err)
}
