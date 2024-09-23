package chatmessagehostedcontent

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

type ListChatMessageHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ChatMessageHostedContent
}

type ListChatMessageHostedContentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ChatMessageHostedContent
}

type ListChatMessageHostedContentsOperationOptions struct {
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

func DefaultListChatMessageHostedContentsOperationOptions() ListChatMessageHostedContentsOperationOptions {
	return ListChatMessageHostedContentsOperationOptions{}
}

func (o ListChatMessageHostedContentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListChatMessageHostedContentsOperationOptions) ToOData() *odata.Query {
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

func (o ListChatMessageHostedContentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListChatMessageHostedContentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListChatMessageHostedContentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListChatMessageHostedContents - Get hostedContents from me. Content in a message hosted by Microsoft Teams - for
// example, images or code snippets.
func (c ChatMessageHostedContentClient) ListChatMessageHostedContents(ctx context.Context, id stable.MeChatIdMessageId, options ListChatMessageHostedContentsOperationOptions) (result ListChatMessageHostedContentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListChatMessageHostedContentsCustomPager{},
		Path:          fmt.Sprintf("%s/hostedContents", id.ID()),
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
		Values *[]stable.ChatMessageHostedContent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListChatMessageHostedContentsComplete retrieves all the results into a single object
func (c ChatMessageHostedContentClient) ListChatMessageHostedContentsComplete(ctx context.Context, id stable.MeChatIdMessageId, options ListChatMessageHostedContentsOperationOptions) (ListChatMessageHostedContentsCompleteResult, error) {
	return c.ListChatMessageHostedContentsCompleteMatchingPredicate(ctx, id, options, ChatMessageHostedContentOperationPredicate{})
}

// ListChatMessageHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatMessageHostedContentClient) ListChatMessageHostedContentsCompleteMatchingPredicate(ctx context.Context, id stable.MeChatIdMessageId, options ListChatMessageHostedContentsOperationOptions, predicate ChatMessageHostedContentOperationPredicate) (result ListChatMessageHostedContentsCompleteResult, err error) {
	items := make([]stable.ChatMessageHostedContent, 0)

	resp, err := c.ListChatMessageHostedContents(ctx, id, options)
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

	result = ListChatMessageHostedContentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
