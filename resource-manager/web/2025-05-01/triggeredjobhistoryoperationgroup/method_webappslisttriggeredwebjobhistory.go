package triggeredjobhistoryoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListTriggeredWebJobHistoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TriggeredJobHistory
}

type WebAppsListTriggeredWebJobHistoryCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TriggeredJobHistory
}

type WebAppsListTriggeredWebJobHistoryCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListTriggeredWebJobHistoryCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListTriggeredWebJobHistory ...
func (c TriggeredJobHistoryOperationGroupClient) WebAppsListTriggeredWebJobHistory(ctx context.Context, id TriggeredWebJobId) (result WebAppsListTriggeredWebJobHistoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListTriggeredWebJobHistoryCustomPager{},
		Path:       fmt.Sprintf("%s/history", id.ID()),
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
		Values *[]TriggeredJobHistory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListTriggeredWebJobHistoryComplete retrieves all the results into a single object
func (c TriggeredJobHistoryOperationGroupClient) WebAppsListTriggeredWebJobHistoryComplete(ctx context.Context, id TriggeredWebJobId) (WebAppsListTriggeredWebJobHistoryCompleteResult, error) {
	return c.WebAppsListTriggeredWebJobHistoryCompleteMatchingPredicate(ctx, id, TriggeredJobHistoryOperationPredicate{})
}

// WebAppsListTriggeredWebJobHistoryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TriggeredJobHistoryOperationGroupClient) WebAppsListTriggeredWebJobHistoryCompleteMatchingPredicate(ctx context.Context, id TriggeredWebJobId, predicate TriggeredJobHistoryOperationPredicate) (result WebAppsListTriggeredWebJobHistoryCompleteResult, err error) {
	items := make([]TriggeredJobHistory, 0)

	resp, err := c.WebAppsListTriggeredWebJobHistory(ctx, id)
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

	result = WebAppsListTriggeredWebJobHistoryCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
