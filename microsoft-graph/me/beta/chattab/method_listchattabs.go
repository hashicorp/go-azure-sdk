package chattab

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListChatTabsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TeamsTab
}

type ListChatTabsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TeamsTab
}

type ListChatTabsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListChatTabsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListChatTabs ...
func (c ChatTabClient) ListChatTabs(ctx context.Context, id MeChatId) (result ListChatTabsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListChatTabsCustomPager{},
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
		Values *[]beta.TeamsTab `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListChatTabsComplete retrieves all the results into a single object
func (c ChatTabClient) ListChatTabsComplete(ctx context.Context, id MeChatId) (ListChatTabsCompleteResult, error) {
	return c.ListChatTabsCompleteMatchingPredicate(ctx, id, TeamsTabOperationPredicate{})
}

// ListChatTabsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatTabClient) ListChatTabsCompleteMatchingPredicate(ctx context.Context, id MeChatId, predicate TeamsTabOperationPredicate) (result ListChatTabsCompleteResult, err error) {
	items := make([]beta.TeamsTab, 0)

	resp, err := c.ListChatTabs(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListChatTabsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
