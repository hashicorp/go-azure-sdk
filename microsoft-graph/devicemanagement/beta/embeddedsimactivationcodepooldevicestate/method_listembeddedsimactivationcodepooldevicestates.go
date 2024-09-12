package embeddedsimactivationcodepooldevicestate

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

type ListEmbeddedSIMActivationCodePoolDeviceStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.EmbeddedSIMDeviceState
}

type ListEmbeddedSIMActivationCodePoolDeviceStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.EmbeddedSIMDeviceState
}

type ListEmbeddedSIMActivationCodePoolDeviceStatesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEmbeddedSIMActivationCodePoolDeviceStatesOperationOptions() ListEmbeddedSIMActivationCodePoolDeviceStatesOperationOptions {
	return ListEmbeddedSIMActivationCodePoolDeviceStatesOperationOptions{}
}

func (o ListEmbeddedSIMActivationCodePoolDeviceStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEmbeddedSIMActivationCodePoolDeviceStatesOperationOptions) ToOData() *odata.Query {
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

func (o ListEmbeddedSIMActivationCodePoolDeviceStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEmbeddedSIMActivationCodePoolDeviceStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEmbeddedSIMActivationCodePoolDeviceStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEmbeddedSIMActivationCodePoolDeviceStates - Get deviceStates from deviceManagement. Navigational property to a
// list of device states for this pool.
func (c EmbeddedSIMActivationCodePoolDeviceStateClient) ListEmbeddedSIMActivationCodePoolDeviceStates(ctx context.Context, id beta.DeviceManagementEmbeddedSIMActivationCodePoolId, options ListEmbeddedSIMActivationCodePoolDeviceStatesOperationOptions) (result ListEmbeddedSIMActivationCodePoolDeviceStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEmbeddedSIMActivationCodePoolDeviceStatesCustomPager{},
		Path:          fmt.Sprintf("%s/deviceStates", id.ID()),
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
		Values *[]beta.EmbeddedSIMDeviceState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEmbeddedSIMActivationCodePoolDeviceStatesComplete retrieves all the results into a single object
func (c EmbeddedSIMActivationCodePoolDeviceStateClient) ListEmbeddedSIMActivationCodePoolDeviceStatesComplete(ctx context.Context, id beta.DeviceManagementEmbeddedSIMActivationCodePoolId, options ListEmbeddedSIMActivationCodePoolDeviceStatesOperationOptions) (ListEmbeddedSIMActivationCodePoolDeviceStatesCompleteResult, error) {
	return c.ListEmbeddedSIMActivationCodePoolDeviceStatesCompleteMatchingPredicate(ctx, id, options, EmbeddedSIMDeviceStateOperationPredicate{})
}

// ListEmbeddedSIMActivationCodePoolDeviceStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EmbeddedSIMActivationCodePoolDeviceStateClient) ListEmbeddedSIMActivationCodePoolDeviceStatesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementEmbeddedSIMActivationCodePoolId, options ListEmbeddedSIMActivationCodePoolDeviceStatesOperationOptions, predicate EmbeddedSIMDeviceStateOperationPredicate) (result ListEmbeddedSIMActivationCodePoolDeviceStatesCompleteResult, err error) {
	items := make([]beta.EmbeddedSIMDeviceState, 0)

	resp, err := c.ListEmbeddedSIMActivationCodePoolDeviceStates(ctx, id, options)
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

	result = ListEmbeddedSIMActivationCodePoolDeviceStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
