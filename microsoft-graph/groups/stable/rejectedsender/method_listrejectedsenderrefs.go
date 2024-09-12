package rejectedsender

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

type ListRejectedSenderRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListRejectedSenderRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListRejectedSenderRefsOperationOptions struct {
	Count   *bool
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Skip    *int64
	Top     *int64
}

func DefaultListRejectedSenderRefsOperationOptions() ListRejectedSenderRefsOperationOptions {
	return ListRejectedSenderRefsOperationOptions{}
}

func (o ListRejectedSenderRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRejectedSenderRefsOperationOptions) ToOData() *odata.Query {
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

func (o ListRejectedSenderRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRejectedSenderRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRejectedSenderRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRejectedSenderRefs - List rejectedSenders. Users in the rejected senders list can't post to conversations of the
// group (identified in the GET request URL). Make sure you don't specify the same user or group in the rejected senders
// and accepted senders lists, otherwise you get an error.
func (c RejectedSenderClient) ListRejectedSenderRefs(ctx context.Context, id stable.GroupId, options ListRejectedSenderRefsOperationOptions) (result ListRejectedSenderRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRejectedSenderRefsCustomPager{},
		Path:          fmt.Sprintf("%s/rejectedSenders/$ref", id.ID()),
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

// ListRejectedSenderRefsComplete retrieves all the results into a single object
func (c RejectedSenderClient) ListRejectedSenderRefsComplete(ctx context.Context, id stable.GroupId, options ListRejectedSenderRefsOperationOptions) (ListRejectedSenderRefsCompleteResult, error) {
	return c.ListRejectedSenderRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListRejectedSenderRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RejectedSenderClient) ListRejectedSenderRefsCompleteMatchingPredicate(ctx context.Context, id stable.GroupId, options ListRejectedSenderRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListRejectedSenderRefsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListRejectedSenderRefs(ctx, id, options)
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

	result = ListRejectedSenderRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
