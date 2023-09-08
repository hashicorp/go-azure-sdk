package directorysubscription

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

type ListDirectorySubscriptionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CompanySubscriptionCollectionResponse
}

type ListDirectorySubscriptionsCompleteResult struct {
	Items []models.CompanySubscriptionCollectionResponse
}

// ListDirectorySubscriptions ...
func (c DirectorySubscriptionClient) ListDirectorySubscriptions(ctx context.Context) (result ListDirectorySubscriptionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/directory/subscriptions",
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
		Values *[]models.CompanySubscriptionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectorySubscriptionsComplete retrieves all the results into a single object
func (c DirectorySubscriptionClient) ListDirectorySubscriptionsComplete(ctx context.Context) (ListDirectorySubscriptionsCompleteResult, error) {
	return c.ListDirectorySubscriptionsCompleteMatchingPredicate(ctx, models.CompanySubscriptionCollectionResponseOperationPredicate{})
}

// ListDirectorySubscriptionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectorySubscriptionClient) ListDirectorySubscriptionsCompleteMatchingPredicate(ctx context.Context, predicate models.CompanySubscriptionCollectionResponseOperationPredicate) (result ListDirectorySubscriptionsCompleteResult, err error) {
	items := make([]models.CompanySubscriptionCollectionResponse, 0)

	resp, err := c.ListDirectorySubscriptions(ctx)
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

	result = ListDirectorySubscriptionsCompleteResult{
		Items: items,
	}
	return
}
