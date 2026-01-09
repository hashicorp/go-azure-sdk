package joinedteamprimarychannelmessagereplyhostedcontent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ChatMessageHostedContent
}

type ListJoinedTeamPrimaryChannelMessageReplyHostedContentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ChatMessageHostedContent
}

type ListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationOptions struct {
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

func DefaultListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationOptions() ListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationOptions {
	return ListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationOptions{}
}

func (o ListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamPrimaryChannelMessageReplyHostedContentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamPrimaryChannelMessageReplyHostedContentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamPrimaryChannelMessageReplyHostedContents - Get hostedContents from users. Content in a message hosted
// by Microsoft Teams - for example, images or code snippets.
func (c JoinedTeamPrimaryChannelMessageReplyHostedContentClient) ListJoinedTeamPrimaryChannelMessageReplyHostedContents(ctx context.Context, id stable.UserIdJoinedTeamIdPrimaryChannelMessageIdReplyId, options ListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationOptions) (result ListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamPrimaryChannelMessageReplyHostedContentsCustomPager{},
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

// ListJoinedTeamPrimaryChannelMessageReplyHostedContentsComplete retrieves all the results into a single object
func (c JoinedTeamPrimaryChannelMessageReplyHostedContentClient) ListJoinedTeamPrimaryChannelMessageReplyHostedContentsComplete(ctx context.Context, id stable.UserIdJoinedTeamIdPrimaryChannelMessageIdReplyId, options ListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationOptions) (ListJoinedTeamPrimaryChannelMessageReplyHostedContentsCompleteResult, error) {
	return c.ListJoinedTeamPrimaryChannelMessageReplyHostedContentsCompleteMatchingPredicate(ctx, id, options, ChatMessageHostedContentOperationPredicate{})
}

// ListJoinedTeamPrimaryChannelMessageReplyHostedContentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPrimaryChannelMessageReplyHostedContentClient) ListJoinedTeamPrimaryChannelMessageReplyHostedContentsCompleteMatchingPredicate(ctx context.Context, id stable.UserIdJoinedTeamIdPrimaryChannelMessageIdReplyId, options ListJoinedTeamPrimaryChannelMessageReplyHostedContentsOperationOptions, predicate ChatMessageHostedContentOperationPredicate) (result ListJoinedTeamPrimaryChannelMessageReplyHostedContentsCompleteResult, err error) {
	items := make([]stable.ChatMessageHostedContent, 0)

	resp, err := c.ListJoinedTeamPrimaryChannelMessageReplyHostedContents(ctx, id, options)
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

	result = ListJoinedTeamPrimaryChannelMessageReplyHostedContentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
