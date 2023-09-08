package serviceprincipalcreatedobject

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

type ListServicePrincipalByIdCreatedObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListServicePrincipalByIdCreatedObjectsCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListServicePrincipalByIdCreatedObjects ...
func (c ServicePrincipalCreatedObjectClient) ListServicePrincipalByIdCreatedObjects(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdCreatedObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/createdObjects", id.ID()),
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

// ListServicePrincipalByIdCreatedObjectsComplete retrieves all the results into a single object
func (c ServicePrincipalCreatedObjectClient) ListServicePrincipalByIdCreatedObjectsComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdCreatedObjectsCompleteResult, error) {
	return c.ListServicePrincipalByIdCreatedObjectsCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdCreatedObjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalCreatedObjectClient) ListServicePrincipalByIdCreatedObjectsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListServicePrincipalByIdCreatedObjectsCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdCreatedObjects(ctx, id)
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

	result = ListServicePrincipalByIdCreatedObjectsCompleteResult{
		Items: items,
	}
	return
}
