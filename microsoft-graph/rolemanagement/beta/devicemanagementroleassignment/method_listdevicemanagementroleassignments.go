package devicemanagementroleassignment

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

type ListDeviceManagementRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignmentMultiple
}

type ListDeviceManagementRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignmentMultiple
}

type ListDeviceManagementRoleAssignmentsOperationOptions struct {
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

func DefaultListDeviceManagementRoleAssignmentsOperationOptions() ListDeviceManagementRoleAssignmentsOperationOptions {
	return ListDeviceManagementRoleAssignmentsOperationOptions{}
}

func (o ListDeviceManagementRoleAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceManagementRoleAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceManagementRoleAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceManagementRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementRoleAssignments - List unifiedRoleAssignmentMultiple. Get a list of unifiedRoleAssignmentMultiple
// objects for an RBAC provider. The following RBAC providers are currently supported: - Cloud PC - device management
// (Intune) For other Microsoft 365 applications (like Microsoft Entra ID), use unifiedRoleAssignment.
func (c DeviceManagementRoleAssignmentClient) ListDeviceManagementRoleAssignments(ctx context.Context, options ListDeviceManagementRoleAssignmentsOperationOptions) (result ListDeviceManagementRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceManagementRoleAssignmentsCustomPager{},
		Path:          "/roleManagement/deviceManagement/roleAssignments",
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
		Values *[]beta.UnifiedRoleAssignmentMultiple `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceManagementRoleAssignmentsComplete retrieves all the results into a single object
func (c DeviceManagementRoleAssignmentClient) ListDeviceManagementRoleAssignmentsComplete(ctx context.Context, options ListDeviceManagementRoleAssignmentsOperationOptions) (ListDeviceManagementRoleAssignmentsCompleteResult, error) {
	return c.ListDeviceManagementRoleAssignmentsCompleteMatchingPredicate(ctx, options, UnifiedRoleAssignmentMultipleOperationPredicate{})
}

// ListDeviceManagementRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementRoleAssignmentClient) ListDeviceManagementRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, options ListDeviceManagementRoleAssignmentsOperationOptions, predicate UnifiedRoleAssignmentMultipleOperationPredicate) (result ListDeviceManagementRoleAssignmentsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignmentMultiple, 0)

	resp, err := c.ListDeviceManagementRoleAssignments(ctx, options)
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

	result = ListDeviceManagementRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
