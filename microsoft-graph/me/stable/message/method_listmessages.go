package message

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Message
}

type ListMessagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Message
}

type ListMessagesOperationOptions struct {
	Count                 *bool
	Expand                *odata.Expand
	Filter                *string
	IncludeHiddenMessages *string
	Metadata              *odata.Metadata
	OrderBy               *odata.OrderBy
	RetryFunc             client.RequestRetryFunc
	Search                *string
	Select                *[]string
	Skip                  *int64
	Top                   *int64
}

func DefaultListMessagesOperationOptions() ListMessagesOperationOptions {
	return ListMessagesOperationOptions{}
}

func (o ListMessagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMessagesOperationOptions) ToOData() *odata.Query {
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

func (o ListMessagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IncludeHiddenMessages != nil {
		out.Append("includeHiddenMessages", fmt.Sprintf("%v", *o.IncludeHiddenMessages))
	}
	return &out
}

type ListMessagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMessagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMessages - List messages. Get the messages in the signed-in user's mailbox (including the Deleted Items and
// Clutter folders). Depending on the page size and mailbox data, getting messages from a mailbox can incur multiple
// requests. The default page size is 10 messages. Use $top to customize the page size, within the range of 1 and 1000.
// To improve the operation response time, use $select to specify the exact properties you need; see example 1 below.
// Fine-tune the values for $select and $top, especially when you must use a larger page size, as returning a page with
// hundreds of messages each with a full response payload may trigger the gateway timeout (HTTP 504). To get the next
// page of messages, simply apply the entire URL returned in @odata.nextLink to the next get-messages request. This URL
// includes any query parameters you may have specified in the initial request. Do not try to extract the $skip value
// from the @odata.nextLink URL to manipulate responses. This API uses the $skip value to keep count of all the items it
// has gone through in the user's mailbox to return a page of message-type items. It's therefore possible that even in
// the initial response, the $skip value is larger than the page size. For more information, see Paging Microsoft Graph
// data in your app. Currently, this operation returns message bodies in only HTML format. There are two scenarios where
// an app can get messages in another user's mail folder
func (c MessageClient) ListMessages(ctx context.Context, options ListMessagesOperationOptions) (result ListMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMessagesCustomPager{},
		Path:          "/me/messages",
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

	temp := make([]stable.Message, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalMessageImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.Message (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListMessagesComplete retrieves all the results into a single object
func (c MessageClient) ListMessagesComplete(ctx context.Context, options ListMessagesOperationOptions) (ListMessagesCompleteResult, error) {
	return c.ListMessagesCompleteMatchingPredicate(ctx, options, MessageOperationPredicate{})
}

// ListMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MessageClient) ListMessagesCompleteMatchingPredicate(ctx context.Context, options ListMessagesOperationOptions, predicate MessageOperationPredicate) (result ListMessagesCompleteResult, err error) {
	items := make([]stable.Message, 0)

	resp, err := c.ListMessages(ctx, options)
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

	result = ListMessagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
