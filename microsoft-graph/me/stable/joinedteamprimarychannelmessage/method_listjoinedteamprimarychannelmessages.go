package joinedteamprimarychannelmessage

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

type ListJoinedTeamPrimaryChannelMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ChatMessage
}

type ListJoinedTeamPrimaryChannelMessagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ChatMessage
}

type ListJoinedTeamPrimaryChannelMessagesOperationOptions struct {
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

func DefaultListJoinedTeamPrimaryChannelMessagesOperationOptions() ListJoinedTeamPrimaryChannelMessagesOperationOptions {
	return ListJoinedTeamPrimaryChannelMessagesOperationOptions{}
}

func (o ListJoinedTeamPrimaryChannelMessagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamPrimaryChannelMessagesOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamPrimaryChannelMessagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamPrimaryChannelMessagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamPrimaryChannelMessagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamPrimaryChannelMessages - Get messages from me. A collection of all the messages in the channel. A
// navigation property. Nullable.
func (c JoinedTeamPrimaryChannelMessageClient) ListJoinedTeamPrimaryChannelMessages(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamPrimaryChannelMessagesOperationOptions) (result ListJoinedTeamPrimaryChannelMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamPrimaryChannelMessagesCustomPager{},
		Path:          fmt.Sprintf("%s/primaryChannel/messages", id.ID()),
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

// ListJoinedTeamPrimaryChannelMessagesComplete retrieves all the results into a single object
func (c JoinedTeamPrimaryChannelMessageClient) ListJoinedTeamPrimaryChannelMessagesComplete(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamPrimaryChannelMessagesOperationOptions) (ListJoinedTeamPrimaryChannelMessagesCompleteResult, error) {
	return c.ListJoinedTeamPrimaryChannelMessagesCompleteMatchingPredicate(ctx, id, options, ChatMessageOperationPredicate{})
}

// ListJoinedTeamPrimaryChannelMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPrimaryChannelMessageClient) ListJoinedTeamPrimaryChannelMessagesCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamId, options ListJoinedTeamPrimaryChannelMessagesOperationOptions, predicate ChatMessageOperationPredicate) (result ListJoinedTeamPrimaryChannelMessagesCompleteResult, err error) {
	items := make([]stable.ChatMessage, 0)

	resp, err := c.ListJoinedTeamPrimaryChannelMessages(ctx, id, options)
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

	result = ListJoinedTeamPrimaryChannelMessagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
