package grouptransitivememberof

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

type ListGroupByIdTransitiveMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListGroupByIdTransitiveMemberOfCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListGroupByIdTransitiveMemberOf ...
func (c GroupTransitiveMemberOfClient) ListGroupByIdTransitiveMemberOf(ctx context.Context, id GroupId) (result ListGroupByIdTransitiveMemberOfOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/transitiveMemberOf", id.ID()),
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
		Values *[]models.DirectoryObjectCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdTransitiveMemberOfComplete retrieves all the results into a single object
func (c GroupTransitiveMemberOfClient) ListGroupByIdTransitiveMemberOfComplete(ctx context.Context, id GroupId) (ListGroupByIdTransitiveMemberOfCompleteResult, error) {
	return c.ListGroupByIdTransitiveMemberOfCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListGroupByIdTransitiveMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupTransitiveMemberOfClient) ListGroupByIdTransitiveMemberOfCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListGroupByIdTransitiveMemberOfCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListGroupByIdTransitiveMemberOf(ctx, id)
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

	result = ListGroupByIdTransitiveMemberOfCompleteResult{
		Items: items,
	}
	return
}
