package mailfoldermessage

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

type ListMailFolderMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Message
}

type ListMailFolderMessagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Message
}

type ListMailFolderMessagesOperationOptions struct {
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

func DefaultListMailFolderMessagesOperationOptions() ListMailFolderMessagesOperationOptions {
	return ListMailFolderMessagesOperationOptions{}
}

func (o ListMailFolderMessagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMailFolderMessagesOperationOptions) ToOData() *odata.Query {
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

func (o ListMailFolderMessagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMailFolderMessagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFolderMessagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolderMessages - List messages. List all the messages in the specified user's mailbox, or those messages in a
// specified folder in the mailbox.
func (c MailFolderMessageClient) ListMailFolderMessages(ctx context.Context, id beta.MeMailFolderId, options ListMailFolderMessagesOperationOptions) (result ListMailFolderMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMailFolderMessagesCustomPager{},
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.Message, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalMessageImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.Message (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListMailFolderMessagesComplete retrieves all the results into a single object
func (c MailFolderMessageClient) ListMailFolderMessagesComplete(ctx context.Context, id beta.MeMailFolderId, options ListMailFolderMessagesOperationOptions) (ListMailFolderMessagesCompleteResult, error) {
	return c.ListMailFolderMessagesCompleteMatchingPredicate(ctx, id, options, MessageOperationPredicate{})
}

// ListMailFolderMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderMessageClient) ListMailFolderMessagesCompleteMatchingPredicate(ctx context.Context, id beta.MeMailFolderId, options ListMailFolderMessagesOperationOptions, predicate MessageOperationPredicate) (result ListMailFolderMessagesCompleteResult, err error) {
	items := make([]beta.Message, 0)

	resp, err := c.ListMailFolderMessages(ctx, id, options)
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

	result = ListMailFolderMessagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
