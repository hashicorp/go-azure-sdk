package person

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

type ListPeopleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Person
}

type ListPeopleCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Person
}

type ListPeopleOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListPeopleOperationOptions() ListPeopleOperationOptions {
	return ListPeopleOperationOptions{}
}

func (o ListPeopleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPeopleOperationOptions) ToOData() *odata.Query {
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

func (o ListPeopleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPeopleCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPeopleCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPeople - List people. Retrieve a collection of person objects ordered by their relevance to the user, which is
// determined by the user's communication and collaboration patterns, and business relationships. You can get this
// information via the People API. For examples, see the Examples section and the article Use the People API to get
// information about the people most relevant to you.
func (c PersonClient) ListPeople(ctx context.Context, options ListPeopleOperationOptions) (result ListPeopleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPeopleCustomPager{},
		Path:          "/me/people",
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
		Values *[]stable.Person `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPeopleComplete retrieves all the results into a single object
func (c PersonClient) ListPeopleComplete(ctx context.Context, options ListPeopleOperationOptions) (ListPeopleCompleteResult, error) {
	return c.ListPeopleCompleteMatchingPredicate(ctx, options, PersonOperationPredicate{})
}

// ListPeopleCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PersonClient) ListPeopleCompleteMatchingPredicate(ctx context.Context, options ListPeopleOperationOptions, predicate PersonOperationPredicate) (result ListPeopleCompleteResult, err error) {
	items := make([]stable.Person, 0)

	resp, err := c.ListPeople(ctx, options)
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

	result = ListPeopleCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
