package profileproject

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

type ListProfileProjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ProjectParticipation
}

type ListProfileProjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ProjectParticipation
}

type ListProfileProjectsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListProfileProjectsOperationOptions() ListProfileProjectsOperationOptions {
	return ListProfileProjectsOperationOptions{}
}

func (o ListProfileProjectsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProfileProjectsOperationOptions) ToOData() *odata.Query {
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

func (o ListProfileProjectsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProfileProjectsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileProjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileProjects - Get projects from users. Represents detailed information about projects associated with a user.
func (c ProfileProjectClient) ListProfileProjects(ctx context.Context, id beta.UserId, options ListProfileProjectsOperationOptions) (result ListProfileProjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProfileProjectsCustomPager{},
		Path:          fmt.Sprintf("%s/profile/projects", id.ID()),
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
		Values *[]beta.ProjectParticipation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileProjectsComplete retrieves all the results into a single object
func (c ProfileProjectClient) ListProfileProjectsComplete(ctx context.Context, id beta.UserId, options ListProfileProjectsOperationOptions) (ListProfileProjectsCompleteResult, error) {
	return c.ListProfileProjectsCompleteMatchingPredicate(ctx, id, options, ProjectParticipationOperationPredicate{})
}

// ListProfileProjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileProjectClient) ListProfileProjectsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListProfileProjectsOperationOptions, predicate ProjectParticipationOperationPredicate) (result ListProfileProjectsCompleteResult, err error) {
	items := make([]beta.ProjectParticipation, 0)

	resp, err := c.ListProfileProjects(ctx, id, options)
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

	result = ListProfileProjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
