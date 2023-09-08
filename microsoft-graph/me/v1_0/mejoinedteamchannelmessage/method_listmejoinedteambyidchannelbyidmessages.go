package mejoinedteamchannelmessage

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

type ListMeJoinedTeamByIdChannelByIdMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChatMessageCollectionResponse
}

type ListMeJoinedTeamByIdChannelByIdMessagesCompleteResult struct {
	Items []models.ChatMessageCollectionResponse
}

// ListMeJoinedTeamByIdChannelByIdMessages ...
func (c MeJoinedTeamChannelMessageClient) ListMeJoinedTeamByIdChannelByIdMessages(ctx context.Context, id MeJoinedTeamChannelId) (result ListMeJoinedTeamByIdChannelByIdMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/messages", id.ID()),
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

// ListMeJoinedTeamByIdChannelByIdMessagesComplete retrieves all the results into a single object
func (c MeJoinedTeamChannelMessageClient) ListMeJoinedTeamByIdChannelByIdMessagesComplete(ctx context.Context, id MeJoinedTeamChannelId) (ListMeJoinedTeamByIdChannelByIdMessagesCompleteResult, error) {
	return c.ListMeJoinedTeamByIdChannelByIdMessagesCompleteMatchingPredicate(ctx, id, models.ChatMessageCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdChannelByIdMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamChannelMessageClient) ListMeJoinedTeamByIdChannelByIdMessagesCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamChannelId, predicate models.ChatMessageCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdChannelByIdMessagesCompleteResult, err error) {
	items := make([]models.ChatMessageCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdChannelByIdMessages(ctx, id)
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

	result = ListMeJoinedTeamByIdChannelByIdMessagesCompleteResult{
		Items: items,
	}
	return
}
