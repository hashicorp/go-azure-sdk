package onenotenotebook

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOnenoteNotebooksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Notebook
}

type ListOnenoteNotebooksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Notebook
}

type ListOnenoteNotebooksOperationOptions struct {
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

func DefaultListOnenoteNotebooksOperationOptions() ListOnenoteNotebooksOperationOptions {
	return ListOnenoteNotebooksOperationOptions{}
}

func (o ListOnenoteNotebooksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnenoteNotebooksOperationOptions) ToOData() *odata.Query {
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

func (o ListOnenoteNotebooksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnenoteNotebooksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnenoteNotebooksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnenoteNotebooks - Get notebooks from groups. The collection of OneNote notebooks that the user or group owns.
// Read-only. Nullable.
func (c OnenoteNotebookClient) ListOnenoteNotebooks(ctx context.Context, id beta.GroupId, options ListOnenoteNotebooksOperationOptions) (result ListOnenoteNotebooksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnenoteNotebooksCustomPager{},
		Path:          fmt.Sprintf("%s/onenote/notebooks", id.ID()),
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
		Values *[]beta.Notebook `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnenoteNotebooksComplete retrieves all the results into a single object
func (c OnenoteNotebookClient) ListOnenoteNotebooksComplete(ctx context.Context, id beta.GroupId, options ListOnenoteNotebooksOperationOptions) (ListOnenoteNotebooksCompleteResult, error) {
	return c.ListOnenoteNotebooksCompleteMatchingPredicate(ctx, id, options, NotebookOperationPredicate{})
}

// ListOnenoteNotebooksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnenoteNotebookClient) ListOnenoteNotebooksCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListOnenoteNotebooksOperationOptions, predicate NotebookOperationPredicate) (result ListOnenoteNotebooksCompleteResult, err error) {
	items := make([]beta.Notebook, 0)

	resp, err := c.ListOnenoteNotebooks(ctx, id, options)
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

	result = ListOnenoteNotebooksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
