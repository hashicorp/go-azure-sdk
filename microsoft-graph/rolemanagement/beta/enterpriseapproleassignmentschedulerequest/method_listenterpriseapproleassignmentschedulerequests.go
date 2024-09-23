package enterpriseapproleassignmentschedulerequest

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

type ListEnterpriseAppRoleAssignmentScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignmentScheduleRequest
}

type ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignmentScheduleRequest
}

type ListEnterpriseAppRoleAssignmentScheduleRequestsOperationOptions struct {
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

func DefaultListEnterpriseAppRoleAssignmentScheduleRequestsOperationOptions() ListEnterpriseAppRoleAssignmentScheduleRequestsOperationOptions {
	return ListEnterpriseAppRoleAssignmentScheduleRequestsOperationOptions{}
}

func (o ListEnterpriseAppRoleAssignmentScheduleRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEnterpriseAppRoleAssignmentScheduleRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListEnterpriseAppRoleAssignmentScheduleRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEnterpriseAppRoleAssignmentScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleAssignmentScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleAssignmentScheduleRequests - Get roleAssignmentScheduleRequests from roleManagement
func (c EnterpriseAppRoleAssignmentScheduleRequestClient) ListEnterpriseAppRoleAssignmentScheduleRequests(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleAssignmentScheduleRequestsOperationOptions) (result ListEnterpriseAppRoleAssignmentScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEnterpriseAppRoleAssignmentScheduleRequestsCustomPager{},
		Path:          fmt.Sprintf("%s/roleAssignmentScheduleRequests", id.ID()),
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
		Values *[]beta.UnifiedRoleAssignmentScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnterpriseAppRoleAssignmentScheduleRequestsComplete retrieves all the results into a single object
func (c EnterpriseAppRoleAssignmentScheduleRequestClient) ListEnterpriseAppRoleAssignmentScheduleRequestsComplete(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleAssignmentScheduleRequestsOperationOptions) (ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteResult, error) {
	return c.ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx, id, options, UnifiedRoleAssignmentScheduleRequestOperationPredicate{})
}

// ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleAssignmentScheduleRequestClient) ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleAssignmentScheduleRequestsOperationOptions, predicate UnifiedRoleAssignmentScheduleRequestOperationPredicate) (result ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignmentScheduleRequest, 0)

	resp, err := c.ListEnterpriseAppRoleAssignmentScheduleRequests(ctx, id, options)
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

	result = ListEnterpriseAppRoleAssignmentScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
