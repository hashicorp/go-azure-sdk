package comanageddevicemanageddevicemobileappconfigurationstate

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

type ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedDeviceMobileAppConfigurationState
}

type ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedDeviceMobileAppConfigurationState
}

type ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationOptions struct {
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

func DefaultListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationOptions() ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationOptions {
	return ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationOptions{}
}

func (o ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationOptions) ToOData() *odata.Query {
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

func (o ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDeviceManagedDeviceMobileAppConfigurationStates - Get managedDeviceMobileAppConfigurationStates from
// deviceManagement. Managed device mobile app configuration states for this device.
func (c ComanagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListComanagedDeviceManagedDeviceMobileAppConfigurationStates(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationOptions) (result ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesCustomPager{},
		Path:          fmt.Sprintf("%s/managedDeviceMobileAppConfigurationStates", id.ID()),
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
		Values *[]beta.ManagedDeviceMobileAppConfigurationState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesComplete retrieves all the results into a single object
func (c ComanagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesComplete(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationOptions) (ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteResult, error) {
	return c.ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate(ctx, id, options, ManagedDeviceMobileAppConfigurationStateOperationPredicate{})
}

// ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceManagedDeviceMobileAppConfigurationStateClient) ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesOperationOptions, predicate ManagedDeviceMobileAppConfigurationStateOperationPredicate) (result ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteResult, err error) {
	items := make([]beta.ManagedDeviceMobileAppConfigurationState, 0)

	resp, err := c.ListComanagedDeviceManagedDeviceMobileAppConfigurationStates(ctx, id, options)
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

	result = ListComanagedDeviceManagedDeviceMobileAppConfigurationStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
