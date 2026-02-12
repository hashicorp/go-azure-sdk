package notebooks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookGetNotebookSummaryByWorkSpaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]NotebookResource
}

type NotebookGetNotebookSummaryByWorkSpaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []NotebookResource
}

type NotebookGetNotebookSummaryByWorkSpaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *NotebookGetNotebookSummaryByWorkSpaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// NotebookGetNotebookSummaryByWorkSpace ...
func (c NotebooksClient) NotebookGetNotebookSummaryByWorkSpace(ctx context.Context) (result NotebookGetNotebookSummaryByWorkSpaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &NotebookGetNotebookSummaryByWorkSpaceCustomPager{},
		Path:       "/notebooks/summary",
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
		Values *[]NotebookResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// NotebookGetNotebookSummaryByWorkSpaceComplete retrieves all the results into a single object
func (c NotebooksClient) NotebookGetNotebookSummaryByWorkSpaceComplete(ctx context.Context) (NotebookGetNotebookSummaryByWorkSpaceCompleteResult, error) {
	return c.NotebookGetNotebookSummaryByWorkSpaceCompleteMatchingPredicate(ctx, NotebookResourceOperationPredicate{})
}

// NotebookGetNotebookSummaryByWorkSpaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NotebooksClient) NotebookGetNotebookSummaryByWorkSpaceCompleteMatchingPredicate(ctx context.Context, predicate NotebookResourceOperationPredicate) (result NotebookGetNotebookSummaryByWorkSpaceCompleteResult, err error) {
	items := make([]NotebookResource, 0)

	resp, err := c.NotebookGetNotebookSummaryByWorkSpace(ctx)
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

	result = NotebookGetNotebookSummaryByWorkSpaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
