package profileanniversary

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

type ListProfileAnniversariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PersonAnnualEvent
}

type ListProfileAnniversariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PersonAnnualEvent
}

type ListProfileAnniversariesOperationOptions struct {
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

func DefaultListProfileAnniversariesOperationOptions() ListProfileAnniversariesOperationOptions {
	return ListProfileAnniversariesOperationOptions{}
}

func (o ListProfileAnniversariesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProfileAnniversariesOperationOptions) ToOData() *odata.Query {
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

func (o ListProfileAnniversariesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProfileAnniversariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileAnniversariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileAnniversaries - Get anniversaries from users. Represents the details of meaningful dates associated with a
// person.
func (c ProfileAnniversaryClient) ListProfileAnniversaries(ctx context.Context, id beta.UserId, options ListProfileAnniversariesOperationOptions) (result ListProfileAnniversariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProfileAnniversariesCustomPager{},
		Path:          fmt.Sprintf("%s/profile/anniversaries", id.ID()),
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
		Values *[]beta.PersonAnnualEvent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileAnniversariesComplete retrieves all the results into a single object
func (c ProfileAnniversaryClient) ListProfileAnniversariesComplete(ctx context.Context, id beta.UserId, options ListProfileAnniversariesOperationOptions) (ListProfileAnniversariesCompleteResult, error) {
	return c.ListProfileAnniversariesCompleteMatchingPredicate(ctx, id, options, PersonAnnualEventOperationPredicate{})
}

// ListProfileAnniversariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileAnniversaryClient) ListProfileAnniversariesCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListProfileAnniversariesOperationOptions, predicate PersonAnnualEventOperationPredicate) (result ListProfileAnniversariesCompleteResult, err error) {
	items := make([]beta.PersonAnnualEvent, 0)

	resp, err := c.ListProfileAnniversaries(ctx, id, options)
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

	result = ListProfileAnniversariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
