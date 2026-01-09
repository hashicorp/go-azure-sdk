package devicecompliancepolicy

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

type AssignDeviceCompliancePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceCompliancePolicyAssignment
}

type AssignDeviceCompliancePoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceCompliancePolicyAssignment
}

type AssignDeviceCompliancePoliciesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultAssignDeviceCompliancePoliciesOperationOptions() AssignDeviceCompliancePoliciesOperationOptions {
	return AssignDeviceCompliancePoliciesOperationOptions{}
}

func (o AssignDeviceCompliancePoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignDeviceCompliancePoliciesOperationOptions) ToOData() *odata.Query {
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

func (o AssignDeviceCompliancePoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type AssignDeviceCompliancePoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *AssignDeviceCompliancePoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AssignDeviceCompliancePolicies - Invoke action assign
func (c DeviceCompliancePolicyClient) AssignDeviceCompliancePolicies(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyId, input AssignDeviceCompliancePoliciesRequest, options AssignDeviceCompliancePoliciesOperationOptions) (result AssignDeviceCompliancePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &AssignDeviceCompliancePoliciesCustomPager{},
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
		Values *[]beta.DeviceCompliancePolicyAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AssignDeviceCompliancePoliciesComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyClient) AssignDeviceCompliancePoliciesComplete(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyId, input AssignDeviceCompliancePoliciesRequest, options AssignDeviceCompliancePoliciesOperationOptions) (AssignDeviceCompliancePoliciesCompleteResult, error) {
	return c.AssignDeviceCompliancePoliciesCompleteMatchingPredicate(ctx, id, input, options, DeviceCompliancePolicyAssignmentOperationPredicate{})
}

// AssignDeviceCompliancePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyClient) AssignDeviceCompliancePoliciesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyId, input AssignDeviceCompliancePoliciesRequest, options AssignDeviceCompliancePoliciesOperationOptions, predicate DeviceCompliancePolicyAssignmentOperationPredicate) (result AssignDeviceCompliancePoliciesCompleteResult, err error) {
	items := make([]beta.DeviceCompliancePolicyAssignment, 0)

	resp, err := c.AssignDeviceCompliancePolicies(ctx, id, input, options)
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

	result = AssignDeviceCompliancePoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
