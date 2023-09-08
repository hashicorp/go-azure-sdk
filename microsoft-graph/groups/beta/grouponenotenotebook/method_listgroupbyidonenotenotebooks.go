package grouponenotenotebook

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdOnenoteNotebooksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.NotebookCollectionResponse
}

type ListGroupByIdOnenoteNotebooksCompleteResult struct {
	Items []models.NotebookCollectionResponse
}

// ListGroupByIdOnenoteNotebooks ...
func (c GroupOnenoteNotebookClient) ListGroupByIdOnenoteNotebooks(ctx context.Context, id GroupId) (result ListGroupByIdOnenoteNotebooksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/onenote/notebooks", id.ID()),
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

// ListGroupByIdOnenoteNotebooksComplete retrieves all the results into a single object
func (c GroupOnenoteNotebookClient) ListGroupByIdOnenoteNotebooksComplete(ctx context.Context, id GroupId) (ListGroupByIdOnenoteNotebooksCompleteResult, error) {
	return c.ListGroupByIdOnenoteNotebooksCompleteMatchingPredicate(ctx, id, models.NotebookCollectionResponseOperationPredicate{})
}

// ListGroupByIdOnenoteNotebooksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupOnenoteNotebookClient) ListGroupByIdOnenoteNotebooksCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.NotebookCollectionResponseOperationPredicate) (result ListGroupByIdOnenoteNotebooksCompleteResult, err error) {
	items := make([]models.NotebookCollectionResponse, 0)

	resp, err := c.ListGroupByIdOnenoteNotebooks(ctx, id)
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

	result = ListGroupByIdOnenoteNotebooksCompleteResult{
		Items: items,
	}
	return
}
