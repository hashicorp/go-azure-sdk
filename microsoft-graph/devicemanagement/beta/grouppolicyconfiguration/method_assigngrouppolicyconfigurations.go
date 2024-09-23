package grouppolicyconfiguration

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

type AssignGroupPolicyConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyConfigurationAssignment
}

type AssignGroupPolicyConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyConfigurationAssignment
}

type AssignGroupPolicyConfigurationsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultAssignGroupPolicyConfigurationsOperationOptions() AssignGroupPolicyConfigurationsOperationOptions {
	return AssignGroupPolicyConfigurationsOperationOptions{}
}

func (o AssignGroupPolicyConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignGroupPolicyConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o AssignGroupPolicyConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AssignGroupPolicyConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignGroupPolicyConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignGroupPolicyConfigurations - Invoke action assign
func (c GroupPolicyConfigurationClient) AssignGroupPolicyConfigurations(ctx context.Context, id beta.DeviceManagementGroupPolicyConfigurationId, input AssignGroupPolicyConfigurationsRequest, options AssignGroupPolicyConfigurationsOperationOptions) (result AssignGroupPolicyConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AssignGroupPolicyConfigurationsCustomPager{},
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
		Values *[]beta.GroupPolicyConfigurationAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignGroupPolicyConfigurationsComplete retrieves all the results into a single object
func (c GroupPolicyConfigurationClient) AssignGroupPolicyConfigurationsComplete(ctx context.Context, id beta.DeviceManagementGroupPolicyConfigurationId, input AssignGroupPolicyConfigurationsRequest, options AssignGroupPolicyConfigurationsOperationOptions) (AssignGroupPolicyConfigurationsCompleteResult, error) {
	return c.AssignGroupPolicyConfigurationsCompleteMatchingPredicate(ctx, id, input, options, GroupPolicyConfigurationAssignmentOperationPredicate{})
}

// AssignGroupPolicyConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyConfigurationClient) AssignGroupPolicyConfigurationsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementGroupPolicyConfigurationId, input AssignGroupPolicyConfigurationsRequest, options AssignGroupPolicyConfigurationsOperationOptions, predicate GroupPolicyConfigurationAssignmentOperationPredicate) (result AssignGroupPolicyConfigurationsCompleteResult, err error) {
	items := make([]beta.GroupPolicyConfigurationAssignment, 0)

	resp, err := c.AssignGroupPolicyConfigurations(ctx, id, input, options)
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

	result = AssignGroupPolicyConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
