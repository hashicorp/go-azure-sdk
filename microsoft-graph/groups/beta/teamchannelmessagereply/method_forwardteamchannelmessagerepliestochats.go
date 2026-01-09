package teamchannelmessagereply

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForwardTeamChannelMessageRepliesToChatsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActionResultPart
}

type ForwardTeamChannelMessageRepliesToChatsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActionResultPart
}

type ForwardTeamChannelMessageRepliesToChatsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultForwardTeamChannelMessageRepliesToChatsOperationOptions() ForwardTeamChannelMessageRepliesToChatsOperationOptions {
	return ForwardTeamChannelMessageRepliesToChatsOperationOptions{}
}

func (o ForwardTeamChannelMessageRepliesToChatsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ForwardTeamChannelMessageRepliesToChatsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ForwardTeamChannelMessageRepliesToChatsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ForwardTeamChannelMessageRepliesToChatsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ForwardTeamChannelMessageRepliesToChatsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ForwardTeamChannelMessageRepliesToChats - Invoke action forwardToChat. Forward a chat message, a channel message, or
// a channel message reply to a chat.
func (c TeamChannelMessageReplyClient) ForwardTeamChannelMessageRepliesToChats(ctx context.Context, id beta.GroupIdTeamChannelIdMessageId, input ForwardTeamChannelMessageRepliesToChatsRequest, options ForwardTeamChannelMessageRepliesToChatsOperationOptions) (result ForwardTeamChannelMessageRepliesToChatsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ForwardTeamChannelMessageRepliesToChatsCustomPager{},
		Path:          fmt.Sprintf("%s/replies/forwardToChat", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.ActionResultPart, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalActionResultPartImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.ActionResultPart (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ForwardTeamChannelMessageRepliesToChatsComplete retrieves all the results into a single object
func (c TeamChannelMessageReplyClient) ForwardTeamChannelMessageRepliesToChatsComplete(ctx context.Context, id beta.GroupIdTeamChannelIdMessageId, input ForwardTeamChannelMessageRepliesToChatsRequest, options ForwardTeamChannelMessageRepliesToChatsOperationOptions) (ForwardTeamChannelMessageRepliesToChatsCompleteResult, error) {
	return c.ForwardTeamChannelMessageRepliesToChatsCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// ForwardTeamChannelMessageRepliesToChatsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamChannelMessageReplyClient) ForwardTeamChannelMessageRepliesToChatsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdTeamChannelIdMessageId, input ForwardTeamChannelMessageRepliesToChatsRequest, options ForwardTeamChannelMessageRepliesToChatsOperationOptions, predicate ActionResultPartOperationPredicate) (result ForwardTeamChannelMessageRepliesToChatsCompleteResult, err error) {
	items := make([]beta.ActionResultPart, 0)

	resp, err := c.ForwardTeamChannelMessageRepliesToChats(ctx, id, input, options)
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

	result = ForwardTeamChannelMessageRepliesToChatsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
