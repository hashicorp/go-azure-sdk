package meonenotenotebook

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeOnenoteNotebooksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.NotebookCollectionResponse
}

type ListMeOnenoteNotebooksCompleteResult struct {
	Items []models.NotebookCollectionResponse
}

// ListMeOnenoteNotebooks ...
func (c MeOnenoteNotebookClient) ListMeOnenoteNotebooks(ctx context.Context) (result ListMeOnenoteNotebooksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/onenote/notebooks",
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
		Values *[]models.NotebookCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeOnenoteNotebooksComplete retrieves all the results into a single object
func (c MeOnenoteNotebookClient) ListMeOnenoteNotebooksComplete(ctx context.Context) (ListMeOnenoteNotebooksCompleteResult, error) {
	return c.ListMeOnenoteNotebooksCompleteMatchingPredicate(ctx, models.NotebookCollectionResponseOperationPredicate{})
}

// ListMeOnenoteNotebooksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnenoteNotebookClient) ListMeOnenoteNotebooksCompleteMatchingPredicate(ctx context.Context, predicate models.NotebookCollectionResponseOperationPredicate) (result ListMeOnenoteNotebooksCompleteResult, err error) {
	items := make([]models.NotebookCollectionResponse, 0)

	resp, err := c.ListMeOnenoteNotebooks(ctx)
	if err != nil {
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

	result = ListMeOnenoteNotebooksCompleteResult{
		Items: items,
	}
	return
}
