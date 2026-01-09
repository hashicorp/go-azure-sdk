package directorytransitiveroleassignment

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

type ListDirectoryTransitiveRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignment
}

type ListDirectoryTransitiveRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignment
}

type ListDirectoryTransitiveRoleAssignmentsOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Expand           *odata.Expand
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	RetryFunc        client.RequestRetryFunc
	Search           *string
	Select           *[]string
	Skip             *int64
	Top              *int64
}

func DefaultListDirectoryTransitiveRoleAssignmentsOperationOptions() ListDirectoryTransitiveRoleAssignmentsOperationOptions {
	return ListDirectoryTransitiveRoleAssignmentsOperationOptions{}
}

func (o ListDirectoryTransitiveRoleAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectoryTransitiveRoleAssignmentsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
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

func (o ListDirectoryTransitiveRoleAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectoryTransitiveRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryTransitiveRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryTransitiveRoleAssignments - List transitiveRoleAssignment. Get the list of direct and transitive
// unifiedRoleAssignment objects for a specific principal. For example, if a user is assigned a Microsoft Entra role
// through group membership, the role assignment is transitive, and this request will list the group's ID as the
// principalId. Results can also be filtered by the roleDefinitionId and directoryScopeId. Supported only for directory
// (Microsoft Entra ID) provider. For more information, see Use Microsoft Entra groups to manage role assignments.
func (c DirectoryTransitiveRoleAssignmentClient) ListDirectoryTransitiveRoleAssignments(ctx context.Context, options ListDirectoryTransitiveRoleAssignmentsOperationOptions) (result ListDirectoryTransitiveRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectoryTransitiveRoleAssignmentsCustomPager{},
		Path:          "/roleManagement/directory/transitiveRoleAssignments",
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
		Values *[]beta.UnifiedRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryTransitiveRoleAssignmentsComplete retrieves all the results into a single object
func (c DirectoryTransitiveRoleAssignmentClient) ListDirectoryTransitiveRoleAssignmentsComplete(ctx context.Context, options ListDirectoryTransitiveRoleAssignmentsOperationOptions) (ListDirectoryTransitiveRoleAssignmentsCompleteResult, error) {
	return c.ListDirectoryTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx, options, UnifiedRoleAssignmentOperationPredicate{})
}

// ListDirectoryTransitiveRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryTransitiveRoleAssignmentClient) ListDirectoryTransitiveRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, options ListDirectoryTransitiveRoleAssignmentsOperationOptions, predicate UnifiedRoleAssignmentOperationPredicate) (result ListDirectoryTransitiveRoleAssignmentsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignment, 0)

	resp, err := c.ListDirectoryTransitiveRoleAssignments(ctx, options)
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

	result = ListDirectoryTransitiveRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
