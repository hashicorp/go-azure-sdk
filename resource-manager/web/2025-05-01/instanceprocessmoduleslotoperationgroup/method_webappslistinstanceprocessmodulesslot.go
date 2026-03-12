package instanceprocessmoduleslotoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListInstanceProcessModulesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProcessModuleInfo
}

type WebAppsListInstanceProcessModulesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProcessModuleInfo
}

type WebAppsListInstanceProcessModulesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListInstanceProcessModulesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListInstanceProcessModulesSlot ...
func (c InstanceProcessModuleSlotOperationGroupClient) WebAppsListInstanceProcessModulesSlot(ctx context.Context, id SlotInstanceProcessId) (result WebAppsListInstanceProcessModulesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListInstanceProcessModulesSlotCustomPager{},
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

// WebAppsListInstanceProcessModulesSlotComplete retrieves all the results into a single object
func (c InstanceProcessModuleSlotOperationGroupClient) WebAppsListInstanceProcessModulesSlotComplete(ctx context.Context, id SlotInstanceProcessId) (WebAppsListInstanceProcessModulesSlotCompleteResult, error) {
	return c.WebAppsListInstanceProcessModulesSlotCompleteMatchingPredicate(ctx, id, ProcessModuleInfoOperationPredicate{})
}

// WebAppsListInstanceProcessModulesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InstanceProcessModuleSlotOperationGroupClient) WebAppsListInstanceProcessModulesSlotCompleteMatchingPredicate(ctx context.Context, id SlotInstanceProcessId, predicate ProcessModuleInfoOperationPredicate) (result WebAppsListInstanceProcessModulesSlotCompleteResult, err error) {
	items := make([]ProcessModuleInfo, 0)

	resp, err := c.WebAppsListInstanceProcessModulesSlot(ctx, id)
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

	result = WebAppsListInstanceProcessModulesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
