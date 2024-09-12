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

type ListAcceptedSenderRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListAcceptedSenderRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListAcceptedSenderRefsOperationOptions struct {
	Count   *bool
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Skip    *int64
	Top     *int64
}

func DefaultListAcceptedSenderRefsOperationOptions() ListAcceptedSenderRefsOperationOptions {
	return ListAcceptedSenderRefsOperationOptions{}
}

func (o ListAcceptedSenderRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAcceptedSenderRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
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
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListAcceptedSenderRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAcceptedSenderRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAcceptedSenderRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAcceptedSenderRefs - List acceptedSenders. Users in the accepted senders list can post to conversations of the
// group (identified in the GET request URL). Make sure you do not specify the same user or group in the accepted
// senders and rejected senders lists, otherwise you will get an error.
func (c AcceptedSenderClient) ListAcceptedSenderRefs(ctx context.Context, id stable.GroupId, options ListAcceptedSenderRefsOperationOptions) (result ListAcceptedSenderRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAcceptedSenderRefsCustomPager{},
		Path:          fmt.Sprintf("%s/acceptedSenders/$ref", id.ID()),
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

// ListAcceptedSenderRefsComplete retrieves all the results into a single object
func (c AcceptedSenderClient) ListAcceptedSenderRefsComplete(ctx context.Context, id stable.GroupId, options ListAcceptedSenderRefsOperationOptions) (ListAcceptedSenderRefsCompleteResult, error) {
	return c.ListAcceptedSenderRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListAcceptedSenderRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AcceptedSenderClient) ListAcceptedSenderRefsCompleteMatchingPredicate(ctx context.Context, id stable.GroupId, options ListAcceptedSenderRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListAcceptedSenderRefsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListAcceptedSenderRefs(ctx, id, options)
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

	result = ListAcceptedSenderRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
