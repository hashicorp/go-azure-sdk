package usernotification

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

type ListUserByIdNotificationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.NotificationCollectionResponse
}

type ListUserByIdNotificationsCompleteResult struct {
	Items []models.NotificationCollectionResponse
}

// ListUserByIdNotifications ...
func (c UserNotificationClient) ListUserByIdNotifications(ctx context.Context, id UserId) (result ListUserByIdNotificationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/notifications", id.ID()),
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

// ListUserByIdNotificationsComplete retrieves all the results into a single object
func (c UserNotificationClient) ListUserByIdNotificationsComplete(ctx context.Context, id UserId) (ListUserByIdNotificationsCompleteResult, error) {
	return c.ListUserByIdNotificationsCompleteMatchingPredicate(ctx, id, models.NotificationCollectionResponseOperationPredicate{})
}

// ListUserByIdNotificationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserNotificationClient) ListUserByIdNotificationsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.NotificationCollectionResponseOperationPredicate) (result ListUserByIdNotificationsCompleteResult, err error) {
	items := make([]models.NotificationCollectionResponse, 0)

	resp, err := c.ListUserByIdNotifications(ctx, id)
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

	result = ListUserByIdNotificationsCompleteResult{
		Items: items,
	}
	return
}
