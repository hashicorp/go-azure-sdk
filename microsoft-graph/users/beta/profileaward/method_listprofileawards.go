package profileaward

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

type ListProfileAwardsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PersonAward
}

type ListProfileAwardsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PersonAward
}

type ListProfileAwardsOperationOptions struct {
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

func DefaultListProfileAwardsOperationOptions() ListProfileAwardsOperationOptions {
	return ListProfileAwardsOperationOptions{}
}

func (o ListProfileAwardsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProfileAwardsOperationOptions) ToOData() *odata.Query {
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

func (o ListProfileAwardsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProfileAwardsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileAwardsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileAwards - Get awards from users. Represents the details of awards or honors associated with a person.
func (c ProfileAwardClient) ListProfileAwards(ctx context.Context, id beta.UserId, options ListProfileAwardsOperationOptions) (result ListProfileAwardsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProfileAwardsCustomPager{},
		Path:          fmt.Sprintf("%s/profile/awards", id.ID()),
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
		Values *[]beta.PersonAward `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileAwardsComplete retrieves all the results into a single object
func (c ProfileAwardClient) ListProfileAwardsComplete(ctx context.Context, id beta.UserId, options ListProfileAwardsOperationOptions) (ListProfileAwardsCompleteResult, error) {
	return c.ListProfileAwardsCompleteMatchingPredicate(ctx, id, options, PersonAwardOperationPredicate{})
}

// ListProfileAwardsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileAwardClient) ListProfileAwardsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListProfileAwardsOperationOptions, predicate PersonAwardOperationPredicate) (result ListProfileAwardsCompleteResult, err error) {
	items := make([]beta.PersonAward, 0)

	resp, err := c.ListProfileAwards(ctx, id, options)
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

	result = ListProfileAwardsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
