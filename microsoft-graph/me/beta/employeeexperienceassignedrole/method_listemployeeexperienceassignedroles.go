package employeeexperienceassignedrole

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

type ListEmployeeExperienceAssignedRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.EngagementRole
}

type ListEmployeeExperienceAssignedRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.EngagementRole
}

type ListEmployeeExperienceAssignedRolesOperationOptions struct {
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

func DefaultListEmployeeExperienceAssignedRolesOperationOptions() ListEmployeeExperienceAssignedRolesOperationOptions {
	return ListEmployeeExperienceAssignedRolesOperationOptions{}
}

func (o ListEmployeeExperienceAssignedRolesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEmployeeExperienceAssignedRolesOperationOptions) ToOData() *odata.Query {
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

func (o ListEmployeeExperienceAssignedRolesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEmployeeExperienceAssignedRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEmployeeExperienceAssignedRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEmployeeExperienceAssignedRoles - List assignedRoles. Get a list of all the roles assigned to a user in Viva
// Engage.
func (c EmployeeExperienceAssignedRoleClient) ListEmployeeExperienceAssignedRoles(ctx context.Context, options ListEmployeeExperienceAssignedRolesOperationOptions) (result ListEmployeeExperienceAssignedRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEmployeeExperienceAssignedRolesCustomPager{},
		Path:          "/me/employeeExperience/assignedRoles",
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
		Values *[]beta.EngagementRole `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEmployeeExperienceAssignedRolesComplete retrieves all the results into a single object
func (c EmployeeExperienceAssignedRoleClient) ListEmployeeExperienceAssignedRolesComplete(ctx context.Context, options ListEmployeeExperienceAssignedRolesOperationOptions) (ListEmployeeExperienceAssignedRolesCompleteResult, error) {
	return c.ListEmployeeExperienceAssignedRolesCompleteMatchingPredicate(ctx, options, EngagementRoleOperationPredicate{})
}

// ListEmployeeExperienceAssignedRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EmployeeExperienceAssignedRoleClient) ListEmployeeExperienceAssignedRolesCompleteMatchingPredicate(ctx context.Context, options ListEmployeeExperienceAssignedRolesOperationOptions, predicate EngagementRoleOperationPredicate) (result ListEmployeeExperienceAssignedRolesCompleteResult, err error) {
	items := make([]beta.EngagementRole, 0)

	resp, err := c.ListEmployeeExperienceAssignedRoles(ctx, options)
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

	result = ListEmployeeExperienceAssignedRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
