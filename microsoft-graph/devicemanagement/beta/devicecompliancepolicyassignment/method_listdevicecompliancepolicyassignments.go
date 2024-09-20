package devicecompliancepolicyassignment

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

type ListDeviceCompliancePolicyAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceCompliancePolicyAssignment
}

type ListDeviceCompliancePolicyAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceCompliancePolicyAssignment
}

type ListDeviceCompliancePolicyAssignmentsOperationOptions struct {
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

func DefaultListDeviceCompliancePolicyAssignmentsOperationOptions() ListDeviceCompliancePolicyAssignmentsOperationOptions {
	return ListDeviceCompliancePolicyAssignmentsOperationOptions{}
}

func (o ListDeviceCompliancePolicyAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceCompliancePolicyAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceCompliancePolicyAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceCompliancePolicyAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePolicyAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicyAssignments - Get assignments from deviceManagement. The collection of assignments for this
// compliance policy.
func (c DeviceCompliancePolicyAssignmentClient) ListDeviceCompliancePolicyAssignments(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyId, options ListDeviceCompliancePolicyAssignmentsOperationOptions) (result ListDeviceCompliancePolicyAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceCompliancePolicyAssignmentsCustomPager{},
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
		Values *[]beta.DeviceCompliancePolicyAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePolicyAssignmentsComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyAssignmentClient) ListDeviceCompliancePolicyAssignmentsComplete(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyId, options ListDeviceCompliancePolicyAssignmentsOperationOptions) (ListDeviceCompliancePolicyAssignmentsCompleteResult, error) {
	return c.ListDeviceCompliancePolicyAssignmentsCompleteMatchingPredicate(ctx, id, options, DeviceCompliancePolicyAssignmentOperationPredicate{})
}

// ListDeviceCompliancePolicyAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyAssignmentClient) ListDeviceCompliancePolicyAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyId, options ListDeviceCompliancePolicyAssignmentsOperationOptions, predicate DeviceCompliancePolicyAssignmentOperationPredicate) (result ListDeviceCompliancePolicyAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceCompliancePolicyAssignment, 0)

	resp, err := c.ListDeviceCompliancePolicyAssignments(ctx, id, options)
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

	result = ListDeviceCompliancePolicyAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
