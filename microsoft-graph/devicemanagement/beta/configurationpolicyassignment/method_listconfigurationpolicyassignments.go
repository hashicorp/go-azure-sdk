package configurationpolicyassignment

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

type ListConfigurationPolicyAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicyAssignment
}

type ListConfigurationPolicyAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicyAssignment
}

type ListConfigurationPolicyAssignmentsOperationOptions struct {
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

func DefaultListConfigurationPolicyAssignmentsOperationOptions() ListConfigurationPolicyAssignmentsOperationOptions {
	return ListConfigurationPolicyAssignmentsOperationOptions{}
}

func (o ListConfigurationPolicyAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListConfigurationPolicyAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListConfigurationPolicyAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListConfigurationPolicyAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConfigurationPolicyAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigurationPolicyAssignments - Get assignments from deviceManagement. Policy assignments
func (c ConfigurationPolicyAssignmentClient) ListConfigurationPolicyAssignments(ctx context.Context, id beta.DeviceManagementConfigurationPolicyId, options ListConfigurationPolicyAssignmentsOperationOptions) (result ListConfigurationPolicyAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListConfigurationPolicyAssignmentsCustomPager{},
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

// ListConfigurationPolicyAssignmentsComplete retrieves all the results into a single object
func (c ConfigurationPolicyAssignmentClient) ListConfigurationPolicyAssignmentsComplete(ctx context.Context, id beta.DeviceManagementConfigurationPolicyId, options ListConfigurationPolicyAssignmentsOperationOptions) (ListConfigurationPolicyAssignmentsCompleteResult, error) {
	return c.ListConfigurationPolicyAssignmentsCompleteMatchingPredicate(ctx, id, options, DeviceManagementConfigurationPolicyAssignmentOperationPredicate{})
}

// ListConfigurationPolicyAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigurationPolicyAssignmentClient) ListConfigurationPolicyAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementConfigurationPolicyId, options ListConfigurationPolicyAssignmentsOperationOptions, predicate DeviceManagementConfigurationPolicyAssignmentOperationPredicate) (result ListConfigurationPolicyAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicyAssignment, 0)

	resp, err := c.ListConfigurationPolicyAssignments(ctx, id, options)
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

	result = ListConfigurationPolicyAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
