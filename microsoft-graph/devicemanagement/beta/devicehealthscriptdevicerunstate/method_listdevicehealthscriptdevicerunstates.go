package devicehealthscriptdevicerunstate

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

type ListDeviceHealthScriptDeviceRunStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceHealthScriptDeviceState
}

type ListDeviceHealthScriptDeviceRunStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceHealthScriptDeviceState
}

type ListDeviceHealthScriptDeviceRunStatesOperationOptions struct {
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

func DefaultListDeviceHealthScriptDeviceRunStatesOperationOptions() ListDeviceHealthScriptDeviceRunStatesOperationOptions {
	return ListDeviceHealthScriptDeviceRunStatesOperationOptions{}
}

func (o ListDeviceHealthScriptDeviceRunStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceHealthScriptDeviceRunStatesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceHealthScriptDeviceRunStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceHealthScriptDeviceRunStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceHealthScriptDeviceRunStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceHealthScriptDeviceRunStates - Get deviceRunStates from deviceManagement. List of run states for the device
// health script across all devices
func (c DeviceHealthScriptDeviceRunStateClient) ListDeviceHealthScriptDeviceRunStates(ctx context.Context, id beta.DeviceManagementDeviceHealthScriptId, options ListDeviceHealthScriptDeviceRunStatesOperationOptions) (result ListDeviceHealthScriptDeviceRunStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceHealthScriptDeviceRunStatesCustomPager{},
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
		Values *[]beta.DeviceHealthScriptDeviceState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceHealthScriptDeviceRunStatesComplete retrieves all the results into a single object
func (c DeviceHealthScriptDeviceRunStateClient) ListDeviceHealthScriptDeviceRunStatesComplete(ctx context.Context, id beta.DeviceManagementDeviceHealthScriptId, options ListDeviceHealthScriptDeviceRunStatesOperationOptions) (ListDeviceHealthScriptDeviceRunStatesCompleteResult, error) {
	return c.ListDeviceHealthScriptDeviceRunStatesCompleteMatchingPredicate(ctx, id, options, DeviceHealthScriptDeviceStateOperationPredicate{})
}

// ListDeviceHealthScriptDeviceRunStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceHealthScriptDeviceRunStateClient) ListDeviceHealthScriptDeviceRunStatesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDeviceHealthScriptId, options ListDeviceHealthScriptDeviceRunStatesOperationOptions, predicate DeviceHealthScriptDeviceStateOperationPredicate) (result ListDeviceHealthScriptDeviceRunStatesCompleteResult, err error) {
	items := make([]beta.DeviceHealthScriptDeviceState, 0)

	resp, err := c.ListDeviceHealthScriptDeviceRunStates(ctx, id, options)
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

	result = ListDeviceHealthScriptDeviceRunStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
