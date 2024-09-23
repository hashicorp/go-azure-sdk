package devicecustomattributeshellscriptgroupassignment

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

type ListDeviceCustomAttributeShellScriptGroupAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptGroupAssignment
}

type ListDeviceCustomAttributeShellScriptGroupAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptGroupAssignment
}

type ListDeviceCustomAttributeShellScriptGroupAssignmentsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListDeviceCustomAttributeShellScriptGroupAssignmentsOperationOptions() ListDeviceCustomAttributeShellScriptGroupAssignmentsOperationOptions {
	return ListDeviceCustomAttributeShellScriptGroupAssignmentsOperationOptions{}
}

func (o ListDeviceCustomAttributeShellScriptGroupAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceCustomAttributeShellScriptGroupAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceCustomAttributeShellScriptGroupAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceCustomAttributeShellScriptGroupAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCustomAttributeShellScriptGroupAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCustomAttributeShellScriptGroupAssignments - Get groupAssignments from deviceManagement. The list of group
// assignments for the device management script.
func (c DeviceCustomAttributeShellScriptGroupAssignmentClient) ListDeviceCustomAttributeShellScriptGroupAssignments(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptId, options ListDeviceCustomAttributeShellScriptGroupAssignmentsOperationOptions) (result ListDeviceCustomAttributeShellScriptGroupAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceCustomAttributeShellScriptGroupAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/groupAssignments", id.ID()),
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
		Values *[]beta.DeviceManagementScriptGroupAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCustomAttributeShellScriptGroupAssignmentsComplete retrieves all the results into a single object
func (c DeviceCustomAttributeShellScriptGroupAssignmentClient) ListDeviceCustomAttributeShellScriptGroupAssignmentsComplete(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptId, options ListDeviceCustomAttributeShellScriptGroupAssignmentsOperationOptions) (ListDeviceCustomAttributeShellScriptGroupAssignmentsCompleteResult, error) {
	return c.ListDeviceCustomAttributeShellScriptGroupAssignmentsCompleteMatchingPredicate(ctx, id, options, DeviceManagementScriptGroupAssignmentOperationPredicate{})
}

// ListDeviceCustomAttributeShellScriptGroupAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCustomAttributeShellScriptGroupAssignmentClient) ListDeviceCustomAttributeShellScriptGroupAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptId, options ListDeviceCustomAttributeShellScriptGroupAssignmentsOperationOptions, predicate DeviceManagementScriptGroupAssignmentOperationPredicate) (result ListDeviceCustomAttributeShellScriptGroupAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptGroupAssignment, 0)

	resp, err := c.ListDeviceCustomAttributeShellScriptGroupAssignments(ctx, id, options)
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

	result = ListDeviceCustomAttributeShellScriptGroupAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
