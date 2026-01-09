package compliancepolicy

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

type AssignCompliancePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicyAssignment
}

type AssignCompliancePoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicyAssignment
}

type AssignCompliancePoliciesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultAssignCompliancePoliciesOperationOptions() AssignCompliancePoliciesOperationOptions {
	return AssignCompliancePoliciesOperationOptions{}
}

func (o AssignCompliancePoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignCompliancePoliciesOperationOptions) ToOData() *odata.Query {
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

func (o AssignCompliancePoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AssignCompliancePoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignCompliancePoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignCompliancePolicies - Invoke action assign
func (c CompliancePolicyClient) AssignCompliancePolicies(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, input AssignCompliancePoliciesRequest, options AssignCompliancePoliciesOperationOptions) (result AssignCompliancePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AssignCompliancePoliciesCustomPager{},
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

// AssignCompliancePoliciesComplete retrieves all the results into a single object
func (c CompliancePolicyClient) AssignCompliancePoliciesComplete(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, input AssignCompliancePoliciesRequest, options AssignCompliancePoliciesOperationOptions) (AssignCompliancePoliciesCompleteResult, error) {
	return c.AssignCompliancePoliciesCompleteMatchingPredicate(ctx, id, input, options, DeviceManagementConfigurationPolicyAssignmentOperationPredicate{})
}

// AssignCompliancePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicyClient) AssignCompliancePoliciesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, input AssignCompliancePoliciesRequest, options AssignCompliancePoliciesOperationOptions, predicate DeviceManagementConfigurationPolicyAssignmentOperationPredicate) (result AssignCompliancePoliciesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicyAssignment, 0)

	resp, err := c.AssignCompliancePolicies(ctx, id, input, options)
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

	result = AssignCompliancePoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
