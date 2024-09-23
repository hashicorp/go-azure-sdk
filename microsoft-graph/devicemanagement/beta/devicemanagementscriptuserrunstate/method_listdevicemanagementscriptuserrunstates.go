package devicemanagementscriptuserrunstate

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

type ListDeviceManagementScriptUserRunStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptUserState
}

type ListDeviceManagementScriptUserRunStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptUserState
}

type ListDeviceManagementScriptUserRunStatesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListDeviceManagementScriptUserRunStatesOperationOptions() ListDeviceManagementScriptUserRunStatesOperationOptions {
	return ListDeviceManagementScriptUserRunStatesOperationOptions{}
}

func (o ListDeviceManagementScriptUserRunStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceManagementScriptUserRunStatesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceManagementScriptUserRunStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceManagementScriptUserRunStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementScriptUserRunStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementScriptUserRunStates - Get userRunStates from deviceManagement. List of run states for this script
// across all users.
func (c DeviceManagementScriptUserRunStateClient) ListDeviceManagementScriptUserRunStates(ctx context.Context, id beta.DeviceManagementDeviceManagementScriptId, options ListDeviceManagementScriptUserRunStatesOperationOptions) (result ListDeviceManagementScriptUserRunStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceManagementScriptUserRunStatesCustomPager{},
		Path:          fmt.Sprintf("%s/userRunStates", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.DeviceManagementScriptUserState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceManagementScriptUserRunStatesComplete retrieves all the results into a single object
func (c DeviceManagementScriptUserRunStateClient) ListDeviceManagementScriptUserRunStatesComplete(ctx context.Context, id beta.DeviceManagementDeviceManagementScriptId, options ListDeviceManagementScriptUserRunStatesOperationOptions) (ListDeviceManagementScriptUserRunStatesCompleteResult, error) {
	return c.ListDeviceManagementScriptUserRunStatesCompleteMatchingPredicate(ctx, id, options, DeviceManagementScriptUserStateOperationPredicate{})
}

// ListDeviceManagementScriptUserRunStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementScriptUserRunStateClient) ListDeviceManagementScriptUserRunStatesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDeviceManagementScriptId, options ListDeviceManagementScriptUserRunStatesOperationOptions, predicate DeviceManagementScriptUserStateOperationPredicate) (result ListDeviceManagementScriptUserRunStatesCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptUserState, 0)

	resp, err := c.ListDeviceManagementScriptUserRunStates(ctx, id, options)
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

	result = ListDeviceManagementScriptUserRunStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
