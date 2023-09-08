package meowneddevice

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

type ListMeOwnedDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListMeOwnedDevicesCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListMeOwnedDevices ...
func (c MeOwnedDeviceClient) ListMeOwnedDevices(ctx context.Context) (result ListMeOwnedDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/ownedDevices",
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

// ListMeOwnedDevicesComplete retrieves all the results into a single object
func (c MeOwnedDeviceClient) ListMeOwnedDevicesComplete(ctx context.Context) (ListMeOwnedDevicesCompleteResult, error) {
	return c.ListMeOwnedDevicesCompleteMatchingPredicate(ctx, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListMeOwnedDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOwnedDeviceClient) ListMeOwnedDevicesCompleteMatchingPredicate(ctx context.Context, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListMeOwnedDevicesCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListMeOwnedDevices(ctx)
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

	result = ListMeOwnedDevicesCompleteResult{
		Items: items,
	}
	return
}
