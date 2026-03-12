package processinfooperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListProcessesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProcessInfo
}

type WebAppsListProcessesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProcessInfo
}

type WebAppsListProcessesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListProcessesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListProcesses ...
func (c ProcessInfoOperationGroupClient) WebAppsListProcesses(ctx context.Context, id commonids.AppServiceId) (result WebAppsListProcessesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListProcessesCustomPager{},
		Path:       fmt.Sprintf("%s/processes", id.ID()),
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
		Values *[]ProcessInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListProcessesComplete retrieves all the results into a single object
func (c ProcessInfoOperationGroupClient) WebAppsListProcessesComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListProcessesCompleteResult, error) {
	return c.WebAppsListProcessesCompleteMatchingPredicate(ctx, id, ProcessInfoOperationPredicate{})
}

// WebAppsListProcessesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProcessInfoOperationGroupClient) WebAppsListProcessesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate ProcessInfoOperationPredicate) (result WebAppsListProcessesCompleteResult, err error) {
	items := make([]ProcessInfo, 0)

	resp, err := c.WebAppsListProcesses(ctx, id)
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

	result = WebAppsListProcessesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
