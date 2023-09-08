package usermanageddevice

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

type ListUserByIdManagedDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ManagedDeviceCollectionResponse
}

type ListUserByIdManagedDevicesCompleteResult struct {
	Items []models.ManagedDeviceCollectionResponse
}

// ListUserByIdManagedDevices ...
func (c UserManagedDeviceClient) ListUserByIdManagedDevices(ctx context.Context, id UserId) (result ListUserByIdManagedDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/managedDevices", id.ID()),
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
		Values *[]models.ManagedDeviceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdManagedDevicesComplete retrieves all the results into a single object
func (c UserManagedDeviceClient) ListUserByIdManagedDevicesComplete(ctx context.Context, id UserId) (ListUserByIdManagedDevicesCompleteResult, error) {
	return c.ListUserByIdManagedDevicesCompleteMatchingPredicate(ctx, id, models.ManagedDeviceCollectionResponseOperationPredicate{})
}

// ListUserByIdManagedDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserManagedDeviceClient) ListUserByIdManagedDevicesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ManagedDeviceCollectionResponseOperationPredicate) (result ListUserByIdManagedDevicesCompleteResult, err error) {
	items := make([]models.ManagedDeviceCollectionResponse, 0)

	resp, err := c.ListUserByIdManagedDevices(ctx, id)
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

	result = ListUserByIdManagedDevicesCompleteResult{
		Items: items,
	}
	return
}
