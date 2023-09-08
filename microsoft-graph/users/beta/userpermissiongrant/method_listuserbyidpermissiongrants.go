package userpermissiongrant

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

type ListUserByIdPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ResourceSpecificPermissionGrantCollectionResponse
}

type ListUserByIdPermissionGrantsCompleteResult struct {
	Items []models.ResourceSpecificPermissionGrantCollectionResponse
}

// ListUserByIdPermissionGrants ...
func (c UserPermissionGrantClient) ListUserByIdPermissionGrants(ctx context.Context, id UserId) (result ListUserByIdPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/permissionGrants", id.ID()),
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
		Values *[]models.ResourceSpecificPermissionGrantCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdPermissionGrantsComplete retrieves all the results into a single object
func (c UserPermissionGrantClient) ListUserByIdPermissionGrantsComplete(ctx context.Context, id UserId) (ListUserByIdPermissionGrantsCompleteResult, error) {
	return c.ListUserByIdPermissionGrantsCompleteMatchingPredicate(ctx, id, models.ResourceSpecificPermissionGrantCollectionResponseOperationPredicate{})
}

// ListUserByIdPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPermissionGrantClient) ListUserByIdPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ResourceSpecificPermissionGrantCollectionResponseOperationPredicate) (result ListUserByIdPermissionGrantsCompleteResult, err error) {
	items := make([]models.ResourceSpecificPermissionGrantCollectionResponse, 0)

	resp, err := c.ListUserByIdPermissionGrants(ctx, id)
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

	result = ListUserByIdPermissionGrantsCompleteResult{
		Items: items,
	}
	return
}
