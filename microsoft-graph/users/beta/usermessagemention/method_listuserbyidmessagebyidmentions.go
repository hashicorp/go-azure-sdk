package usermessagemention

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

type ListUserByIdMessageByIdMentionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MentionCollectionResponse
}

type ListUserByIdMessageByIdMentionsCompleteResult struct {
	Items []models.MentionCollectionResponse
}

// ListUserByIdMessageByIdMentions ...
func (c UserMessageMentionClient) ListUserByIdMessageByIdMentions(ctx context.Context, id UserMessageId) (result ListUserByIdMessageByIdMentionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/mentions", id.ID()),
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
		Values *[]models.MentionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdMessageByIdMentionsComplete retrieves all the results into a single object
func (c UserMessageMentionClient) ListUserByIdMessageByIdMentionsComplete(ctx context.Context, id UserMessageId) (ListUserByIdMessageByIdMentionsCompleteResult, error) {
	return c.ListUserByIdMessageByIdMentionsCompleteMatchingPredicate(ctx, id, models.MentionCollectionResponseOperationPredicate{})
}

// ListUserByIdMessageByIdMentionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserMessageMentionClient) ListUserByIdMessageByIdMentionsCompleteMatchingPredicate(ctx context.Context, id UserMessageId, predicate models.MentionCollectionResponseOperationPredicate) (result ListUserByIdMessageByIdMentionsCompleteResult, err error) {
	items := make([]models.MentionCollectionResponse, 0)

	resp, err := c.ListUserByIdMessageByIdMentions(ctx, id)
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

	result = ListUserByIdMessageByIdMentionsCompleteResult{
		Items: items,
	}
	return
}
