package userjoinedteamchannelmessagereply

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

type ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChatMessageCollectionResponse
}

type ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesCompleteResult struct {
	Items []models.ChatMessageCollectionResponse
}

// ListUserByIdJoinedTeamByIdChannelByIdMessageByIdReplies ...
func (c UserJoinedTeamChannelMessageReplyClient) ListUserByIdJoinedTeamByIdChannelByIdMessageByIdReplies(ctx context.Context, id UserJoinedTeamChannelMessageId) (result ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/replies", id.ID()),
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
		Values *[]models.ChatMessageCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesComplete retrieves all the results into a single object
func (c UserJoinedTeamChannelMessageReplyClient) ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesComplete(ctx context.Context, id UserJoinedTeamChannelMessageId) (ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesCompleteResult, error) {
	return c.ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesCompleteMatchingPredicate(ctx, id, models.ChatMessageCollectionResponseOperationPredicate{})
}

// ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserJoinedTeamChannelMessageReplyClient) ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesCompleteMatchingPredicate(ctx context.Context, id UserJoinedTeamChannelMessageId, predicate models.ChatMessageCollectionResponseOperationPredicate) (result ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesCompleteResult, err error) {
	items := make([]models.ChatMessageCollectionResponse, 0)

	resp, err := c.ListUserByIdJoinedTeamByIdChannelByIdMessageByIdReplies(ctx, id)
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

	result = ListUserByIdJoinedTeamByIdChannelByIdMessageByIdRepliesCompleteResult{
		Items: items,
	}
	return
}
