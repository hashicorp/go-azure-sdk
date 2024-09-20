package chatoperation

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

type ListChatOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TeamsAsyncOperation
}

type ListChatOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TeamsAsyncOperation
}

type ListChatOperationsOperationOptions struct {
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

func DefaultListChatOperationsOperationOptions() ListChatOperationsOperationOptions {
	return ListChatOperationsOperationOptions{}
}

func (o ListChatOperationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListChatOperationsOperationOptions) ToOData() *odata.Query {
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

func (o ListChatOperationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListChatOperationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListChatOperationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListChatOperations - Get operations from users. A collection of all the Teams async operations that ran or are
// running on the chat. Nullable.
func (c ChatOperationClient) ListChatOperations(ctx context.Context, id beta.UserIdChatId, options ListChatOperationsOperationOptions) (result ListChatOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListChatOperationsCustomPager{},
		Path:          fmt.Sprintf("%s/operations", id.ID()),
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
		Values *[]beta.TeamsAsyncOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListChatOperationsComplete retrieves all the results into a single object
func (c ChatOperationClient) ListChatOperationsComplete(ctx context.Context, id beta.UserIdChatId, options ListChatOperationsOperationOptions) (ListChatOperationsCompleteResult, error) {
	return c.ListChatOperationsCompleteMatchingPredicate(ctx, id, options, TeamsAsyncOperationOperationPredicate{})
}

// ListChatOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatOperationClient) ListChatOperationsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdChatId, options ListChatOperationsOperationOptions, predicate TeamsAsyncOperationOperationPredicate) (result ListChatOperationsCompleteResult, err error) {
	items := make([]beta.TeamsAsyncOperation, 0)

	resp, err := c.ListChatOperations(ctx, id, options)
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

	result = ListChatOperationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
