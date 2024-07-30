package thread

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

type ListThreadsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ConversationThread
}

type ListThreadsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ConversationThread
}

type ListThreadsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListThreadsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListThreads ...
func (c ThreadClient) ListThreads(ctx context.Context, id GroupId) (result ListThreadsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListThreadsCustomPager{},
		Path:       fmt.Sprintf("%s/threads", id.ID()),
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
		Values *[]beta.ConversationThread `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListThreadsComplete retrieves all the results into a single object
func (c ThreadClient) ListThreadsComplete(ctx context.Context, id GroupId) (ListThreadsCompleteResult, error) {
	return c.ListThreadsCompleteMatchingPredicate(ctx, id, ConversationThreadOperationPredicate{})
}

// ListThreadsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ThreadClient) ListThreadsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate ConversationThreadOperationPredicate) (result ListThreadsCompleteResult, err error) {
	items := make([]beta.ConversationThread, 0)

	resp, err := c.ListThreads(ctx, id)
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

	result = ListThreadsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
