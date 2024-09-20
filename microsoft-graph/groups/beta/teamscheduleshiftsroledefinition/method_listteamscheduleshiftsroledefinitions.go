package teamscheduleshiftsroledefinition

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

type ListTeamScheduleShiftsRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ShiftsRoleDefinition
}

type ListTeamScheduleShiftsRoleDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ShiftsRoleDefinition
}

type ListTeamScheduleShiftsRoleDefinitionsOperationOptions struct {
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

func DefaultListTeamScheduleShiftsRoleDefinitionsOperationOptions() ListTeamScheduleShiftsRoleDefinitionsOperationOptions {
	return ListTeamScheduleShiftsRoleDefinitionsOperationOptions{}
}

func (o ListTeamScheduleShiftsRoleDefinitionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamScheduleShiftsRoleDefinitionsOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamScheduleShiftsRoleDefinitionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamScheduleShiftsRoleDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamScheduleShiftsRoleDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamScheduleShiftsRoleDefinitions - Get shiftsRoleDefinitions from groups. The definitions of the roles in the
// schedule.
func (c TeamScheduleShiftsRoleDefinitionClient) ListTeamScheduleShiftsRoleDefinitions(ctx context.Context, id beta.GroupId, options ListTeamScheduleShiftsRoleDefinitionsOperationOptions) (result ListTeamScheduleShiftsRoleDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamScheduleShiftsRoleDefinitionsCustomPager{},
		Path:          fmt.Sprintf("%s/team/schedule/shiftsRoleDefinitions", id.ID()),
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
		Values *[]beta.ShiftsRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamScheduleShiftsRoleDefinitionsComplete retrieves all the results into a single object
func (c TeamScheduleShiftsRoleDefinitionClient) ListTeamScheduleShiftsRoleDefinitionsComplete(ctx context.Context, id beta.GroupId, options ListTeamScheduleShiftsRoleDefinitionsOperationOptions) (ListTeamScheduleShiftsRoleDefinitionsCompleteResult, error) {
	return c.ListTeamScheduleShiftsRoleDefinitionsCompleteMatchingPredicate(ctx, id, options, ShiftsRoleDefinitionOperationPredicate{})
}

// ListTeamScheduleShiftsRoleDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamScheduleShiftsRoleDefinitionClient) ListTeamScheduleShiftsRoleDefinitionsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListTeamScheduleShiftsRoleDefinitionsOperationOptions, predicate ShiftsRoleDefinitionOperationPredicate) (result ListTeamScheduleShiftsRoleDefinitionsCompleteResult, err error) {
	items := make([]beta.ShiftsRoleDefinition, 0)

	resp, err := c.ListTeamScheduleShiftsRoleDefinitions(ctx, id, options)
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

	result = ListTeamScheduleShiftsRoleDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
