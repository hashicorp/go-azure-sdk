package memanageddevice

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

type ListMeManagedDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ManagedDeviceCollectionResponse
}

type ListMeManagedDevicesCompleteResult struct {
	Items []models.ManagedDeviceCollectionResponse
}

// ListMeManagedDevices ...
func (c MeManagedDeviceClient) ListMeManagedDevices(ctx context.Context) (result ListMeManagedDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/managedDevices",
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

// ListMeManagedDevicesComplete retrieves all the results into a single object
func (c MeManagedDeviceClient) ListMeManagedDevicesComplete(ctx context.Context) (ListMeManagedDevicesCompleteResult, error) {
	return c.ListMeManagedDevicesCompleteMatchingPredicate(ctx, models.ManagedDeviceCollectionResponseOperationPredicate{})
}

// ListMeManagedDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeManagedDeviceClient) ListMeManagedDevicesCompleteMatchingPredicate(ctx context.Context, predicate models.ManagedDeviceCollectionResponseOperationPredicate) (result ListMeManagedDevicesCompleteResult, err error) {
	items := make([]models.ManagedDeviceCollectionResponse, 0)

	resp, err := c.ListMeManagedDevices(ctx)
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

	result = ListMeManagedDevicesCompleteResult{
		Items: items,
	}
	return
}
