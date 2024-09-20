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

type ListDeviceManagementScriptAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptAssignment
}

type ListDeviceManagementScriptAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptAssignment
}

type ListDeviceManagementScriptAssignmentsOperationOptions struct {
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

func DefaultListDeviceManagementScriptAssignmentsOperationOptions() ListDeviceManagementScriptAssignmentsOperationOptions {
	return ListDeviceManagementScriptAssignmentsOperationOptions{}
}

func (o ListDeviceManagementScriptAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceManagementScriptAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceManagementScriptAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceManagementScriptAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementScriptAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementScriptAssignments - Get assignments from deviceManagement. The list of group assignments for the
// device management script.
func (c DeviceManagementScriptAssignmentClient) ListDeviceManagementScriptAssignments(ctx context.Context, id beta.DeviceManagementDeviceManagementScriptId, options ListDeviceManagementScriptAssignmentsOperationOptions) (result ListDeviceManagementScriptAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceManagementScriptAssignmentsCustomPager{},
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

// ListDeviceManagementScriptAssignmentsComplete retrieves all the results into a single object
func (c DeviceManagementScriptAssignmentClient) ListDeviceManagementScriptAssignmentsComplete(ctx context.Context, id beta.DeviceManagementDeviceManagementScriptId, options ListDeviceManagementScriptAssignmentsOperationOptions) (ListDeviceManagementScriptAssignmentsCompleteResult, error) {
	return c.ListDeviceManagementScriptAssignmentsCompleteMatchingPredicate(ctx, id, options, DeviceManagementScriptAssignmentOperationPredicate{})
}

// ListDeviceManagementScriptAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementScriptAssignmentClient) ListDeviceManagementScriptAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDeviceManagementScriptId, options ListDeviceManagementScriptAssignmentsOperationOptions, predicate DeviceManagementScriptAssignmentOperationPredicate) (result ListDeviceManagementScriptAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptAssignment, 0)

	resp, err := c.ListDeviceManagementScriptAssignments(ctx, id, options)
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

	result = ListDeviceManagementScriptAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
