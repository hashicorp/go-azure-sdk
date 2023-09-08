package grouponenoteresource

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

type ListGroupByIdOnenoteResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenoteResourceCollectionResponse
}

type ListGroupByIdOnenoteResourcesCompleteResult struct {
	Items []models.OnenoteResourceCollectionResponse
}

// ListGroupByIdOnenoteResources ...
func (c GroupOnenoteResourceClient) ListGroupByIdOnenoteResources(ctx context.Context, id GroupId) (result ListGroupByIdOnenoteResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/onenote/resources", id.ID()),
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
		Values *[]models.OnenoteResourceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdOnenoteResourcesComplete retrieves all the results into a single object
func (c GroupOnenoteResourceClient) ListGroupByIdOnenoteResourcesComplete(ctx context.Context, id GroupId) (ListGroupByIdOnenoteResourcesCompleteResult, error) {
	return c.ListGroupByIdOnenoteResourcesCompleteMatchingPredicate(ctx, id, models.OnenoteResourceCollectionResponseOperationPredicate{})
}

// ListGroupByIdOnenoteResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupOnenoteResourceClient) ListGroupByIdOnenoteResourcesCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.OnenoteResourceCollectionResponseOperationPredicate) (result ListGroupByIdOnenoteResourcesCompleteResult, err error) {
	items := make([]models.OnenoteResourceCollectionResponse, 0)

	resp, err := c.ListGroupByIdOnenoteResources(ctx, id)
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

	result = ListGroupByIdOnenoteResourcesCompleteResult{
		Items: items,
	}
	return
}
