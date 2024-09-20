package acceptedsender

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

type ListAcceptedSendersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListAcceptedSendersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListAcceptedSendersOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListAcceptedSendersOperationOptions() ListAcceptedSendersOperationOptions {
	return ListAcceptedSendersOperationOptions{}
}

func (o ListAcceptedSendersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAcceptedSendersOperationOptions) ToOData() *odata.Query {
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

func (o ListAcceptedSendersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAcceptedSendersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAcceptedSendersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAcceptedSenders - List acceptedSenders. Users in the accepted senders list can post to conversations of the group
// (identified in the GET request URL). Make sure you do not specify the same user or group in the accepted senders and
// rejected senders lists, otherwise you will get an error.
func (c AcceptedSenderClient) ListAcceptedSenders(ctx context.Context, id stable.GroupId, options ListAcceptedSendersOperationOptions) (result ListAcceptedSendersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAcceptedSendersCustomPager{},
		Path:          fmt.Sprintf("%s/acceptedSenders", id.ID()),
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

	temp := make([]stable.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListAcceptedSendersComplete retrieves all the results into a single object
func (c AcceptedSenderClient) ListAcceptedSendersComplete(ctx context.Context, id stable.GroupId, options ListAcceptedSendersOperationOptions) (ListAcceptedSendersCompleteResult, error) {
	return c.ListAcceptedSendersCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListAcceptedSendersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AcceptedSenderClient) ListAcceptedSendersCompleteMatchingPredicate(ctx context.Context, id stable.GroupId, options ListAcceptedSendersOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListAcceptedSendersCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListAcceptedSenders(ctx, id, options)
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

	result = ListAcceptedSendersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
