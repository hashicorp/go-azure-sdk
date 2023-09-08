package mejoinedteamprimarychannelmessagehostedcontent

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

type ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChatMessageHostedContentCollectionResponse
}

type ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContentsCompleteResult struct {
	Items []models.ChatMessageHostedContentCollectionResponse
}

// ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContents ...
func (c MeJoinedTeamPrimaryChannelMessageHostedContentClient) ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContents(ctx context.Context, id MeJoinedTeamPrimaryChannelMessageId) (result ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/hostedContents", id.ID()),
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
		Values *[]models.ChatMessageHostedContentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContentsComplete retrieves all the results into a single object
func (c MeJoinedTeamPrimaryChannelMessageHostedContentClient) ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContentsComplete(ctx context.Context, id MeJoinedTeamPrimaryChannelMessageId) (ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContentsCompleteResult, error) {
	return c.ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContentsCompleteMatchingPredicate(ctx, id, models.ChatMessageHostedContentCollectionResponseOperationPredicate{})
}

// ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeJoinedTeamPrimaryChannelMessageHostedContentClient) ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContentsCompleteMatchingPredicate(ctx context.Context, id MeJoinedTeamPrimaryChannelMessageId, predicate models.ChatMessageHostedContentCollectionResponseOperationPredicate) (result ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContentsCompleteResult, err error) {
	items := make([]models.ChatMessageHostedContentCollectionResponse, 0)

	resp, err := c.ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContents(ctx, id)
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

	result = ListMeJoinedTeamByIdPrimaryChannelMessageByIdHostedContentsCompleteResult{
		Items: items,
	}
	return
}
