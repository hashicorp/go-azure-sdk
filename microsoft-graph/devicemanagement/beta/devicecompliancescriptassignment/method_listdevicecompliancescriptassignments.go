package devicecompliancescriptassignment

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

type ListDeviceComplianceScriptAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceHealthScriptAssignment
}

type ListDeviceComplianceScriptAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceHealthScriptAssignment
}

type ListDeviceComplianceScriptAssignmentsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListDeviceComplianceScriptAssignmentsOperationOptions() ListDeviceComplianceScriptAssignmentsOperationOptions {
	return ListDeviceComplianceScriptAssignmentsOperationOptions{}
}

func (o ListDeviceComplianceScriptAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceComplianceScriptAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceComplianceScriptAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceComplianceScriptAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceComplianceScriptAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceComplianceScriptAssignments - Get assignments from deviceManagement. The list of group assignments for the
// device compliance script
func (c DeviceComplianceScriptAssignmentClient) ListDeviceComplianceScriptAssignments(ctx context.Context, id beta.DeviceManagementDeviceComplianceScriptId, options ListDeviceComplianceScriptAssignmentsOperationOptions) (result ListDeviceComplianceScriptAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceComplianceScriptAssignmentsCustomPager{},
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
		Values *[]beta.DeviceHealthScriptAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceComplianceScriptAssignmentsComplete retrieves all the results into a single object
func (c DeviceComplianceScriptAssignmentClient) ListDeviceComplianceScriptAssignmentsComplete(ctx context.Context, id beta.DeviceManagementDeviceComplianceScriptId, options ListDeviceComplianceScriptAssignmentsOperationOptions) (ListDeviceComplianceScriptAssignmentsCompleteResult, error) {
	return c.ListDeviceComplianceScriptAssignmentsCompleteMatchingPredicate(ctx, id, options, DeviceHealthScriptAssignmentOperationPredicate{})
}

// ListDeviceComplianceScriptAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceComplianceScriptAssignmentClient) ListDeviceComplianceScriptAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDeviceComplianceScriptId, options ListDeviceComplianceScriptAssignmentsOperationOptions, predicate DeviceHealthScriptAssignmentOperationPredicate) (result ListDeviceComplianceScriptAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceHealthScriptAssignment, 0)

	resp, err := c.ListDeviceComplianceScriptAssignments(ctx, id, options)
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

	result = ListDeviceComplianceScriptAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
