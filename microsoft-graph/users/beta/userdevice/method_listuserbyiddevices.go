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

type ListUserByIdDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DeviceCollectionResponse
}

type ListUserByIdDevicesCompleteResult struct {
	Items []models.DeviceCollectionResponse
}

// ListUserByIdDevices ...
func (c UserDeviceClient) ListUserByIdDevices(ctx context.Context, id UserId) (result ListUserByIdDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/devices", id.ID()),
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
		Values *[]models.DeviceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdDevicesComplete retrieves all the results into a single object
func (c UserDeviceClient) ListUserByIdDevicesComplete(ctx context.Context, id UserId) (ListUserByIdDevicesCompleteResult, error) {
	return c.ListUserByIdDevicesCompleteMatchingPredicate(ctx, id, models.DeviceCollectionResponseOperationPredicate{})
}

// ListUserByIdDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserDeviceClient) ListUserByIdDevicesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.DeviceCollectionResponseOperationPredicate) (result ListUserByIdDevicesCompleteResult, err error) {
	items := make([]models.DeviceCollectionResponse, 0)

	resp, err := c.ListUserByIdDevices(ctx, id)
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

	result = ListUserByIdDevicesCompleteResult{
		Items: items,
	}
	return
}
