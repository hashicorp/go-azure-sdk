package devicemanagementscriptassignment

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

type ListScriptAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptAssignment
}

type ListScriptAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptAssignment
}

type ListScriptAssignmentsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListScriptAssignmentsOperationOptions() ListScriptAssignmentsOperationOptions {
	return ListScriptAssignmentsOperationOptions{}
}

func (o ListScriptAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListScriptAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListScriptAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListScriptAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListScriptAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListScriptAssignments - Get assignments from deviceManagement. The list of group assignments for the device
// management script.
func (c DeviceManagementScriptAssignmentClient) ListScriptAssignments(ctx context.Context, id beta.DeviceManagementDeviceManagementScriptId, options ListScriptAssignmentsOperationOptions) (result ListScriptAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListScriptAssignmentsCustomPager{},
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
		Values *[]beta.DeviceManagementScriptAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListScriptAssignmentsComplete retrieves all the results into a single object
func (c DeviceManagementScriptAssignmentClient) ListScriptAssignmentsComplete(ctx context.Context, id beta.DeviceManagementDeviceManagementScriptId, options ListScriptAssignmentsOperationOptions) (ListScriptAssignmentsCompleteResult, error) {
	return c.ListScriptAssignmentsCompleteMatchingPredicate(ctx, id, options, DeviceManagementScriptAssignmentOperationPredicate{})
}

// ListScriptAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementScriptAssignmentClient) ListScriptAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDeviceManagementScriptId, options ListScriptAssignmentsOperationOptions, predicate DeviceManagementScriptAssignmentOperationPredicate) (result ListScriptAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptAssignment, 0)

	resp, err := c.ListScriptAssignments(ctx, id, options)
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

	result = ListScriptAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
