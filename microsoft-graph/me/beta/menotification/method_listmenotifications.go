package menotification

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

type ListMeNotificationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.NotificationCollectionResponse
}

type ListMeNotificationsCompleteResult struct {
	Items []models.NotificationCollectionResponse
}

// ListMeNotifications ...
func (c MeNotificationClient) ListMeNotifications(ctx context.Context) (result ListMeNotificationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/notifications",
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
		Values *[]models.NotificationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeNotificationsComplete retrieves all the results into a single object
func (c MeNotificationClient) ListMeNotificationsComplete(ctx context.Context) (ListMeNotificationsCompleteResult, error) {
	return c.ListMeNotificationsCompleteMatchingPredicate(ctx, models.NotificationCollectionResponseOperationPredicate{})
}

// ListMeNotificationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeNotificationClient) ListMeNotificationsCompleteMatchingPredicate(ctx context.Context, predicate models.NotificationCollectionResponseOperationPredicate) (result ListMeNotificationsCompleteResult, err error) {
	items := make([]models.NotificationCollectionResponse, 0)

	resp, err := c.ListMeNotifications(ctx)
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

	result = ListMeNotificationsCompleteResult{
		Items: items,
	}
	return
}
