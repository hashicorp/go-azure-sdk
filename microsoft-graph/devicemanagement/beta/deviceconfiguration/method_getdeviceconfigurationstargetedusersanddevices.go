package deviceconfiguration

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

type GetDeviceConfigurationsTargetedUsersAndDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceConfigurationTargetedUserAndDevice
}

type GetDeviceConfigurationsTargetedUsersAndDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceConfigurationTargetedUserAndDevice
}

type GetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultGetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions() GetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions {
	return GetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions{}
}

func (o GetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o GetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetDeviceConfigurationsTargetedUsersAndDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetDeviceConfigurationsTargetedUsersAndDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetDeviceConfigurationsTargetedUsersAndDevices - Invoke action getTargetedUsersAndDevices
func (c DeviceConfigurationClient) GetDeviceConfigurationsTargetedUsersAndDevices(ctx context.Context, input GetDeviceConfigurationsTargetedUsersAndDevicesRequest, options GetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions) (result GetDeviceConfigurationsTargetedUsersAndDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetDeviceConfigurationsTargetedUsersAndDevicesCustomPager{},
		Path:          "/deviceManagement/deviceConfigurations/getTargetedUsersAndDevices",
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
		Values *[]beta.DeviceConfigurationTargetedUserAndDevice `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetDeviceConfigurationsTargetedUsersAndDevicesComplete retrieves all the results into a single object
func (c DeviceConfigurationClient) GetDeviceConfigurationsTargetedUsersAndDevicesComplete(ctx context.Context, input GetDeviceConfigurationsTargetedUsersAndDevicesRequest, options GetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions) (GetDeviceConfigurationsTargetedUsersAndDevicesCompleteResult, error) {
	return c.GetDeviceConfigurationsTargetedUsersAndDevicesCompleteMatchingPredicate(ctx, input, options, DeviceConfigurationTargetedUserAndDeviceOperationPredicate{})
}

// GetDeviceConfigurationsTargetedUsersAndDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationClient) GetDeviceConfigurationsTargetedUsersAndDevicesCompleteMatchingPredicate(ctx context.Context, input GetDeviceConfigurationsTargetedUsersAndDevicesRequest, options GetDeviceConfigurationsTargetedUsersAndDevicesOperationOptions, predicate DeviceConfigurationTargetedUserAndDeviceOperationPredicate) (result GetDeviceConfigurationsTargetedUsersAndDevicesCompleteResult, err error) {
	items := make([]beta.DeviceConfigurationTargetedUserAndDevice, 0)

	resp, err := c.GetDeviceConfigurationsTargetedUsersAndDevices(ctx, input, options)
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

	result = GetDeviceConfigurationsTargetedUsersAndDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
