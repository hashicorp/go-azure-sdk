package comanageddevicesecuritybaselinestate

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

type ListComanagedDeviceSecurityBaselineStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SecurityBaselineState
}

type ListComanagedDeviceSecurityBaselineStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SecurityBaselineState
}

type ListComanagedDeviceSecurityBaselineStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDeviceSecurityBaselineStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDeviceSecurityBaselineStates ...
func (c ComanagedDeviceSecurityBaselineStateClient) ListComanagedDeviceSecurityBaselineStates(ctx context.Context, id DeviceManagementComanagedDeviceId) (result ListComanagedDeviceSecurityBaselineStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListComanagedDeviceSecurityBaselineStatesCustomPager{},
		Path:       fmt.Sprintf("%s/securityBaselineStates", id.ID()),
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
		Values *[]beta.SecurityBaselineState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagedDeviceSecurityBaselineStatesComplete retrieves all the results into a single object
func (c ComanagedDeviceSecurityBaselineStateClient) ListComanagedDeviceSecurityBaselineStatesComplete(ctx context.Context, id DeviceManagementComanagedDeviceId) (ListComanagedDeviceSecurityBaselineStatesCompleteResult, error) {
	return c.ListComanagedDeviceSecurityBaselineStatesCompleteMatchingPredicate(ctx, id, SecurityBaselineStateOperationPredicate{})
}

// ListComanagedDeviceSecurityBaselineStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceSecurityBaselineStateClient) ListComanagedDeviceSecurityBaselineStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementComanagedDeviceId, predicate SecurityBaselineStateOperationPredicate) (result ListComanagedDeviceSecurityBaselineStatesCompleteResult, err error) {
	items := make([]beta.SecurityBaselineState, 0)

	resp, err := c.ListComanagedDeviceSecurityBaselineStates(ctx, id)
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

	result = ListComanagedDeviceSecurityBaselineStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
