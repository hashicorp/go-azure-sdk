package teamchannelmessage

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

type ListTeamChannelMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ChatMessage
}

type ListTeamChannelMessagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ChatMessage
}

type ListTeamChannelMessagesOperationOptions struct {
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

func DefaultListTeamChannelMessagesOperationOptions() ListTeamChannelMessagesOperationOptions {
	return ListTeamChannelMessagesOperationOptions{}
}

func (o ListTeamChannelMessagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamChannelMessagesOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamChannelMessagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamChannelMessagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamChannelMessagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamChannelMessages - Get messages from groups. A collection of all the messages in the channel. Nullable.
func (c TeamChannelMessageClient) ListTeamChannelMessages(ctx context.Context, id beta.GroupIdTeamChannelId, options ListTeamChannelMessagesOperationOptions) (result ListTeamChannelMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamChannelMessagesCustomPager{},
		Path:          fmt.Sprintf("%s/messages", id.ID()),
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
		Values *[]beta.ChatMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamChannelMessagesComplete retrieves all the results into a single object
func (c TeamChannelMessageClient) ListTeamChannelMessagesComplete(ctx context.Context, id beta.GroupIdTeamChannelId, options ListTeamChannelMessagesOperationOptions) (ListTeamChannelMessagesCompleteResult, error) {
	return c.ListTeamChannelMessagesCompleteMatchingPredicate(ctx, id, options, ChatMessageOperationPredicate{})
}

// ListTeamChannelMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamChannelMessageClient) ListTeamChannelMessagesCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdTeamChannelId, options ListTeamChannelMessagesOperationOptions, predicate ChatMessageOperationPredicate) (result ListTeamChannelMessagesCompleteResult, err error) {
	items := make([]beta.ChatMessage, 0)

	resp, err := c.ListTeamChannelMessages(ctx, id, options)
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

	result = ListTeamChannelMessagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
