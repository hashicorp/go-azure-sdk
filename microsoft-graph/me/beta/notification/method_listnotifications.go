package notification

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListNotificationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Notification
}

type ListNotificationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Notification
}

type ListNotificationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListNotificationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListNotifications ...
func (c NotificationClient) ListNotifications(ctx context.Context) (result ListNotificationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListNotificationsCustomPager{},
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
		Values *[]beta.Notification `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListNotificationsComplete retrieves all the results into a single object
func (c NotificationClient) ListNotificationsComplete(ctx context.Context) (ListNotificationsCompleteResult, error) {
	return c.ListNotificationsCompleteMatchingPredicate(ctx, NotificationOperationPredicate{})
}

// ListNotificationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NotificationClient) ListNotificationsCompleteMatchingPredicate(ctx context.Context, predicate NotificationOperationPredicate) (result ListNotificationsCompleteResult, err error) {
	items := make([]beta.Notification, 0)

	resp, err := c.ListNotifications(ctx)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListNotificationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
