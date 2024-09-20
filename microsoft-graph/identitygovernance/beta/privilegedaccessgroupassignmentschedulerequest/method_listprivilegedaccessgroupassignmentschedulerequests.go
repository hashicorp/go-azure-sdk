package privilegedaccessgroupassignmentschedulerequest

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

type ListPrivilegedAccessGroupAssignmentScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PrivilegedAccessGroupAssignmentScheduleRequest
}

type ListPrivilegedAccessGroupAssignmentScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PrivilegedAccessGroupAssignmentScheduleRequest
}

type ListPrivilegedAccessGroupAssignmentScheduleRequestsOperationOptions struct {
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

func DefaultListPrivilegedAccessGroupAssignmentScheduleRequestsOperationOptions() ListPrivilegedAccessGroupAssignmentScheduleRequestsOperationOptions {
	return ListPrivilegedAccessGroupAssignmentScheduleRequestsOperationOptions{}
}

func (o ListPrivilegedAccessGroupAssignmentScheduleRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPrivilegedAccessGroupAssignmentScheduleRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListPrivilegedAccessGroupAssignmentScheduleRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPrivilegedAccessGroupAssignmentScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccessGroupAssignmentScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccessGroupAssignmentScheduleRequests - List assignmentScheduleRequests. Get a list of the
// privilegedAccessGroupAssignmentScheduleRequest objects and their properties.
func (c PrivilegedAccessGroupAssignmentScheduleRequestClient) ListPrivilegedAccessGroupAssignmentScheduleRequests(ctx context.Context, options ListPrivilegedAccessGroupAssignmentScheduleRequestsOperationOptions) (result ListPrivilegedAccessGroupAssignmentScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPrivilegedAccessGroupAssignmentScheduleRequestsCustomPager{},
		Path:          "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests",
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
		Values *[]beta.PrivilegedAccessGroupAssignmentScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccessGroupAssignmentScheduleRequestsComplete retrieves all the results into a single object
func (c PrivilegedAccessGroupAssignmentScheduleRequestClient) ListPrivilegedAccessGroupAssignmentScheduleRequestsComplete(ctx context.Context, options ListPrivilegedAccessGroupAssignmentScheduleRequestsOperationOptions) (ListPrivilegedAccessGroupAssignmentScheduleRequestsCompleteResult, error) {
	return c.ListPrivilegedAccessGroupAssignmentScheduleRequestsCompleteMatchingPredicate(ctx, options, PrivilegedAccessGroupAssignmentScheduleRequestOperationPredicate{})
}

// ListPrivilegedAccessGroupAssignmentScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccessGroupAssignmentScheduleRequestClient) ListPrivilegedAccessGroupAssignmentScheduleRequestsCompleteMatchingPredicate(ctx context.Context, options ListPrivilegedAccessGroupAssignmentScheduleRequestsOperationOptions, predicate PrivilegedAccessGroupAssignmentScheduleRequestOperationPredicate) (result ListPrivilegedAccessGroupAssignmentScheduleRequestsCompleteResult, err error) {
	items := make([]beta.PrivilegedAccessGroupAssignmentScheduleRequest, 0)

	resp, err := c.ListPrivilegedAccessGroupAssignmentScheduleRequests(ctx, options)
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

	result = ListPrivilegedAccessGroupAssignmentScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
