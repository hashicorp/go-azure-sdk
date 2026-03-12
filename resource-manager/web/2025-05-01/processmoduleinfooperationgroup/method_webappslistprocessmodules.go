package processmoduleinfooperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListProcessModulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProcessModuleInfo
}

type WebAppsListProcessModulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProcessModuleInfo
}

type WebAppsListProcessModulesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListProcessModulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListProcessModules ...
func (c ProcessModuleInfoOperationGroupClient) WebAppsListProcessModules(ctx context.Context, id ProcessId) (result WebAppsListProcessModulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListProcessModulesCustomPager{},
		Path:       fmt.Sprintf("%s/modules", id.ID()),
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
		Values *[]ProcessModuleInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListProcessModulesComplete retrieves all the results into a single object
func (c ProcessModuleInfoOperationGroupClient) WebAppsListProcessModulesComplete(ctx context.Context, id ProcessId) (WebAppsListProcessModulesCompleteResult, error) {
	return c.WebAppsListProcessModulesCompleteMatchingPredicate(ctx, id, ProcessModuleInfoOperationPredicate{})
}

// WebAppsListProcessModulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProcessModuleInfoOperationGroupClient) WebAppsListProcessModulesCompleteMatchingPredicate(ctx context.Context, id ProcessId, predicate ProcessModuleInfoOperationPredicate) (result WebAppsListProcessModulesCompleteResult, err error) {
	items := make([]ProcessModuleInfo, 0)

	resp, err := c.WebAppsListProcessModules(ctx, id)
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

	result = WebAppsListProcessModulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
