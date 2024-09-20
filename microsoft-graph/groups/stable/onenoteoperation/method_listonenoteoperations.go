package onenoteoperation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOnenoteOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OnenoteOperation
}

type ListOnenoteOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OnenoteOperation
}

type ListOnenoteOperationsOperationOptions struct {
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

func DefaultListOnenoteOperationsOperationOptions() ListOnenoteOperationsOperationOptions {
	return ListOnenoteOperationsOperationOptions{}
}

func (o ListOnenoteOperationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnenoteOperationsOperationOptions) ToOData() *odata.Query {
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

func (o ListOnenoteOperationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnenoteOperationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnenoteOperationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnenoteOperations - Get operations from groups. The status of OneNote operations. Getting an operations
// collection isn't supported, but you can get the status of long-running operations if the Operation-Location header is
// returned in the response. Read-only. Nullable.
func (c OnenoteOperationClient) ListOnenoteOperations(ctx context.Context, id stable.GroupId, options ListOnenoteOperationsOperationOptions) (result ListOnenoteOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnenoteOperationsCustomPager{},
		Path:          fmt.Sprintf("%s/onenote/operations", id.ID()),
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
		Values *[]stable.OnenoteOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnenoteOperationsComplete retrieves all the results into a single object
func (c OnenoteOperationClient) ListOnenoteOperationsComplete(ctx context.Context, id stable.GroupId, options ListOnenoteOperationsOperationOptions) (ListOnenoteOperationsCompleteResult, error) {
	return c.ListOnenoteOperationsCompleteMatchingPredicate(ctx, id, options, OnenoteOperationOperationPredicate{})
}

// ListOnenoteOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnenoteOperationClient) ListOnenoteOperationsCompleteMatchingPredicate(ctx context.Context, id stable.GroupId, options ListOnenoteOperationsOperationOptions, predicate OnenoteOperationOperationPredicate) (result ListOnenoteOperationsCompleteResult, err error) {
	items := make([]stable.OnenoteOperation, 0)

	resp, err := c.ListOnenoteOperations(ctx, id, options)
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

	result = ListOnenoteOperationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
