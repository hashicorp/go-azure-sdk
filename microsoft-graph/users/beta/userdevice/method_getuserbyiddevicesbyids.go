package userdevice

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

type GetUserByIdDevicesByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObject
}

type GetUserByIdDevicesByIdsCompleteResult struct {
	Items []models.DirectoryObject
}

// GetUserByIdDevicesByIds ...
func (c UserDeviceClient) GetUserByIdDevicesByIds(ctx context.Context, id UserId) (result GetUserByIdDevicesByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/devices/getByIds", id.ID()),
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
		Values *[]models.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetUserByIdDevicesByIdsComplete retrieves all the results into a single object
func (c UserDeviceClient) GetUserByIdDevicesByIdsComplete(ctx context.Context, id UserId) (GetUserByIdDevicesByIdsCompleteResult, error) {
	return c.GetUserByIdDevicesByIdsCompleteMatchingPredicate(ctx, id, models.DirectoryObjectOperationPredicate{})
}

// GetUserByIdDevicesByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserDeviceClient) GetUserByIdDevicesByIdsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.DirectoryObjectOperationPredicate) (result GetUserByIdDevicesByIdsCompleteResult, err error) {
	items := make([]models.DirectoryObject, 0)

	resp, err := c.GetUserByIdDevicesByIds(ctx, id)
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

	result = GetUserByIdDevicesByIdsCompleteResult{
		Items: items,
	}
	return
}
