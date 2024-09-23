package joinedteamprimarychannelmessagereply

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

type ListJoinedTeamPrimaryChannelMessageRepliesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ChatMessage
}

type ListJoinedTeamPrimaryChannelMessageRepliesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ChatMessage
}

type ListJoinedTeamPrimaryChannelMessageRepliesOperationOptions struct {
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

func DefaultListJoinedTeamPrimaryChannelMessageRepliesOperationOptions() ListJoinedTeamPrimaryChannelMessageRepliesOperationOptions {
	return ListJoinedTeamPrimaryChannelMessageRepliesOperationOptions{}
}

func (o ListJoinedTeamPrimaryChannelMessageRepliesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamPrimaryChannelMessageRepliesOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamPrimaryChannelMessageRepliesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamPrimaryChannelMessageRepliesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamPrimaryChannelMessageRepliesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamPrimaryChannelMessageReplies - Get replies from me. Replies for a specified message. Supports $expand
// for channel messages.
func (c JoinedTeamPrimaryChannelMessageReplyClient) ListJoinedTeamPrimaryChannelMessageReplies(ctx context.Context, id stable.MeJoinedTeamIdPrimaryChannelMessageId, options ListJoinedTeamPrimaryChannelMessageRepliesOperationOptions) (result ListJoinedTeamPrimaryChannelMessageRepliesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamPrimaryChannelMessageRepliesCustomPager{},
		Path:          fmt.Sprintf("%s/replies", id.ID()),
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
		Values *[]stable.ChatMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamPrimaryChannelMessageRepliesComplete retrieves all the results into a single object
func (c JoinedTeamPrimaryChannelMessageReplyClient) ListJoinedTeamPrimaryChannelMessageRepliesComplete(ctx context.Context, id stable.MeJoinedTeamIdPrimaryChannelMessageId, options ListJoinedTeamPrimaryChannelMessageRepliesOperationOptions) (ListJoinedTeamPrimaryChannelMessageRepliesCompleteResult, error) {
	return c.ListJoinedTeamPrimaryChannelMessageRepliesCompleteMatchingPredicate(ctx, id, options, ChatMessageOperationPredicate{})
}

// ListJoinedTeamPrimaryChannelMessageRepliesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPrimaryChannelMessageReplyClient) ListJoinedTeamPrimaryChannelMessageRepliesCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamIdPrimaryChannelMessageId, options ListJoinedTeamPrimaryChannelMessageRepliesOperationOptions, predicate ChatMessageOperationPredicate) (result ListJoinedTeamPrimaryChannelMessageRepliesCompleteResult, err error) {
	items := make([]stable.ChatMessage, 0)

	resp, err := c.ListJoinedTeamPrimaryChannelMessageReplies(ctx, id, options)
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

	result = ListJoinedTeamPrimaryChannelMessageRepliesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
