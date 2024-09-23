package directoryroleassignmentschedule

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

type ListDirectoryRoleAssignmentSchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleAssignmentSchedule
}

type ListDirectoryRoleAssignmentSchedulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleAssignmentSchedule
}

type ListDirectoryRoleAssignmentSchedulesOperationOptions struct {
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

func DefaultListDirectoryRoleAssignmentSchedulesOperationOptions() ListDirectoryRoleAssignmentSchedulesOperationOptions {
	return ListDirectoryRoleAssignmentSchedulesOperationOptions{}
}

func (o ListDirectoryRoleAssignmentSchedulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectoryRoleAssignmentSchedulesOperationOptions) ToOData() *odata.Query {
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

func (o ListDirectoryRoleAssignmentSchedulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectoryRoleAssignmentSchedulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleAssignmentSchedulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleAssignmentSchedules - List roleAssignmentSchedules. Get the schedules for active role assignment
// operations.
func (c DirectoryRoleAssignmentScheduleClient) ListDirectoryRoleAssignmentSchedules(ctx context.Context, options ListDirectoryRoleAssignmentSchedulesOperationOptions) (result ListDirectoryRoleAssignmentSchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectoryRoleAssignmentSchedulesCustomPager{},
		Path:          "/roleManagement/directory/roleAssignmentSchedules",
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
		Values *[]stable.UnifiedRoleAssignmentSchedule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleAssignmentSchedulesComplete retrieves all the results into a single object
func (c DirectoryRoleAssignmentScheduleClient) ListDirectoryRoleAssignmentSchedulesComplete(ctx context.Context, options ListDirectoryRoleAssignmentSchedulesOperationOptions) (ListDirectoryRoleAssignmentSchedulesCompleteResult, error) {
	return c.ListDirectoryRoleAssignmentSchedulesCompleteMatchingPredicate(ctx, options, UnifiedRoleAssignmentScheduleOperationPredicate{})
}

// ListDirectoryRoleAssignmentSchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleAssignmentScheduleClient) ListDirectoryRoleAssignmentSchedulesCompleteMatchingPredicate(ctx context.Context, options ListDirectoryRoleAssignmentSchedulesOperationOptions, predicate UnifiedRoleAssignmentScheduleOperationPredicate) (result ListDirectoryRoleAssignmentSchedulesCompleteResult, err error) {
	items := make([]stable.UnifiedRoleAssignmentSchedule, 0)

	resp, err := c.ListDirectoryRoleAssignmentSchedules(ctx, options)
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

	result = ListDirectoryRoleAssignmentSchedulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
