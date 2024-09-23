package deviceenrollmentconfiguration

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDeviceEnrollmentConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceEnrollmentConfiguration
}

type ListDeviceEnrollmentConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceEnrollmentConfiguration
}

type ListDeviceEnrollmentConfigurationsOperationOptions struct {
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

func DefaultListDeviceEnrollmentConfigurationsOperationOptions() ListDeviceEnrollmentConfigurationsOperationOptions {
	return ListDeviceEnrollmentConfigurationsOperationOptions{}
}

func (o ListDeviceEnrollmentConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceEnrollmentConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceEnrollmentConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceEnrollmentConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceEnrollmentConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceEnrollmentConfigurations - Get deviceEnrollmentConfigurations from deviceManagement. The list of device
// enrollment configurations
func (c DeviceEnrollmentConfigurationClient) ListDeviceEnrollmentConfigurations(ctx context.Context, options ListDeviceEnrollmentConfigurationsOperationOptions) (result ListDeviceEnrollmentConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceEnrollmentConfigurationsCustomPager{},
		Path:          "/deviceManagement/deviceEnrollmentConfigurations",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.DeviceEnrollmentConfiguration, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDeviceEnrollmentConfigurationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DeviceEnrollmentConfiguration (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDeviceEnrollmentConfigurationsComplete retrieves all the results into a single object
func (c DeviceEnrollmentConfigurationClient) ListDeviceEnrollmentConfigurationsComplete(ctx context.Context, options ListDeviceEnrollmentConfigurationsOperationOptions) (ListDeviceEnrollmentConfigurationsCompleteResult, error) {
	return c.ListDeviceEnrollmentConfigurationsCompleteMatchingPredicate(ctx, options, DeviceEnrollmentConfigurationOperationPredicate{})
}

// ListDeviceEnrollmentConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceEnrollmentConfigurationClient) ListDeviceEnrollmentConfigurationsCompleteMatchingPredicate(ctx context.Context, options ListDeviceEnrollmentConfigurationsOperationOptions, predicate DeviceEnrollmentConfigurationOperationPredicate) (result ListDeviceEnrollmentConfigurationsCompleteResult, err error) {
	items := make([]beta.DeviceEnrollmentConfiguration, 0)

	resp, err := c.ListDeviceEnrollmentConfigurations(ctx, options)
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

	result = ListDeviceEnrollmentConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
