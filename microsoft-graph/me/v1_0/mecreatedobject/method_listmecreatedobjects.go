package mecreatedobject

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

type ListMeCreatedObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListMeCreatedObjectsCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListMeCreatedObjects ...
func (c MeCreatedObjectClient) ListMeCreatedObjects(ctx context.Context) (result ListMeCreatedObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/createdObjects",
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

// ListMeCreatedObjectsComplete retrieves all the results into a single object
func (c MeCreatedObjectClient) ListMeCreatedObjectsComplete(ctx context.Context) (ListMeCreatedObjectsCompleteResult, error) {
	return c.ListMeCreatedObjectsCompleteMatchingPredicate(ctx, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListMeCreatedObjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCreatedObjectClient) ListMeCreatedObjectsCompleteMatchingPredicate(ctx context.Context, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListMeCreatedObjectsCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListMeCreatedObjects(ctx)
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

	result = ListMeCreatedObjectsCompleteResult{
		Items: items,
	}
	return
}
