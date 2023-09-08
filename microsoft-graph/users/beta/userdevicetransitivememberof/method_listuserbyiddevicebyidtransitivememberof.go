package userdevicetransitivememberof

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

type ListUserByIdDeviceByIdTransitiveMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListUserByIdDeviceByIdTransitiveMemberOfCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListUserByIdDeviceByIdTransitiveMemberOf ...
func (c UserDeviceTransitiveMemberOfClient) ListUserByIdDeviceByIdTransitiveMemberOf(ctx context.Context, id UserDeviceId) (result ListUserByIdDeviceByIdTransitiveMemberOfOperationResponse, err error) {
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

// ListUserByIdDeviceByIdTransitiveMemberOfComplete retrieves all the results into a single object
func (c UserDeviceTransitiveMemberOfClient) ListUserByIdDeviceByIdTransitiveMemberOfComplete(ctx context.Context, id UserDeviceId) (ListUserByIdDeviceByIdTransitiveMemberOfCompleteResult, error) {
	return c.ListUserByIdDeviceByIdTransitiveMemberOfCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListUserByIdDeviceByIdTransitiveMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserDeviceTransitiveMemberOfClient) ListUserByIdDeviceByIdTransitiveMemberOfCompleteMatchingPredicate(ctx context.Context, id UserDeviceId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListUserByIdDeviceByIdTransitiveMemberOfCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListUserByIdDeviceByIdTransitiveMemberOf(ctx, id)
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

	result = ListUserByIdDeviceByIdTransitiveMemberOfCompleteResult{
		Items: items,
	}
	return
}
