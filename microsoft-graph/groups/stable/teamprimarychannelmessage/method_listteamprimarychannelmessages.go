package teamprimarychannelmessage

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

type ListTeamPrimaryChannelMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ChatMessage
}

type ListTeamPrimaryChannelMessagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ChatMessage
}

type ListTeamPrimaryChannelMessagesOperationOptions struct {
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

func DefaultListTeamPrimaryChannelMessagesOperationOptions() ListTeamPrimaryChannelMessagesOperationOptions {
	return ListTeamPrimaryChannelMessagesOperationOptions{}
}

func (o ListTeamPrimaryChannelMessagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamPrimaryChannelMessagesOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamPrimaryChannelMessagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamPrimaryChannelMessagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamPrimaryChannelMessagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamPrimaryChannelMessages - Get messages from groups. A collection of all the messages in the channel. A
// navigation property. Nullable.
func (c TeamPrimaryChannelMessageClient) ListTeamPrimaryChannelMessages(ctx context.Context, id stable.GroupId, options ListTeamPrimaryChannelMessagesOperationOptions) (result ListTeamPrimaryChannelMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamPrimaryChannelMessagesCustomPager{},
		Path:          fmt.Sprintf("%s/team/primaryChannel/messages", id.ID()),
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

// ListTeamPrimaryChannelMessagesComplete retrieves all the results into a single object
func (c TeamPrimaryChannelMessageClient) ListTeamPrimaryChannelMessagesComplete(ctx context.Context, id stable.GroupId, options ListTeamPrimaryChannelMessagesOperationOptions) (ListTeamPrimaryChannelMessagesCompleteResult, error) {
	return c.ListTeamPrimaryChannelMessagesCompleteMatchingPredicate(ctx, id, options, ChatMessageOperationPredicate{})
}

// ListTeamPrimaryChannelMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamPrimaryChannelMessageClient) ListTeamPrimaryChannelMessagesCompleteMatchingPredicate(ctx context.Context, id stable.GroupId, options ListTeamPrimaryChannelMessagesOperationOptions, predicate ChatMessageOperationPredicate) (result ListTeamPrimaryChannelMessagesCompleteResult, err error) {
	items := make([]stable.ChatMessage, 0)

	resp, err := c.ListTeamPrimaryChannelMessages(ctx, id, options)
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

	result = ListTeamPrimaryChannelMessagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
