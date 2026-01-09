package hardwareconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignHardwareConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HardwareConfigurationAssignment
}

type AssignHardwareConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HardwareConfigurationAssignment
}

type AssignHardwareConfigurationsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultAssignHardwareConfigurationsOperationOptions() AssignHardwareConfigurationsOperationOptions {
	return AssignHardwareConfigurationsOperationOptions{}
}

func (o AssignHardwareConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignHardwareConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o AssignHardwareConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AssignHardwareConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignHardwareConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignHardwareConfigurations - Invoke action assign
func (c HardwareConfigurationClient) AssignHardwareConfigurations(ctx context.Context, id beta.DeviceManagementHardwareConfigurationId, input AssignHardwareConfigurationsRequest, options AssignHardwareConfigurationsOperationOptions) (result AssignHardwareConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AssignHardwareConfigurationsCustomPager{},
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
		Values *[]beta.HardwareConfigurationAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignHardwareConfigurationsComplete retrieves all the results into a single object
func (c HardwareConfigurationClient) AssignHardwareConfigurationsComplete(ctx context.Context, id beta.DeviceManagementHardwareConfigurationId, input AssignHardwareConfigurationsRequest, options AssignHardwareConfigurationsOperationOptions) (AssignHardwareConfigurationsCompleteResult, error) {
	return c.AssignHardwareConfigurationsCompleteMatchingPredicate(ctx, id, input, options, HardwareConfigurationAssignmentOperationPredicate{})
}

// AssignHardwareConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HardwareConfigurationClient) AssignHardwareConfigurationsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementHardwareConfigurationId, input AssignHardwareConfigurationsRequest, options AssignHardwareConfigurationsOperationOptions, predicate HardwareConfigurationAssignmentOperationPredicate) (result AssignHardwareConfigurationsCompleteResult, err error) {
	items := make([]beta.HardwareConfigurationAssignment, 0)

	resp, err := c.AssignHardwareConfigurations(ctx, id, input, options)
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

	result = AssignHardwareConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
