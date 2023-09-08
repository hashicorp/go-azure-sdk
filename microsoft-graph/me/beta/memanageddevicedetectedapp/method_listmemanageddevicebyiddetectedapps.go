package memanageddevicedetectedapp

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

type ListMeManagedDeviceByIdDetectedAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DetectedAppCollectionResponse
}

type ListMeManagedDeviceByIdDetectedAppsCompleteResult struct {
	Items []models.DetectedAppCollectionResponse
}

// ListMeManagedDeviceByIdDetectedApps ...
func (c MeManagedDeviceDetectedAppClient) ListMeManagedDeviceByIdDetectedApps(ctx context.Context, id MeManagedDeviceId) (result ListMeManagedDeviceByIdDetectedAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/detectedApps", id.ID()),
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
		Values *[]models.DetectedAppCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeManagedDeviceByIdDetectedAppsComplete retrieves all the results into a single object
func (c MeManagedDeviceDetectedAppClient) ListMeManagedDeviceByIdDetectedAppsComplete(ctx context.Context, id MeManagedDeviceId) (ListMeManagedDeviceByIdDetectedAppsCompleteResult, error) {
	return c.ListMeManagedDeviceByIdDetectedAppsCompleteMatchingPredicate(ctx, id, models.DetectedAppCollectionResponseOperationPredicate{})
}

// ListMeManagedDeviceByIdDetectedAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeManagedDeviceDetectedAppClient) ListMeManagedDeviceByIdDetectedAppsCompleteMatchingPredicate(ctx context.Context, id MeManagedDeviceId, predicate models.DetectedAppCollectionResponseOperationPredicate) (result ListMeManagedDeviceByIdDetectedAppsCompleteResult, err error) {
	items := make([]models.DetectedAppCollectionResponse, 0)

	resp, err := c.ListMeManagedDeviceByIdDetectedApps(ctx, id)
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

	result = ListMeManagedDeviceByIdDetectedAppsCompleteResult{
		Items: items,
	}
	return
}
