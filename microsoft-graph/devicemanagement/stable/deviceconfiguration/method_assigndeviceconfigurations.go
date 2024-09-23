package deviceconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceConfigurationAssignment
}

type AssignDeviceConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceConfigurationAssignment
}

type AssignDeviceConfigurationsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultAssignDeviceConfigurationsOperationOptions() AssignDeviceConfigurationsOperationOptions {
	return AssignDeviceConfigurationsOperationOptions{}
}

func (o AssignDeviceConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignDeviceConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o AssignDeviceConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AssignDeviceConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignDeviceConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignDeviceConfigurations - Invoke action assign. Not yet documented
func (c DeviceConfigurationClient) AssignDeviceConfigurations(ctx context.Context, id stable.DeviceManagementDeviceConfigurationId, input AssignDeviceConfigurationsRequest, options AssignDeviceConfigurationsOperationOptions) (result AssignDeviceConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AssignDeviceConfigurationsCustomPager{},
		Path:          fmt.Sprintf("%s/assign", id.ID()),
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
		Values *[]stable.DeviceConfigurationAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignDeviceConfigurationsComplete retrieves all the results into a single object
func (c DeviceConfigurationClient) AssignDeviceConfigurationsComplete(ctx context.Context, id stable.DeviceManagementDeviceConfigurationId, input AssignDeviceConfigurationsRequest, options AssignDeviceConfigurationsOperationOptions) (AssignDeviceConfigurationsCompleteResult, error) {
	return c.AssignDeviceConfigurationsCompleteMatchingPredicate(ctx, id, input, options, DeviceConfigurationAssignmentOperationPredicate{})
}

// AssignDeviceConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationClient) AssignDeviceConfigurationsCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementDeviceConfigurationId, input AssignDeviceConfigurationsRequest, options AssignDeviceConfigurationsOperationOptions, predicate DeviceConfigurationAssignmentOperationPredicate) (result AssignDeviceConfigurationsCompleteResult, err error) {
	items := make([]stable.DeviceConfigurationAssignment, 0)

	resp, err := c.AssignDeviceConfigurations(ctx, id, input, options)
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

	result = AssignDeviceConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
