package chatmessage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForwardChatMessagesToChatsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActionResultPart
}

type ForwardChatMessagesToChatsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActionResultPart
}

type ForwardChatMessagesToChatsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultForwardChatMessagesToChatsOperationOptions() ForwardChatMessagesToChatsOperationOptions {
	return ForwardChatMessagesToChatsOperationOptions{}
}

func (o ForwardChatMessagesToChatsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ForwardChatMessagesToChatsOperationOptions) ToOData() *odata.Query {
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

func (o ForwardChatMessagesToChatsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ForwardChatMessagesToChatsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ForwardChatMessagesToChatsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ForwardChatMessagesToChats - Invoke action forwardToChat. Forward a chat message, a channel message, or a channel
// message reply to a chat.
func (c ChatMessageClient) ForwardChatMessagesToChats(ctx context.Context, id beta.MeChatId, input ForwardChatMessagesToChatsRequest, options ForwardChatMessagesToChatsOperationOptions) (result ForwardChatMessagesToChatsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ForwardChatMessagesToChatsCustomPager{},
		Path:          fmt.Sprintf("%s/messages/forwardToChat", id.ID()),
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

// ForwardChatMessagesToChatsComplete retrieves all the results into a single object
func (c ChatMessageClient) ForwardChatMessagesToChatsComplete(ctx context.Context, id beta.MeChatId, input ForwardChatMessagesToChatsRequest, options ForwardChatMessagesToChatsOperationOptions) (ForwardChatMessagesToChatsCompleteResult, error) {
	return c.ForwardChatMessagesToChatsCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// ForwardChatMessagesToChatsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatMessageClient) ForwardChatMessagesToChatsCompleteMatchingPredicate(ctx context.Context, id beta.MeChatId, input ForwardChatMessagesToChatsRequest, options ForwardChatMessagesToChatsOperationOptions, predicate ActionResultPartOperationPredicate) (result ForwardChatMessagesToChatsCompleteResult, err error) {
	items := make([]beta.ActionResultPart, 0)

	resp, err := c.ForwardChatMessagesToChats(ctx, id, input, options)
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

	result = ForwardChatMessagesToChatsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
