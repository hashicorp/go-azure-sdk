package processinfooperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListProcessThreadsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProcessThreadInfo
}

type WebAppsListProcessThreadsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProcessThreadInfo
}

type WebAppsListProcessThreadsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListProcessThreadsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListProcessThreads ...
func (c ProcessInfoOperationGroupClient) WebAppsListProcessThreads(ctx context.Context, id ProcessId) (result WebAppsListProcessThreadsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListProcessThreadsCustomPager{},
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
		Values *[]ProcessThreadInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListProcessThreadsComplete retrieves all the results into a single object
func (c ProcessInfoOperationGroupClient) WebAppsListProcessThreadsComplete(ctx context.Context, id ProcessId) (WebAppsListProcessThreadsCompleteResult, error) {
	return c.WebAppsListProcessThreadsCompleteMatchingPredicate(ctx, id, ProcessThreadInfoOperationPredicate{})
}

// WebAppsListProcessThreadsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProcessInfoOperationGroupClient) WebAppsListProcessThreadsCompleteMatchingPredicate(ctx context.Context, id ProcessId, predicate ProcessThreadInfoOperationPredicate) (result WebAppsListProcessThreadsCompleteResult, err error) {
	items := make([]ProcessThreadInfo, 0)

	resp, err := c.WebAppsListProcessThreads(ctx, id)
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

	result = WebAppsListProcessThreadsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
