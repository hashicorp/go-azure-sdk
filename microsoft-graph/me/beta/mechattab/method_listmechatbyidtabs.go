package mechattab

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

type ListMeChatByIdTabsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TeamsTabCollectionResponse
}

type ListMeChatByIdTabsCompleteResult struct {
	Items []models.TeamsTabCollectionResponse
}

// ListMeChatByIdTabs ...
func (c MeChatTabClient) ListMeChatByIdTabs(ctx context.Context, id MeChatId) (result ListMeChatByIdTabsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/tabs", id.ID()),
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
		Values *[]models.TeamsTabCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeChatByIdTabsComplete retrieves all the results into a single object
func (c MeChatTabClient) ListMeChatByIdTabsComplete(ctx context.Context, id MeChatId) (ListMeChatByIdTabsCompleteResult, error) {
	return c.ListMeChatByIdTabsCompleteMatchingPredicate(ctx, id, models.TeamsTabCollectionResponseOperationPredicate{})
}

// ListMeChatByIdTabsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeChatTabClient) ListMeChatByIdTabsCompleteMatchingPredicate(ctx context.Context, id MeChatId, predicate models.TeamsTabCollectionResponseOperationPredicate) (result ListMeChatByIdTabsCompleteResult, err error) {
	items := make([]models.TeamsTabCollectionResponse, 0)

	resp, err := c.ListMeChatByIdTabs(ctx, id)
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

	result = ListMeChatByIdTabsCompleteResult{
		Items: items,
	}
	return
}
