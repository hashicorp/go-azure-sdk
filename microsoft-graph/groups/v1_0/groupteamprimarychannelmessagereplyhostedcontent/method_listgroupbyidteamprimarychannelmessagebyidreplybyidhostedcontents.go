package groupteamprimarychannelmessagereplyhostedcontent

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

type ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChatMessageHostedContentCollectionResponse
}

type ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContentsCompleteResult struct {
	Items []models.ChatMessageHostedContentCollectionResponse
}

// ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContents ...
func (c GroupTeamPrimaryChannelMessageReplyHostedContentClient) ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContents(ctx context.Context, id GroupTeamPrimaryChannelMessageReplyId) (result ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContentsOperationResponse, err error) {
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

// ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContentsComplete retrieves all the results into a single object
func (c GroupTeamPrimaryChannelMessageReplyHostedContentClient) ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContentsComplete(ctx context.Context, id GroupTeamPrimaryChannelMessageReplyId) (ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContentsCompleteResult, error) {
	return c.ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContentsCompleteMatchingPredicate(ctx, id, models.ChatMessageHostedContentCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamPrimaryChannelMessageReplyHostedContentClient) ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContentsCompleteMatchingPredicate(ctx context.Context, id GroupTeamPrimaryChannelMessageReplyId, predicate models.ChatMessageHostedContentCollectionResponseOperationPredicate) (result ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContentsCompleteResult, err error) {
	items := make([]models.ChatMessageHostedContentCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContents(ctx, id)
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

	result = ListGroupByIdTeamPrimaryChannelMessageByIdReplyByIdHostedContentsCompleteResult{
		Items: items,
	}
	return
}
