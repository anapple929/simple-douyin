package to_relation

import (
	"context"
	"message/model"
	"message/service/to_relation"
)

type ToRelationService struct {
}

func (*ToRelationService) QueryMessagesByUsers(ctx context.Context, in *to_relation.QueryMessagesByUsersRequest, out *to_relation.QueryMessagesByUsersResponse) error {
	var resultQueryBody []*to_relation.QueryBody
	for _, querybody := range in.QueryBody {
		query_body_temp := model.NewMessageDaoInstance().QueryMessage(querybody.ToUserId, querybody.FromUserId)
		resultQueryBody = append(resultQueryBody, query_body_temp)
	}
	out.StatusCode = 0
	out.QueryBody = resultQueryBody
	return nil
}
