package deviceshellscriptuserrunstatedevicerunstate

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

type ListDeviceShellScriptUserRunStateDeviceRunStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptDeviceState
}

type ListDeviceShellScriptUserRunStateDeviceRunStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptDeviceState
}

type ListDeviceShellScriptUserRunStateDeviceRunStatesOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListDeviceShellScriptUserRunStateDeviceRunStatesOperationOptions() ListDeviceShellScriptUserRunStateDeviceRunStatesOperationOptions {
	return ListDeviceShellScriptUserRunStateDeviceRunStatesOperationOptions{}
}

func (o ListDeviceShellScriptUserRunStateDeviceRunStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceShellScriptUserRunStateDeviceRunStatesOperationOptions) ToOData() *odata.Query {
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
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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

func (o ListDeviceShellScriptUserRunStateDeviceRunStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceShellScriptUserRunStateDeviceRunStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceShellScriptUserRunStateDeviceRunStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceShellScriptUserRunStateDeviceRunStates - Get deviceRunStates from deviceManagement. List of run states for
// this script across all devices of specific user.
func (c DeviceShellScriptUserRunStateDeviceRunStateClient) ListDeviceShellScriptUserRunStateDeviceRunStates(ctx context.Context, id beta.DeviceManagementDeviceShellScriptIdUserRunStateId, options ListDeviceShellScriptUserRunStateDeviceRunStatesOperationOptions) (result ListDeviceShellScriptUserRunStateDeviceRunStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceShellScriptUserRunStateDeviceRunStatesCustomPager{},
		Path:          fmt.Sprintf("%s/deviceRunStates", id.ID()),
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
		Values *[]beta.DeviceManagementScriptDeviceState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceShellScriptUserRunStateDeviceRunStatesComplete retrieves all the results into a single object
func (c DeviceShellScriptUserRunStateDeviceRunStateClient) ListDeviceShellScriptUserRunStateDeviceRunStatesComplete(ctx context.Context, id beta.DeviceManagementDeviceShellScriptIdUserRunStateId, options ListDeviceShellScriptUserRunStateDeviceRunStatesOperationOptions) (ListDeviceShellScriptUserRunStateDeviceRunStatesCompleteResult, error) {
	return c.ListDeviceShellScriptUserRunStateDeviceRunStatesCompleteMatchingPredicate(ctx, id, options, DeviceManagementScriptDeviceStateOperationPredicate{})
}

// ListDeviceShellScriptUserRunStateDeviceRunStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceShellScriptUserRunStateDeviceRunStateClient) ListDeviceShellScriptUserRunStateDeviceRunStatesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDeviceShellScriptIdUserRunStateId, options ListDeviceShellScriptUserRunStateDeviceRunStatesOperationOptions, predicate DeviceManagementScriptDeviceStateOperationPredicate) (result ListDeviceShellScriptUserRunStateDeviceRunStatesCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptDeviceState, 0)

	resp, err := c.ListDeviceShellScriptUserRunStateDeviceRunStates(ctx, id, options)
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

	result = ListDeviceShellScriptUserRunStateDeviceRunStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
