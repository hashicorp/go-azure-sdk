package hardwareconfiguration

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

type ListHardwareConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HardwareConfiguration
}

type ListHardwareConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HardwareConfiguration
}

type ListHardwareConfigurationsOperationOptions struct {
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

func DefaultListHardwareConfigurationsOperationOptions() ListHardwareConfigurationsOperationOptions {
	return ListHardwareConfigurationsOperationOptions{}
}

func (o ListHardwareConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListHardwareConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o ListHardwareConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListHardwareConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListHardwareConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListHardwareConfigurations - Get hardwareConfigurations from deviceManagement. BIOS configuration and other settings
// provides customers the ability to configure hardware/bios settings on the enrolled Windows 10/11 Entra ID joined
// devices by uploading a configuration file generated with their OEM tool (e.g. Dell Command tool). A BIOS
// configuration policy can be assigned to multiple devices, allowing admins to remotely control a device's hardware
// properties (e.g. enable Secure Boot) from the Intune Portal. Supported for Dell only at this time.
func (c HardwareConfigurationClient) ListHardwareConfigurations(ctx context.Context, options ListHardwareConfigurationsOperationOptions) (result ListHardwareConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListHardwareConfigurationsCustomPager{},
		Path:          "/deviceManagement/hardwareConfigurations",
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
		Values *[]beta.HardwareConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListHardwareConfigurationsComplete retrieves all the results into a single object
func (c HardwareConfigurationClient) ListHardwareConfigurationsComplete(ctx context.Context, options ListHardwareConfigurationsOperationOptions) (ListHardwareConfigurationsCompleteResult, error) {
	return c.ListHardwareConfigurationsCompleteMatchingPredicate(ctx, options, HardwareConfigurationOperationPredicate{})
}

// ListHardwareConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HardwareConfigurationClient) ListHardwareConfigurationsCompleteMatchingPredicate(ctx context.Context, options ListHardwareConfigurationsOperationOptions, predicate HardwareConfigurationOperationPredicate) (result ListHardwareConfigurationsCompleteResult, err error) {
	items := make([]beta.HardwareConfiguration, 0)

	resp, err := c.ListHardwareConfigurations(ctx, options)
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

	result = ListHardwareConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
