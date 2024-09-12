package mailfolderchildfoldermessage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMailFolderChildFolderMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Message
}

type ListMailFolderChildFolderMessagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Message
}

type ListMailFolderChildFolderMessagesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListMailFolderChildFolderMessagesOperationOptions() ListMailFolderChildFolderMessagesOperationOptions {
	return ListMailFolderChildFolderMessagesOperationOptions{}
}

func (o ListMailFolderChildFolderMessagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMailFolderChildFolderMessagesOperationOptions) ToOData() *odata.Query {
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

func (o ListMailFolderChildFolderMessagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMailFolderChildFolderMessagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFolderChildFolderMessagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolderChildFolderMessages - Get messages from users. The collection of messages in the mailFolder.
func (c MailFolderChildFolderMessageClient) ListMailFolderChildFolderMessages(ctx context.Context, id stable.UserIdMailFolderIdChildFolderId, options ListMailFolderChildFolderMessagesOperationOptions) (result ListMailFolderChildFolderMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMailFolderChildFolderMessagesCustomPager{},
		Path:          fmt.Sprintf("%s/messages", id.ID()),
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

// ListMailFolderChildFolderMessagesComplete retrieves all the results into a single object
func (c MailFolderChildFolderMessageClient) ListMailFolderChildFolderMessagesComplete(ctx context.Context, id stable.UserIdMailFolderIdChildFolderId, options ListMailFolderChildFolderMessagesOperationOptions) (ListMailFolderChildFolderMessagesCompleteResult, error) {
	return c.ListMailFolderChildFolderMessagesCompleteMatchingPredicate(ctx, id, options, MessageOperationPredicate{})
}

// ListMailFolderChildFolderMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderChildFolderMessageClient) ListMailFolderChildFolderMessagesCompleteMatchingPredicate(ctx context.Context, id stable.UserIdMailFolderIdChildFolderId, options ListMailFolderChildFolderMessagesOperationOptions, predicate MessageOperationPredicate) (result ListMailFolderChildFolderMessagesCompleteResult, err error) {
	items := make([]stable.Message, 0)

	resp, err := c.ListMailFolderChildFolderMessages(ctx, id, options)
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

	result = ListMailFolderChildFolderMessagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
