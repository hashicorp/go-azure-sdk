package compliancepolicyassignment

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

type ListCompliancePolicyAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicyAssignment
}

type ListCompliancePolicyAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicyAssignment
}

type ListCompliancePolicyAssignmentsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListCompliancePolicyAssignmentsOperationOptions() ListCompliancePolicyAssignmentsOperationOptions {
	return ListCompliancePolicyAssignmentsOperationOptions{}
}

func (o ListCompliancePolicyAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCompliancePolicyAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListCompliancePolicyAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCompliancePolicyAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCompliancePolicyAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCompliancePolicyAssignments - Get assignments from deviceManagement. Policy assignments
func (c CompliancePolicyAssignmentClient) ListCompliancePolicyAssignments(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, options ListCompliancePolicyAssignmentsOperationOptions) (result ListCompliancePolicyAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCompliancePolicyAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/assignments", id.ID()),
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

// ListCompliancePolicyAssignmentsComplete retrieves all the results into a single object
func (c CompliancePolicyAssignmentClient) ListCompliancePolicyAssignmentsComplete(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, options ListCompliancePolicyAssignmentsOperationOptions) (ListCompliancePolicyAssignmentsCompleteResult, error) {
	return c.ListCompliancePolicyAssignmentsCompleteMatchingPredicate(ctx, id, options, DeviceManagementConfigurationPolicyAssignmentOperationPredicate{})
}

// ListCompliancePolicyAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicyAssignmentClient) ListCompliancePolicyAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, options ListCompliancePolicyAssignmentsOperationOptions, predicate DeviceManagementConfigurationPolicyAssignmentOperationPredicate) (result ListCompliancePolicyAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicyAssignment, 0)

	resp, err := c.ListCompliancePolicyAssignments(ctx, id, options)
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

	result = ListCompliancePolicyAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
