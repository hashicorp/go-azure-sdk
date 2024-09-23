package chattab

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

type ListChatTabsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TeamsTab
}

type ListChatTabsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TeamsTab
}

type ListChatTabsOperationOptions struct {
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

func DefaultListChatTabsOperationOptions() ListChatTabsOperationOptions {
	return ListChatTabsOperationOptions{}
}

func (o ListChatTabsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListChatTabsOperationOptions) ToOData() *odata.Query {
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

func (o ListChatTabsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListChatTabsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListChatTabsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListChatTabs - Get tabs from users. A collection of all the tabs in the chat. Nullable.
func (c ChatTabClient) ListChatTabs(ctx context.Context, id beta.UserIdChatId, options ListChatTabsOperationOptions) (result ListChatTabsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListChatTabsCustomPager{},
		Path:          fmt.Sprintf("%s/tabs", id.ID()),
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
		Values *[]beta.TeamsTab `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListChatTabsComplete retrieves all the results into a single object
func (c ChatTabClient) ListChatTabsComplete(ctx context.Context, id beta.UserIdChatId, options ListChatTabsOperationOptions) (ListChatTabsCompleteResult, error) {
	return c.ListChatTabsCompleteMatchingPredicate(ctx, id, options, TeamsTabOperationPredicate{})
}

// ListChatTabsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatTabClient) ListChatTabsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdChatId, options ListChatTabsOperationOptions, predicate TeamsTabOperationPredicate) (result ListChatTabsCompleteResult, err error) {
	items := make([]beta.TeamsTab, 0)

	resp, err := c.ListChatTabs(ctx, id, options)
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

	result = ListChatTabsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
