package manageddevicesecuritybaselinestatesettingstate

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

type ListManagedDeviceSecurityBaselineStateSettingStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SecurityBaselineSettingState
}

type ListManagedDeviceSecurityBaselineStateSettingStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SecurityBaselineSettingState
}

type ListManagedDeviceSecurityBaselineStateSettingStatesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListManagedDeviceSecurityBaselineStateSettingStatesOperationOptions() ListManagedDeviceSecurityBaselineStateSettingStatesOperationOptions {
	return ListManagedDeviceSecurityBaselineStateSettingStatesOperationOptions{}
}

func (o ListManagedDeviceSecurityBaselineStateSettingStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListManagedDeviceSecurityBaselineStateSettingStatesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListManagedDeviceSecurityBaselineStateSettingStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListManagedDeviceSecurityBaselineStateSettingStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceSecurityBaselineStateSettingStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceSecurityBaselineStateSettingStates - Get settingStates from me. The security baseline state for
// different settings for a device
func (c ManagedDeviceSecurityBaselineStateSettingStateClient) ListManagedDeviceSecurityBaselineStateSettingStates(ctx context.Context, id beta.MeManagedDeviceIdSecurityBaselineStateId, options ListManagedDeviceSecurityBaselineStateSettingStatesOperationOptions) (result ListManagedDeviceSecurityBaselineStateSettingStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListManagedDeviceSecurityBaselineStateSettingStatesCustomPager{},
		Path:          fmt.Sprintf("%s/settingStates", id.ID()),
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

// ListManagedDeviceSecurityBaselineStateSettingStatesComplete retrieves all the results into a single object
func (c ManagedDeviceSecurityBaselineStateSettingStateClient) ListManagedDeviceSecurityBaselineStateSettingStatesComplete(ctx context.Context, id beta.MeManagedDeviceIdSecurityBaselineStateId, options ListManagedDeviceSecurityBaselineStateSettingStatesOperationOptions) (ListManagedDeviceSecurityBaselineStateSettingStatesCompleteResult, error) {
	return c.ListManagedDeviceSecurityBaselineStateSettingStatesCompleteMatchingPredicate(ctx, id, options, SecurityBaselineSettingStateOperationPredicate{})
}

// ListManagedDeviceSecurityBaselineStateSettingStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceSecurityBaselineStateSettingStateClient) ListManagedDeviceSecurityBaselineStateSettingStatesCompleteMatchingPredicate(ctx context.Context, id beta.MeManagedDeviceIdSecurityBaselineStateId, options ListManagedDeviceSecurityBaselineStateSettingStatesOperationOptions, predicate SecurityBaselineSettingStateOperationPredicate) (result ListManagedDeviceSecurityBaselineStateSettingStatesCompleteResult, err error) {
	items := make([]beta.SecurityBaselineSettingState, 0)

	resp, err := c.ListManagedDeviceSecurityBaselineStateSettingStates(ctx, id, options)
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

	result = ListManagedDeviceSecurityBaselineStateSettingStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
