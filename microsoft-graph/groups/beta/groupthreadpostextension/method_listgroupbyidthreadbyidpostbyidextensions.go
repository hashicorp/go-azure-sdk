package groupthreadpostextension

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

type ListGroupByIdThreadByIdPostByIdExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ExtensionCollectionResponse
}

type ListGroupByIdThreadByIdPostByIdExtensionsCompleteResult struct {
	Items []models.ExtensionCollectionResponse
}

// ListGroupByIdThreadByIdPostByIdExtensions ...
func (c GroupThreadPostExtensionClient) ListGroupByIdThreadByIdPostByIdExtensions(ctx context.Context, id GroupThreadPostId) (result ListGroupByIdThreadByIdPostByIdExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/extensions", id.ID()),
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
		Values *[]models.ExtensionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdThreadByIdPostByIdExtensionsComplete retrieves all the results into a single object
func (c GroupThreadPostExtensionClient) ListGroupByIdThreadByIdPostByIdExtensionsComplete(ctx context.Context, id GroupThreadPostId) (ListGroupByIdThreadByIdPostByIdExtensionsCompleteResult, error) {
	return c.ListGroupByIdThreadByIdPostByIdExtensionsCompleteMatchingPredicate(ctx, id, models.ExtensionCollectionResponseOperationPredicate{})
}

// ListGroupByIdThreadByIdPostByIdExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupThreadPostExtensionClient) ListGroupByIdThreadByIdPostByIdExtensionsCompleteMatchingPredicate(ctx context.Context, id GroupThreadPostId, predicate models.ExtensionCollectionResponseOperationPredicate) (result ListGroupByIdThreadByIdPostByIdExtensionsCompleteResult, err error) {
	items := make([]models.ExtensionCollectionResponse, 0)

	resp, err := c.ListGroupByIdThreadByIdPostByIdExtensions(ctx, id)
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

	result = ListGroupByIdThreadByIdPostByIdExtensionsCompleteResult{
		Items: items,
	}
	return
}
