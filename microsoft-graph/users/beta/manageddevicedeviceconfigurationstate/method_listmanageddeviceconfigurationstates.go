package manageddevicedeviceconfigurationstate

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

type ListManagedDeviceConfigurationStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceConfigurationState
}

type ListManagedDeviceConfigurationStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceConfigurationState
}

type ListManagedDeviceConfigurationStatesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListManagedDeviceConfigurationStatesOperationOptions() ListManagedDeviceConfigurationStatesOperationOptions {
	return ListManagedDeviceConfigurationStatesOperationOptions{}
}

func (o ListManagedDeviceConfigurationStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListManagedDeviceConfigurationStatesOperationOptions) ToOData() *odata.Query {
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

func (o ListManagedDeviceConfigurationStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListManagedDeviceConfigurationStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceConfigurationStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceConfigurationStates - Get deviceConfigurationStates from users. Device configuration states for this
// device.
func (c ManagedDeviceDeviceConfigurationStateClient) ListManagedDeviceConfigurationStates(ctx context.Context, id beta.UserIdManagedDeviceId, options ListManagedDeviceConfigurationStatesOperationOptions) (result ListManagedDeviceConfigurationStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListManagedDeviceConfigurationStatesCustomPager{},
		Path:          fmt.Sprintf("%s/deviceConfigurationStates", id.ID()),
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
		Values *[]beta.DeviceConfigurationState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedDeviceConfigurationStatesComplete retrieves all the results into a single object
func (c ManagedDeviceDeviceConfigurationStateClient) ListManagedDeviceConfigurationStatesComplete(ctx context.Context, id beta.UserIdManagedDeviceId, options ListManagedDeviceConfigurationStatesOperationOptions) (ListManagedDeviceConfigurationStatesCompleteResult, error) {
	return c.ListManagedDeviceConfigurationStatesCompleteMatchingPredicate(ctx, id, options, DeviceConfigurationStateOperationPredicate{})
}

// ListManagedDeviceConfigurationStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceDeviceConfigurationStateClient) ListManagedDeviceConfigurationStatesCompleteMatchingPredicate(ctx context.Context, id beta.UserIdManagedDeviceId, options ListManagedDeviceConfigurationStatesOperationOptions, predicate DeviceConfigurationStateOperationPredicate) (result ListManagedDeviceConfigurationStatesCompleteResult, err error) {
	items := make([]beta.DeviceConfigurationState, 0)

	resp, err := c.ListManagedDeviceConfigurationStates(ctx, id, options)
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

	result = ListManagedDeviceConfigurationStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
