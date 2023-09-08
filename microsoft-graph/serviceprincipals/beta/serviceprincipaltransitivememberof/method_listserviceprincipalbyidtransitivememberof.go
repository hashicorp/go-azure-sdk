package serviceprincipaltransitivememberof

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

type ListServicePrincipalByIdTransitiveMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListServicePrincipalByIdTransitiveMemberOfCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListServicePrincipalByIdTransitiveMemberOf ...
func (c ServicePrincipalTransitiveMemberOfClient) ListServicePrincipalByIdTransitiveMemberOf(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdTransitiveMemberOfOperationResponse, err error) {
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

// ListServicePrincipalByIdTransitiveMemberOfComplete retrieves all the results into a single object
func (c ServicePrincipalTransitiveMemberOfClient) ListServicePrincipalByIdTransitiveMemberOfComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdTransitiveMemberOfCompleteResult, error) {
	return c.ListServicePrincipalByIdTransitiveMemberOfCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdTransitiveMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalTransitiveMemberOfClient) ListServicePrincipalByIdTransitiveMemberOfCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListServicePrincipalByIdTransitiveMemberOfCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdTransitiveMemberOf(ctx, id)
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

	result = ListServicePrincipalByIdTransitiveMemberOfCompleteResult{
		Items: items,
	}
	return
}
