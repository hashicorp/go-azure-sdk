package userdeviceusageright

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

type ListUserByIdDeviceByIdUsageRightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UsageRightCollectionResponse
}

type ListUserByIdDeviceByIdUsageRightsCompleteResult struct {
	Items []models.UsageRightCollectionResponse
}

// ListUserByIdDeviceByIdUsageRights ...
func (c UserDeviceUsageRightClient) ListUserByIdDeviceByIdUsageRights(ctx context.Context, id UserDeviceId) (result ListUserByIdDeviceByIdUsageRightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/usageRights", id.ID()),
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
		Values *[]models.UsageRightCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdDeviceByIdUsageRightsComplete retrieves all the results into a single object
func (c UserDeviceUsageRightClient) ListUserByIdDeviceByIdUsageRightsComplete(ctx context.Context, id UserDeviceId) (ListUserByIdDeviceByIdUsageRightsCompleteResult, error) {
	return c.ListUserByIdDeviceByIdUsageRightsCompleteMatchingPredicate(ctx, id, models.UsageRightCollectionResponseOperationPredicate{})
}

// ListUserByIdDeviceByIdUsageRightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserDeviceUsageRightClient) ListUserByIdDeviceByIdUsageRightsCompleteMatchingPredicate(ctx context.Context, id UserDeviceId, predicate models.UsageRightCollectionResponseOperationPredicate) (result ListUserByIdDeviceByIdUsageRightsCompleteResult, err error) {
	items := make([]models.UsageRightCollectionResponse, 0)

	resp, err := c.ListUserByIdDeviceByIdUsageRights(ctx, id)
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

	result = ListUserByIdDeviceByIdUsageRightsCompleteResult{
		Items: items,
	}
	return
}
