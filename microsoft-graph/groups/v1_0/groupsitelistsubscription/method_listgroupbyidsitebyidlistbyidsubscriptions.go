package groupsitelistsubscription

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

type ListGroupByIdSiteByIdListByIdSubscriptionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SubscriptionCollectionResponse
}

type ListGroupByIdSiteByIdListByIdSubscriptionsCompleteResult struct {
	Items []models.SubscriptionCollectionResponse
}

// ListGroupByIdSiteByIdListByIdSubscriptions ...
func (c GroupSiteListSubscriptionClient) ListGroupByIdSiteByIdListByIdSubscriptions(ctx context.Context, id GroupSiteListId) (result ListGroupByIdSiteByIdListByIdSubscriptionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/subscriptions", id.ID()),
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
		Values *[]models.SubscriptionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdListByIdSubscriptionsComplete retrieves all the results into a single object
func (c GroupSiteListSubscriptionClient) ListGroupByIdSiteByIdListByIdSubscriptionsComplete(ctx context.Context, id GroupSiteListId) (ListGroupByIdSiteByIdListByIdSubscriptionsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdListByIdSubscriptionsCompleteMatchingPredicate(ctx, id, models.SubscriptionCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdListByIdSubscriptionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteListSubscriptionClient) ListGroupByIdSiteByIdListByIdSubscriptionsCompleteMatchingPredicate(ctx context.Context, id GroupSiteListId, predicate models.SubscriptionCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdListByIdSubscriptionsCompleteResult, err error) {
	items := make([]models.SubscriptionCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdListByIdSubscriptions(ctx, id)
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

	result = ListGroupByIdSiteByIdListByIdSubscriptionsCompleteResult{
		Items: items,
	}
	return
}
