package onenoteoperation

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

type ListOnenoteOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OnenoteOperation
}

type ListOnenoteOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OnenoteOperation
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

// ListOnenoteOperations - Get onenoteOperation. Get the status of a long-running OneNote operation. The status applies
// to operations that return the Operation-Location header in the response, such as CopyNotebook, CopyToNotebook,
// CopyToSectionGroup, and CopyToSection. You can poll the Operation-Location endpoint until the status property returns
// completed or failed. If the status is completed, the resourceLocation property contains the resource endpoint URI. If
// the status is failed, the error and @api.diagnostics properties provide error information.
func (c OnenoteOperationClient) ListOnenoteOperations(ctx context.Context, options ListOnenoteOperationsOperationOptions) (result ListOnenoteOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnenoteOperationsCustomPager{},
		Path:          "/me/onenote/operations",
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
		Values *[]beta.OnenoteOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOnenoteOperationsComplete retrieves all the results into a single object
func (c OnenoteOperationClient) ListOnenoteOperationsComplete(ctx context.Context, options ListOnenoteOperationsOperationOptions) (ListOnenoteOperationsCompleteResult, error) {
	return c.ListOnenoteOperationsCompleteMatchingPredicate(ctx, options, OnenoteOperationOperationPredicate{})
}

// ListOnenoteOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnenoteOperationClient) ListOnenoteOperationsCompleteMatchingPredicate(ctx context.Context, options ListOnenoteOperationsOperationOptions, predicate OnenoteOperationOperationPredicate) (result ListOnenoteOperationsCompleteResult, err error) {
	items := make([]beta.OnenoteOperation, 0)

	resp, err := c.ListOnenoteOperations(ctx, options)
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
