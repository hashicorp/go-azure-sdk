package directoryroleassignmentscheduleinstance

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

type ListDirectoryRoleAssignmentScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignmentScheduleInstance
}

type ListDirectoryRoleAssignmentScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignmentScheduleInstance
}

type ListDirectoryRoleAssignmentScheduleInstancesOperationOptions struct {
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

func DefaultListDirectoryRoleAssignmentScheduleInstancesOperationOptions() ListDirectoryRoleAssignmentScheduleInstancesOperationOptions {
	return ListDirectoryRoleAssignmentScheduleInstancesOperationOptions{}
}

func (o ListDirectoryRoleAssignmentScheduleInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectoryRoleAssignmentScheduleInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListDirectoryRoleAssignmentScheduleInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectoryRoleAssignmentScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleAssignmentScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleAssignmentScheduleInstances - List roleAssignmentScheduleInstances. Get the instances of active role
// assignments in your tenant. The active assignments include those made through assignments and activation requests,
// and directly through the role assignments API.
func (c DirectoryRoleAssignmentScheduleInstanceClient) ListDirectoryRoleAssignmentScheduleInstances(ctx context.Context, options ListDirectoryRoleAssignmentScheduleInstancesOperationOptions) (result ListDirectoryRoleAssignmentScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectoryRoleAssignmentScheduleInstancesCustomPager{},
		Path:          "/roleManagement/directory/roleAssignmentScheduleInstances",
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
		Values *[]beta.UnifiedRoleAssignmentScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleAssignmentScheduleInstancesComplete retrieves all the results into a single object
func (c DirectoryRoleAssignmentScheduleInstanceClient) ListDirectoryRoleAssignmentScheduleInstancesComplete(ctx context.Context, options ListDirectoryRoleAssignmentScheduleInstancesOperationOptions) (ListDirectoryRoleAssignmentScheduleInstancesCompleteResult, error) {
	return c.ListDirectoryRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx, options, UnifiedRoleAssignmentScheduleInstanceOperationPredicate{})
}

// ListDirectoryRoleAssignmentScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleAssignmentScheduleInstanceClient) ListDirectoryRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx context.Context, options ListDirectoryRoleAssignmentScheduleInstancesOperationOptions, predicate UnifiedRoleAssignmentScheduleInstanceOperationPredicate) (result ListDirectoryRoleAssignmentScheduleInstancesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignmentScheduleInstance, 0)

	resp, err := c.ListDirectoryRoleAssignmentScheduleInstances(ctx, options)
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

	result = ListDirectoryRoleAssignmentScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
