package groupteamprimarychannelmessagehostedcontent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdTeamPrimaryChannelMessageByIdHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ChatMessageHostedContentCollectionResponse
}

type ListGroupByIdTeamPrimaryChannelMessageByIdHostedContentsCompleteResult struct {
	Items []models.ChatMessageHostedContentCollectionResponse
}

// ListGroupByIdTeamPrimaryChannelMessageByIdHostedContents ...
func (c GroupTeamPrimaryChannelMessageHostedContentClient) ListGroupByIdTeamPrimaryChannelMessageByIdHostedContents(ctx context.Context, id GroupTeamPrimaryChannelMessageId) (result ListGroupByIdTeamPrimaryChannelMessageByIdHostedContentsOperationResponse, err error) {
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

// ListGroupByIdTeamPrimaryChannelMessageByIdHostedContentsComplete retrieves all the results into a single object
func (c GroupTeamPrimaryChannelMessageHostedContentClient) ListGroupByIdTeamPrimaryChannelMessageByIdHostedContentsComplete(ctx context.Context, id GroupTeamPrimaryChannelMessageId) (ListGroupByIdTeamPrimaryChannelMessageByIdHostedContentsCompleteResult, error) {
	return c.ListGroupByIdTeamPrimaryChannelMessageByIdHostedContentsCompleteMatchingPredicate(ctx, id, models.ChatMessageHostedContentCollectionResponseOperationPredicate{})
}

// ListGroupByIdTeamPrimaryChannelMessageByIdHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTeamPrimaryChannelMessageHostedContentClient) ListGroupByIdTeamPrimaryChannelMessageByIdHostedContentsCompleteMatchingPredicate(ctx context.Context, id GroupTeamPrimaryChannelMessageId, predicate models.ChatMessageHostedContentCollectionResponseOperationPredicate) (result ListGroupByIdTeamPrimaryChannelMessageByIdHostedContentsCompleteResult, err error) {
	items := make([]models.ChatMessageHostedContentCollectionResponse, 0)

	resp, err := c.ListGroupByIdTeamPrimaryChannelMessageByIdHostedContents(ctx, id)
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

	result = ListGroupByIdTeamPrimaryChannelMessageByIdHostedContentsCompleteResult{
		Items: items,
	}
	return
}
