package virtualendpointprovisioningpolicyassignment

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

type ListVirtualEndpointProvisioningPolicyAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCProvisioningPolicyAssignment
}

type ListVirtualEndpointProvisioningPolicyAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCProvisioningPolicyAssignment
}

type ListVirtualEndpointProvisioningPolicyAssignmentsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListVirtualEndpointProvisioningPolicyAssignmentsOperationOptions() ListVirtualEndpointProvisioningPolicyAssignmentsOperationOptions {
	return ListVirtualEndpointProvisioningPolicyAssignmentsOperationOptions{}
}

func (o ListVirtualEndpointProvisioningPolicyAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVirtualEndpointProvisioningPolicyAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListVirtualEndpointProvisioningPolicyAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListVirtualEndpointProvisioningPolicyAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointProvisioningPolicyAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointProvisioningPolicyAssignments - Get assignments from deviceManagement. A defined collection of
// provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Microsoft Entra ID
// that have provisioning policy assigned. Returned only on $expand. For an example about how to get the assignments
// relationship, see Get cloudPcProvisioningPolicy.
func (c VirtualEndpointProvisioningPolicyAssignmentClient) ListVirtualEndpointProvisioningPolicyAssignments(ctx context.Context, id beta.DeviceManagementVirtualEndpointProvisioningPolicyId, options ListVirtualEndpointProvisioningPolicyAssignmentsOperationOptions) (result ListVirtualEndpointProvisioningPolicyAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListVirtualEndpointProvisioningPolicyAssignmentsCustomPager{},
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
		Values *[]beta.CloudPCProvisioningPolicyAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointProvisioningPolicyAssignmentsComplete retrieves all the results into a single object
func (c VirtualEndpointProvisioningPolicyAssignmentClient) ListVirtualEndpointProvisioningPolicyAssignmentsComplete(ctx context.Context, id beta.DeviceManagementVirtualEndpointProvisioningPolicyId, options ListVirtualEndpointProvisioningPolicyAssignmentsOperationOptions) (ListVirtualEndpointProvisioningPolicyAssignmentsCompleteResult, error) {
	return c.ListVirtualEndpointProvisioningPolicyAssignmentsCompleteMatchingPredicate(ctx, id, options, CloudPCProvisioningPolicyAssignmentOperationPredicate{})
}

// ListVirtualEndpointProvisioningPolicyAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointProvisioningPolicyAssignmentClient) ListVirtualEndpointProvisioningPolicyAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementVirtualEndpointProvisioningPolicyId, options ListVirtualEndpointProvisioningPolicyAssignmentsOperationOptions, predicate CloudPCProvisioningPolicyAssignmentOperationPredicate) (result ListVirtualEndpointProvisioningPolicyAssignmentsCompleteResult, err error) {
	items := make([]beta.CloudPCProvisioningPolicyAssignment, 0)

	resp, err := c.ListVirtualEndpointProvisioningPolicyAssignments(ctx, id, options)
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

	result = ListVirtualEndpointProvisioningPolicyAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
