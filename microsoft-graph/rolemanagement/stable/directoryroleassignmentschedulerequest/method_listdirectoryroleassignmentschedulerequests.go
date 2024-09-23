package directoryroleassignmentschedulerequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDirectoryRoleAssignmentScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleAssignmentScheduleRequest
}

type ListDirectoryRoleAssignmentScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleAssignmentScheduleRequest
}

type ListDirectoryRoleAssignmentScheduleRequestsOperationOptions struct {
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

func DefaultListDirectoryRoleAssignmentScheduleRequestsOperationOptions() ListDirectoryRoleAssignmentScheduleRequestsOperationOptions {
	return ListDirectoryRoleAssignmentScheduleRequestsOperationOptions{}
}

func (o ListDirectoryRoleAssignmentScheduleRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectoryRoleAssignmentScheduleRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListDirectoryRoleAssignmentScheduleRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectoryRoleAssignmentScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleAssignmentScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleAssignmentScheduleRequests - List roleAssignmentScheduleRequests. Retrieve the requests for active
// role assignments to principals. The active assignments include those made through assignments and activation
// requests, and directly through the role assignments API. The role assignments can be permanently active with or
// without an expiry date, or temporarily active after user activation of eligible assignments.
func (c DirectoryRoleAssignmentScheduleRequestClient) ListDirectoryRoleAssignmentScheduleRequests(ctx context.Context, options ListDirectoryRoleAssignmentScheduleRequestsOperationOptions) (result ListDirectoryRoleAssignmentScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectoryRoleAssignmentScheduleRequestsCustomPager{},
		Path:          "/roleManagement/directory/roleAssignmentScheduleRequests",
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
		Values *[]stable.UnifiedRoleAssignmentScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleAssignmentScheduleRequestsComplete retrieves all the results into a single object
func (c DirectoryRoleAssignmentScheduleRequestClient) ListDirectoryRoleAssignmentScheduleRequestsComplete(ctx context.Context, options ListDirectoryRoleAssignmentScheduleRequestsOperationOptions) (ListDirectoryRoleAssignmentScheduleRequestsCompleteResult, error) {
	return c.ListDirectoryRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx, options, UnifiedRoleAssignmentScheduleRequestOperationPredicate{})
}

// ListDirectoryRoleAssignmentScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleAssignmentScheduleRequestClient) ListDirectoryRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx context.Context, options ListDirectoryRoleAssignmentScheduleRequestsOperationOptions, predicate UnifiedRoleAssignmentScheduleRequestOperationPredicate) (result ListDirectoryRoleAssignmentScheduleRequestsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleAssignmentScheduleRequest, 0)

	resp, err := c.ListDirectoryRoleAssignmentScheduleRequests(ctx, options)
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

	result = ListDirectoryRoleAssignmentScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
