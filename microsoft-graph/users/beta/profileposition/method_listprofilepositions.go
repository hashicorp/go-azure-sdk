package profileposition

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

type ListProfilePositionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WorkPosition
}

type ListProfilePositionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WorkPosition
}

type ListProfilePositionsOperationOptions struct {
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

func DefaultListProfilePositionsOperationOptions() ListProfilePositionsOperationOptions {
	return ListProfilePositionsOperationOptions{}
}

func (o ListProfilePositionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProfilePositionsOperationOptions) ToOData() *odata.Query {
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

func (o ListProfilePositionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProfilePositionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfilePositionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfilePositions - Get positions from users. Represents detailed information about work positions associated with
// a user's profile.
func (c ProfilePositionClient) ListProfilePositions(ctx context.Context, id beta.UserId, options ListProfilePositionsOperationOptions) (result ListProfilePositionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProfilePositionsCustomPager{},
		Path:          fmt.Sprintf("%s/profile/positions", id.ID()),
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
		Values *[]beta.WorkPosition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfilePositionsComplete retrieves all the results into a single object
func (c ProfilePositionClient) ListProfilePositionsComplete(ctx context.Context, id beta.UserId, options ListProfilePositionsOperationOptions) (ListProfilePositionsCompleteResult, error) {
	return c.ListProfilePositionsCompleteMatchingPredicate(ctx, id, options, WorkPositionOperationPredicate{})
}

// ListProfilePositionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfilePositionClient) ListProfilePositionsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListProfilePositionsOperationOptions, predicate WorkPositionOperationPredicate) (result ListProfilePositionsCompleteResult, err error) {
	items := make([]beta.WorkPosition, 0)

	resp, err := c.ListProfilePositions(ctx, id, options)
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

	result = ListProfilePositionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
