package memanageddevicelogcollectionrequest

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

type ListMeManagedDeviceByIdLogCollectionRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DeviceLogCollectionResponseCollectionResponse
}

type ListMeManagedDeviceByIdLogCollectionRequestsCompleteResult struct {
	Items []models.DeviceLogCollectionResponseCollectionResponse
}

// ListMeManagedDeviceByIdLogCollectionRequests ...
func (c MeManagedDeviceLogCollectionRequestClient) ListMeManagedDeviceByIdLogCollectionRequests(ctx context.Context, id MeManagedDeviceId) (result ListMeManagedDeviceByIdLogCollectionRequestsOperationResponse, err error) {
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

// ListMeManagedDeviceByIdLogCollectionRequestsComplete retrieves all the results into a single object
func (c MeManagedDeviceLogCollectionRequestClient) ListMeManagedDeviceByIdLogCollectionRequestsComplete(ctx context.Context, id MeManagedDeviceId) (ListMeManagedDeviceByIdLogCollectionRequestsCompleteResult, error) {
	return c.ListMeManagedDeviceByIdLogCollectionRequestsCompleteMatchingPredicate(ctx, id, models.DeviceLogCollectionResponseCollectionResponseOperationPredicate{})
}

// ListMeManagedDeviceByIdLogCollectionRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeManagedDeviceLogCollectionRequestClient) ListMeManagedDeviceByIdLogCollectionRequestsCompleteMatchingPredicate(ctx context.Context, id MeManagedDeviceId, predicate models.DeviceLogCollectionResponseCollectionResponseOperationPredicate) (result ListMeManagedDeviceByIdLogCollectionRequestsCompleteResult, err error) {
	items := make([]models.DeviceLogCollectionResponseCollectionResponse, 0)

	resp, err := c.ListMeManagedDeviceByIdLogCollectionRequests(ctx, id)
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

	result = ListMeManagedDeviceByIdLogCollectionRequestsCompleteResult{
		Items: items,
	}
	return
}
