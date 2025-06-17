package authenticationmethoddevicehardwareoathdevice

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

type ListAuthenticationMethodDeviceHardwareOathDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HardwareOathTokenAuthenticationMethodDevice
}

type ListAuthenticationMethodDeviceHardwareOathDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HardwareOathTokenAuthenticationMethodDevice
}

type ListAuthenticationMethodDeviceHardwareOathDevicesOperationOptions struct {
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

func DefaultListAuthenticationMethodDeviceHardwareOathDevicesOperationOptions() ListAuthenticationMethodDeviceHardwareOathDevicesOperationOptions {
	return ListAuthenticationMethodDeviceHardwareOathDevicesOperationOptions{}
}

func (o ListAuthenticationMethodDeviceHardwareOathDevicesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationMethodDeviceHardwareOathDevicesOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationMethodDeviceHardwareOathDevicesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAuthenticationMethodDeviceHardwareOathDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationMethodDeviceHardwareOathDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationMethodDeviceHardwareOathDevices - List hardwareOathDevices. List all hardware OATH tokens in the
// directory.
func (c AuthenticationMethodDeviceHardwareOathDeviceClient) ListAuthenticationMethodDeviceHardwareOathDevices(ctx context.Context, options ListAuthenticationMethodDeviceHardwareOathDevicesOperationOptions) (result ListAuthenticationMethodDeviceHardwareOathDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationMethodDeviceHardwareOathDevicesCustomPager{},
		Path:          "/directory/authenticationMethodDevices/hardwareOathDevices",
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
		Values *[]beta.HardwareOathTokenAuthenticationMethodDevice `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationMethodDeviceHardwareOathDevicesComplete retrieves all the results into a single object
func (c AuthenticationMethodDeviceHardwareOathDeviceClient) ListAuthenticationMethodDeviceHardwareOathDevicesComplete(ctx context.Context, options ListAuthenticationMethodDeviceHardwareOathDevicesOperationOptions) (ListAuthenticationMethodDeviceHardwareOathDevicesCompleteResult, error) {
	return c.ListAuthenticationMethodDeviceHardwareOathDevicesCompleteMatchingPredicate(ctx, options, HardwareOathTokenAuthenticationMethodDeviceOperationPredicate{})
}

// ListAuthenticationMethodDeviceHardwareOathDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationMethodDeviceHardwareOathDeviceClient) ListAuthenticationMethodDeviceHardwareOathDevicesCompleteMatchingPredicate(ctx context.Context, options ListAuthenticationMethodDeviceHardwareOathDevicesOperationOptions, predicate HardwareOathTokenAuthenticationMethodDeviceOperationPredicate) (result ListAuthenticationMethodDeviceHardwareOathDevicesCompleteResult, err error) {
	items := make([]beta.HardwareOathTokenAuthenticationMethodDevice, 0)

	resp, err := c.ListAuthenticationMethodDeviceHardwareOathDevices(ctx, options)
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

	result = ListAuthenticationMethodDeviceHardwareOathDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
