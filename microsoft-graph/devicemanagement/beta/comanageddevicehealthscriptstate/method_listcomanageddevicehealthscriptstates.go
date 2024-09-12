package comanageddevicehealthscriptstate

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

type ListComanagedDeviceHealthScriptStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceHealthScriptPolicyState
}

type ListComanagedDeviceHealthScriptStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceHealthScriptPolicyState
}

type ListComanagedDeviceHealthScriptStatesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListComanagedDeviceHealthScriptStatesOperationOptions() ListComanagedDeviceHealthScriptStatesOperationOptions {
	return ListComanagedDeviceHealthScriptStatesOperationOptions{}
}

func (o ListComanagedDeviceHealthScriptStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListComanagedDeviceHealthScriptStatesOperationOptions) ToOData() *odata.Query {
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

func (o ListComanagedDeviceHealthScriptStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListComanagedDeviceHealthScriptStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDeviceHealthScriptStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDeviceHealthScriptStates - Get deviceHealthScriptStates from deviceManagement. Results of device health
// scripts that ran for this device. Default is empty list. This property is read-only.
func (c ComanagedDeviceHealthScriptStateClient) ListComanagedDeviceHealthScriptStates(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceHealthScriptStatesOperationOptions) (result ListComanagedDeviceHealthScriptStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListComanagedDeviceHealthScriptStatesCustomPager{},
		Path:          fmt.Sprintf("%s/deviceHealthScriptStates", id.ID()),
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
		Values *[]beta.DeviceHealthScriptPolicyState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagedDeviceHealthScriptStatesComplete retrieves all the results into a single object
func (c ComanagedDeviceHealthScriptStateClient) ListComanagedDeviceHealthScriptStatesComplete(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceHealthScriptStatesOperationOptions) (ListComanagedDeviceHealthScriptStatesCompleteResult, error) {
	return c.ListComanagedDeviceHealthScriptStatesCompleteMatchingPredicate(ctx, id, options, DeviceHealthScriptPolicyStateOperationPredicate{})
}

// ListComanagedDeviceHealthScriptStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceHealthScriptStateClient) ListComanagedDeviceHealthScriptStatesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceHealthScriptStatesOperationOptions, predicate DeviceHealthScriptPolicyStateOperationPredicate) (result ListComanagedDeviceHealthScriptStatesCompleteResult, err error) {
	items := make([]beta.DeviceHealthScriptPolicyState, 0)

	resp, err := c.ListComanagedDeviceHealthScriptStates(ctx, id, options)
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

	result = ListComanagedDeviceHealthScriptStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
