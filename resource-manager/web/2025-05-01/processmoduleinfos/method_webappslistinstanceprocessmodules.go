package processmoduleinfos

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListInstanceProcessModulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProcessModuleInfo
}

type WebAppsListInstanceProcessModulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProcessModuleInfo
}

type WebAppsListInstanceProcessModulesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListInstanceProcessModulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListInstanceProcessModules ...
func (c ProcessModuleInfosClient) WebAppsListInstanceProcessModules(ctx context.Context, id InstanceProcessId) (result WebAppsListInstanceProcessModulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListInstanceProcessModulesCustomPager{},
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

// WebAppsListInstanceProcessModulesComplete retrieves all the results into a single object
func (c ProcessModuleInfosClient) WebAppsListInstanceProcessModulesComplete(ctx context.Context, id InstanceProcessId) (WebAppsListInstanceProcessModulesCompleteResult, error) {
	return c.WebAppsListInstanceProcessModulesCompleteMatchingPredicate(ctx, id, ProcessModuleInfoOperationPredicate{})
}

// WebAppsListInstanceProcessModulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProcessModuleInfosClient) WebAppsListInstanceProcessModulesCompleteMatchingPredicate(ctx context.Context, id InstanceProcessId, predicate ProcessModuleInfoOperationPredicate) (result WebAppsListInstanceProcessModulesCompleteResult, err error) {
	items := make([]ProcessModuleInfo, 0)

	resp, err := c.WebAppsListInstanceProcessModules(ctx, id)
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

	result = WebAppsListInstanceProcessModulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
