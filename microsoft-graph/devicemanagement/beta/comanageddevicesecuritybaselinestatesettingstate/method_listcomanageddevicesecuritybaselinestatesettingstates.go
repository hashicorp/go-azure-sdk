package comanageddevicesecuritybaselinestatesettingstate

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

type ListComanagedDeviceSecurityBaselineStateSettingStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SecurityBaselineSettingState
}

type ListComanagedDeviceSecurityBaselineStateSettingStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SecurityBaselineSettingState
}

type ListComanagedDeviceSecurityBaselineStateSettingStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDeviceSecurityBaselineStateSettingStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDeviceSecurityBaselineStateSettingStates ...
func (c ComanagedDeviceSecurityBaselineStateSettingStateClient) ListComanagedDeviceSecurityBaselineStateSettingStates(ctx context.Context, id DeviceManagementComanagedDeviceIdSecurityBaselineStateId) (result ListComanagedDeviceSecurityBaselineStateSettingStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListComanagedDeviceSecurityBaselineStateSettingStatesCustomPager{},
		Path:       fmt.Sprintf("%s/settingStates", id.ID()),
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
		Values *[]beta.SecurityBaselineSettingState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagedDeviceSecurityBaselineStateSettingStatesComplete retrieves all the results into a single object
func (c ComanagedDeviceSecurityBaselineStateSettingStateClient) ListComanagedDeviceSecurityBaselineStateSettingStatesComplete(ctx context.Context, id DeviceManagementComanagedDeviceIdSecurityBaselineStateId) (ListComanagedDeviceSecurityBaselineStateSettingStatesCompleteResult, error) {
	return c.ListComanagedDeviceSecurityBaselineStateSettingStatesCompleteMatchingPredicate(ctx, id, SecurityBaselineSettingStateOperationPredicate{})
}

// ListComanagedDeviceSecurityBaselineStateSettingStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceSecurityBaselineStateSettingStateClient) ListComanagedDeviceSecurityBaselineStateSettingStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementComanagedDeviceIdSecurityBaselineStateId, predicate SecurityBaselineSettingStateOperationPredicate) (result ListComanagedDeviceSecurityBaselineStateSettingStatesCompleteResult, err error) {
	items := make([]beta.SecurityBaselineSettingState, 0)

	resp, err := c.ListComanagedDeviceSecurityBaselineStateSettingStates(ctx, id)
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

	result = ListComanagedDeviceSecurityBaselineStateSettingStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
