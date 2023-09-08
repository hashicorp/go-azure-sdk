package memessagemention

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

type ListMeMessageByIdMentionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MentionCollectionResponse
}

type ListMeMessageByIdMentionsCompleteResult struct {
	Items []models.MentionCollectionResponse
}

// ListMeMessageByIdMentions ...
func (c MeMessageMentionClient) ListMeMessageByIdMentions(ctx context.Context, id MeMessageId) (result ListMeMessageByIdMentionsOperationResponse, err error) {
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

// ListMeMessageByIdMentionsComplete retrieves all the results into a single object
func (c MeMessageMentionClient) ListMeMessageByIdMentionsComplete(ctx context.Context, id MeMessageId) (ListMeMessageByIdMentionsCompleteResult, error) {
	return c.ListMeMessageByIdMentionsCompleteMatchingPredicate(ctx, id, models.MentionCollectionResponseOperationPredicate{})
}

// ListMeMessageByIdMentionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeMessageMentionClient) ListMeMessageByIdMentionsCompleteMatchingPredicate(ctx context.Context, id MeMessageId, predicate models.MentionCollectionResponseOperationPredicate) (result ListMeMessageByIdMentionsCompleteResult, err error) {
	items := make([]models.MentionCollectionResponse, 0)

	resp, err := c.ListMeMessageByIdMentions(ctx, id)
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

	result = ListMeMessageByIdMentionsCompleteResult{
		Items: items,
	}
	return
}
