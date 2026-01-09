package configurationpolicy

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

type AssignConfigurationPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicyAssignment
}

type AssignConfigurationPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicyAssignment
}

type AssignConfigurationPoliciesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultAssignConfigurationPoliciesOperationOptions() AssignConfigurationPoliciesOperationOptions {
	return AssignConfigurationPoliciesOperationOptions{}
}

func (o AssignConfigurationPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignConfigurationPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o AssignConfigurationPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AssignConfigurationPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignConfigurationPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignConfigurationPolicies - Invoke action assign
func (c ConfigurationPolicyClient) AssignConfigurationPolicies(ctx context.Context, id beta.DeviceManagementConfigurationPolicyId, input AssignConfigurationPoliciesRequest, options AssignConfigurationPoliciesOperationOptions) (result AssignConfigurationPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AssignConfigurationPoliciesCustomPager{},
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
		Values *[]beta.DeviceManagementConfigurationPolicyAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignConfigurationPoliciesComplete retrieves all the results into a single object
func (c ConfigurationPolicyClient) AssignConfigurationPoliciesComplete(ctx context.Context, id beta.DeviceManagementConfigurationPolicyId, input AssignConfigurationPoliciesRequest, options AssignConfigurationPoliciesOperationOptions) (AssignConfigurationPoliciesCompleteResult, error) {
	return c.AssignConfigurationPoliciesCompleteMatchingPredicate(ctx, id, input, options, DeviceManagementConfigurationPolicyAssignmentOperationPredicate{})
}

// AssignConfigurationPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigurationPolicyClient) AssignConfigurationPoliciesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementConfigurationPolicyId, input AssignConfigurationPoliciesRequest, options AssignConfigurationPoliciesOperationOptions, predicate DeviceManagementConfigurationPolicyAssignmentOperationPredicate) (result AssignConfigurationPoliciesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicyAssignment, 0)

	resp, err := c.AssignConfigurationPolicies(ctx, id, input, options)
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

	result = AssignConfigurationPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
