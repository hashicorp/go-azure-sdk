package userusageright

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

type ListUserByIdUsageRightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UsageRightCollectionResponse
}

type ListUserByIdUsageRightsCompleteResult struct {
	Items []models.UsageRightCollectionResponse
}

// ListUserByIdUsageRights ...
func (c UserUsageRightClient) ListUserByIdUsageRights(ctx context.Context, id UserId) (result ListUserByIdUsageRightsOperationResponse, err error) {
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

// ListUserByIdUsageRightsComplete retrieves all the results into a single object
func (c UserUsageRightClient) ListUserByIdUsageRightsComplete(ctx context.Context, id UserId) (ListUserByIdUsageRightsCompleteResult, error) {
	return c.ListUserByIdUsageRightsCompleteMatchingPredicate(ctx, id, models.UsageRightCollectionResponseOperationPredicate{})
}

// ListUserByIdUsageRightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserUsageRightClient) ListUserByIdUsageRightsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.UsageRightCollectionResponseOperationPredicate) (result ListUserByIdUsageRightsCompleteResult, err error) {
	items := make([]models.UsageRightCollectionResponse, 0)

	resp, err := c.ListUserByIdUsageRights(ctx, id)
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

	result = ListUserByIdUsageRightsCompleteResult{
		Items: items,
	}
	return
}
