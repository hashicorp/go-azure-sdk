package manageddevicesecuritybaselinestate

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

type ListManagedDeviceSecurityBaselineStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SecurityBaselineState
}

type ListManagedDeviceSecurityBaselineStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SecurityBaselineState
}

type ListManagedDeviceSecurityBaselineStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceSecurityBaselineStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceSecurityBaselineStates ...
func (c ManagedDeviceSecurityBaselineStateClient) ListManagedDeviceSecurityBaselineStates(ctx context.Context, id UserIdManagedDeviceId) (result ListManagedDeviceSecurityBaselineStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListManagedDeviceSecurityBaselineStatesCustomPager{},
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

// ListManagedDeviceSecurityBaselineStatesComplete retrieves all the results into a single object
func (c ManagedDeviceSecurityBaselineStateClient) ListManagedDeviceSecurityBaselineStatesComplete(ctx context.Context, id UserIdManagedDeviceId) (ListManagedDeviceSecurityBaselineStatesCompleteResult, error) {
	return c.ListManagedDeviceSecurityBaselineStatesCompleteMatchingPredicate(ctx, id, SecurityBaselineStateOperationPredicate{})
}

// ListManagedDeviceSecurityBaselineStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceSecurityBaselineStateClient) ListManagedDeviceSecurityBaselineStatesCompleteMatchingPredicate(ctx context.Context, id UserIdManagedDeviceId, predicate SecurityBaselineStateOperationPredicate) (result ListManagedDeviceSecurityBaselineStatesCompleteResult, err error) {
	items := make([]beta.SecurityBaselineState, 0)

	resp, err := c.ListManagedDeviceSecurityBaselineStates(ctx, id)
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

	result = ListManagedDeviceSecurityBaselineStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
