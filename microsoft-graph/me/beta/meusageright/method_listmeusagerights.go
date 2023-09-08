package meusageright

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

type ListMeUsageRightsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.UsageRightCollectionResponse
}

type ListMeUsageRightsCompleteResult struct {
	Items []models.UsageRightCollectionResponse
}

// ListMeUsageRights ...
func (c MeUsageRightClient) ListMeUsageRights(ctx context.Context) (result ListMeUsageRightsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/usageRights",
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

// ListMeUsageRightsComplete retrieves all the results into a single object
func (c MeUsageRightClient) ListMeUsageRightsComplete(ctx context.Context) (ListMeUsageRightsCompleteResult, error) {
	return c.ListMeUsageRightsCompleteMatchingPredicate(ctx, models.UsageRightCollectionResponseOperationPredicate{})
}

// ListMeUsageRightsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeUsageRightClient) ListMeUsageRightsCompleteMatchingPredicate(ctx context.Context, predicate models.UsageRightCollectionResponseOperationPredicate) (result ListMeUsageRightsCompleteResult, err error) {
	items := make([]models.UsageRightCollectionResponse, 0)

	resp, err := c.ListMeUsageRights(ctx)
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

	result = ListMeUsageRightsCompleteResult{
		Items: items,
	}
	return
}
