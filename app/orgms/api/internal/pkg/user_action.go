package pkg

import (
	"context"
	"encoding/json"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"uims/api/ums/rpc"
	"uims/pkg/defines"
	"uims/pkg/kafka"
)

type UserAction struct {
	data *rpc.UserActionMsg
	kp   *kafka.KafkaPubClient
}

func NewUserAction(kp *kafka.KafkaPubClient, uid int64, actionType rpc.ACTION_TYPE, detail interface{}) *UserAction {
	b, _ := json.Marshal(detail)
	ua := &UserAction{
		kp: kp,
		data: &rpc.UserActionMsg{
			Uid:          uid,
			ActionId:     1,
			AppId:        defines.APP_ID_ORGMS,
			ActionType:   actionType,
			ActionDetail: string(b),
			ActionTime:   timestamppb.New(time.Now()),
		},
	}

	return ua
}

func (ua *UserAction) Send(ctx context.Context) {
	msg, _ := json.Marshal(ua.data)
	ua.kp.Pub(ctx, defines.USER_ACTION_TOPIC, msg)
}
