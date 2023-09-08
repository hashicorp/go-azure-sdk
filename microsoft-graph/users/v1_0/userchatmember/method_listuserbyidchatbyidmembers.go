package userchatmember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdChatByIdMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ConversationMemberCollectionResponse
}

type ListUserByIdChatByIdMembersCompleteResult struct {
	Items []models.ConversationMemberCollectionResponse
}

// ListUserByIdChatByIdMembers ...
func (c UserChatMemberClient) ListUserByIdChatByIdMembers(ctx context.Context, id UserChatId) (result ListUserByIdChatByIdMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/members", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]models.ConversationMemberCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdChatByIdMembersComplete retrieves all the results into a single object
func (c UserChatMemberClient) ListUserByIdChatByIdMembersComplete(ctx context.Context, id UserChatId) (ListUserByIdChatByIdMembersCompleteResult, error) {
	return c.ListUserByIdChatByIdMembersCompleteMatchingPredicate(ctx, id, models.ConversationMemberCollectionResponseOperationPredicate{})
}

// ListUserByIdChatByIdMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserChatMemberClient) ListUserByIdChatByIdMembersCompleteMatchingPredicate(ctx context.Context, id UserChatId, predicate models.ConversationMemberCollectionResponseOperationPredicate) (result ListUserByIdChatByIdMembersCompleteResult, err error) {
	items := make([]models.ConversationMemberCollectionResponse, 0)

	resp, err := c.ListUserByIdChatByIdMembers(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListUserByIdChatByIdMembersCompleteResult{
		Items: items,
	}
	return
}
