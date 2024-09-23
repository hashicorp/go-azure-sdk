package teamprimarychannelmessagereplyhostedcontent

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

type ListTeamPrimaryChannelMessageReplyHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ChatMessageHostedContent
}

type ListTeamPrimaryChannelMessageReplyHostedContentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ChatMessageHostedContent
}

type ListTeamPrimaryChannelMessageReplyHostedContentsOperationOptions struct {
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

func DefaultListTeamPrimaryChannelMessageReplyHostedContentsOperationOptions() ListTeamPrimaryChannelMessageReplyHostedContentsOperationOptions {
	return ListTeamPrimaryChannelMessageReplyHostedContentsOperationOptions{}
}

func (o ListTeamPrimaryChannelMessageReplyHostedContentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamPrimaryChannelMessageReplyHostedContentsOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamPrimaryChannelMessageReplyHostedContentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamPrimaryChannelMessageReplyHostedContentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamPrimaryChannelMessageReplyHostedContentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamPrimaryChannelMessageReplyHostedContents - Get hostedContents from groups. Content in a message hosted by
// Microsoft Teams - for example, images or code snippets.
func (c TeamPrimaryChannelMessageReplyHostedContentClient) ListTeamPrimaryChannelMessageReplyHostedContents(ctx context.Context, id beta.GroupIdTeamPrimaryChannelMessageIdReplyId, options ListTeamPrimaryChannelMessageReplyHostedContentsOperationOptions) (result ListTeamPrimaryChannelMessageReplyHostedContentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamPrimaryChannelMessageReplyHostedContentsCustomPager{},
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
		Values *[]beta.ChatMessageHostedContent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamPrimaryChannelMessageReplyHostedContentsComplete retrieves all the results into a single object
func (c TeamPrimaryChannelMessageReplyHostedContentClient) ListTeamPrimaryChannelMessageReplyHostedContentsComplete(ctx context.Context, id beta.GroupIdTeamPrimaryChannelMessageIdReplyId, options ListTeamPrimaryChannelMessageReplyHostedContentsOperationOptions) (ListTeamPrimaryChannelMessageReplyHostedContentsCompleteResult, error) {
	return c.ListTeamPrimaryChannelMessageReplyHostedContentsCompleteMatchingPredicate(ctx, id, options, ChatMessageHostedContentOperationPredicate{})
}

// ListTeamPrimaryChannelMessageReplyHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamPrimaryChannelMessageReplyHostedContentClient) ListTeamPrimaryChannelMessageReplyHostedContentsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdTeamPrimaryChannelMessageIdReplyId, options ListTeamPrimaryChannelMessageReplyHostedContentsOperationOptions, predicate ChatMessageHostedContentOperationPredicate) (result ListTeamPrimaryChannelMessageReplyHostedContentsCompleteResult, err error) {
	items := make([]beta.ChatMessageHostedContent, 0)

	resp, err := c.ListTeamPrimaryChannelMessageReplyHostedContents(ctx, id, options)
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

	result = ListTeamPrimaryChannelMessageReplyHostedContentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
