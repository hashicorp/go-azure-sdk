package teamchannelmessage

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

type ForwardTeamChannelMessagesToChatsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ActionResultPart
}

type ForwardTeamChannelMessagesToChatsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ActionResultPart
}

type ForwardTeamChannelMessagesToChatsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultForwardTeamChannelMessagesToChatsOperationOptions() ForwardTeamChannelMessagesToChatsOperationOptions {
	return ForwardTeamChannelMessagesToChatsOperationOptions{}
}

func (o ForwardTeamChannelMessagesToChatsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ForwardTeamChannelMessagesToChatsOperationOptions) ToOData() *odata.Query {
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

func (o ForwardTeamChannelMessagesToChatsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ForwardTeamChannelMessagesToChatsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ForwardTeamChannelMessagesToChatsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ForwardTeamChannelMessagesToChats - Invoke action forwardToChat. Forward a chat message, a channel message, or a
// channel message reply to a chat.
func (c TeamChannelMessageClient) ForwardTeamChannelMessagesToChats(ctx context.Context, id beta.GroupIdTeamChannelId, input ForwardTeamChannelMessagesToChatsRequest, options ForwardTeamChannelMessagesToChatsOperationOptions) (result ForwardTeamChannelMessagesToChatsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ForwardTeamChannelMessagesToChatsCustomPager{},
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

// ForwardTeamChannelMessagesToChatsComplete retrieves all the results into a single object
func (c TeamChannelMessageClient) ForwardTeamChannelMessagesToChatsComplete(ctx context.Context, id beta.GroupIdTeamChannelId, input ForwardTeamChannelMessagesToChatsRequest, options ForwardTeamChannelMessagesToChatsOperationOptions) (ForwardTeamChannelMessagesToChatsCompleteResult, error) {
	return c.ForwardTeamChannelMessagesToChatsCompleteMatchingPredicate(ctx, id, input, options, ActionResultPartOperationPredicate{})
}

// ForwardTeamChannelMessagesToChatsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamChannelMessageClient) ForwardTeamChannelMessagesToChatsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdTeamChannelId, input ForwardTeamChannelMessagesToChatsRequest, options ForwardTeamChannelMessagesToChatsOperationOptions, predicate ActionResultPartOperationPredicate) (result ForwardTeamChannelMessagesToChatsCompleteResult, err error) {
	items := make([]beta.ActionResultPart, 0)

	resp, err := c.ForwardTeamChannelMessagesToChats(ctx, id, input, options)
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

	result = ForwardTeamChannelMessagesToChatsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
