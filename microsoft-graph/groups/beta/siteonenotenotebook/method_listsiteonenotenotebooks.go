package siteonenotenotebook

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

type ListSiteOnenoteNotebooksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Notebook
}

type ListSiteOnenoteNotebooksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Notebook
}

type ListSiteOnenoteNotebooksOperationOptions struct {
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

func DefaultListSiteOnenoteNotebooksOperationOptions() ListSiteOnenoteNotebooksOperationOptions {
	return ListSiteOnenoteNotebooksOperationOptions{}
}

func (o ListSiteOnenoteNotebooksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteOnenoteNotebooksOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteOnenoteNotebooksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteOnenoteNotebooksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteOnenoteNotebooksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteOnenoteNotebooks - Get notebooks from groups. The collection of OneNote notebooks that the user or group
// owns. Read-only. Nullable.
func (c SiteOnenoteNotebookClient) ListSiteOnenoteNotebooks(ctx context.Context, id beta.GroupIdSiteId, options ListSiteOnenoteNotebooksOperationOptions) (result ListSiteOnenoteNotebooksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteOnenoteNotebooksCustomPager{},
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

// ListSiteOnenoteNotebooksComplete retrieves all the results into a single object
func (c SiteOnenoteNotebookClient) ListSiteOnenoteNotebooksComplete(ctx context.Context, id beta.GroupIdSiteId, options ListSiteOnenoteNotebooksOperationOptions) (ListSiteOnenoteNotebooksCompleteResult, error) {
	return c.ListSiteOnenoteNotebooksCompleteMatchingPredicate(ctx, id, options, NotebookOperationPredicate{})
}

// ListSiteOnenoteNotebooksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteOnenoteNotebookClient) ListSiteOnenoteNotebooksCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteId, options ListSiteOnenoteNotebooksOperationOptions, predicate NotebookOperationPredicate) (result ListSiteOnenoteNotebooksCompleteResult, err error) {
	items := make([]beta.Notebook, 0)

	resp, err := c.ListSiteOnenoteNotebooks(ctx, id, options)
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

	result = ListSiteOnenoteNotebooksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
