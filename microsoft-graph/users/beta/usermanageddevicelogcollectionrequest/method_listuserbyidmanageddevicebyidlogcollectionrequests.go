package usermanageddevicelogcollectionrequest

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

type ListUserByIdManagedDeviceByIdLogCollectionRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DeviceLogCollectionResponseCollectionResponse
}

type ListUserByIdManagedDeviceByIdLogCollectionRequestsCompleteResult struct {
	Items []models.DeviceLogCollectionResponseCollectionResponse
}

// ListUserByIdManagedDeviceByIdLogCollectionRequests ...
func (c UserManagedDeviceLogCollectionRequestClient) ListUserByIdManagedDeviceByIdLogCollectionRequests(ctx context.Context, id UserManagedDeviceId) (result ListUserByIdManagedDeviceByIdLogCollectionRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/logCollectionRequests", id.ID()),
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
		Values *[]models.DeviceLogCollectionResponseCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdManagedDeviceByIdLogCollectionRequestsComplete retrieves all the results into a single object
func (c UserManagedDeviceLogCollectionRequestClient) ListUserByIdManagedDeviceByIdLogCollectionRequestsComplete(ctx context.Context, id UserManagedDeviceId) (ListUserByIdManagedDeviceByIdLogCollectionRequestsCompleteResult, error) {
	return c.ListUserByIdManagedDeviceByIdLogCollectionRequestsCompleteMatchingPredicate(ctx, id, models.DeviceLogCollectionResponseCollectionResponseOperationPredicate{})
}

// ListUserByIdManagedDeviceByIdLogCollectionRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserManagedDeviceLogCollectionRequestClient) ListUserByIdManagedDeviceByIdLogCollectionRequestsCompleteMatchingPredicate(ctx context.Context, id UserManagedDeviceId, predicate models.DeviceLogCollectionResponseCollectionResponseOperationPredicate) (result ListUserByIdManagedDeviceByIdLogCollectionRequestsCompleteResult, err error) {
	items := make([]models.DeviceLogCollectionResponseCollectionResponse, 0)

	resp, err := c.ListUserByIdManagedDeviceByIdLogCollectionRequests(ctx, id)
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

	result = ListUserByIdManagedDeviceByIdLogCollectionRequestsCompleteResult{
		Items: items,
	}
	return
}
