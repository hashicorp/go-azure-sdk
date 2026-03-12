package processmoduleslotoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListProcessModulesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProcessModuleInfo
}

type WebAppsListProcessModulesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProcessModuleInfo
}

type WebAppsListProcessModulesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListProcessModulesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListProcessModulesSlot ...
func (c ProcessModuleSlotOperationGroupClient) WebAppsListProcessModulesSlot(ctx context.Context, id SlotProcessId) (result WebAppsListProcessModulesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListProcessModulesSlotCustomPager{},
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

// WebAppsListProcessModulesSlotComplete retrieves all the results into a single object
func (c ProcessModuleSlotOperationGroupClient) WebAppsListProcessModulesSlotComplete(ctx context.Context, id SlotProcessId) (WebAppsListProcessModulesSlotCompleteResult, error) {
	return c.WebAppsListProcessModulesSlotCompleteMatchingPredicate(ctx, id, ProcessModuleInfoOperationPredicate{})
}

// WebAppsListProcessModulesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProcessModuleSlotOperationGroupClient) WebAppsListProcessModulesSlotCompleteMatchingPredicate(ctx context.Context, id SlotProcessId, predicate ProcessModuleInfoOperationPredicate) (result WebAppsListProcessModulesSlotCompleteResult, err error) {
	items := make([]ProcessModuleInfo, 0)

	resp, err := c.WebAppsListProcessModulesSlot(ctx, id)
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

	result = WebAppsListProcessModulesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
