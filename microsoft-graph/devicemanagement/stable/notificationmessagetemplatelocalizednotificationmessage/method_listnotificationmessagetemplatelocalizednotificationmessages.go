package notificationmessagetemplatelocalizednotificationmessage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListNotificationMessageTemplateLocalizedNotificationMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.LocalizedNotificationMessage
}

type ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.LocalizedNotificationMessage
}

type ListNotificationMessageTemplateLocalizedNotificationMessagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListNotificationMessageTemplateLocalizedNotificationMessagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListNotificationMessageTemplateLocalizedNotificationMessages ...
func (c NotificationMessageTemplateLocalizedNotificationMessageClient) ListNotificationMessageTemplateLocalizedNotificationMessages(ctx context.Context, id DeviceManagementNotificationMessageTemplateId) (result ListNotificationMessageTemplateLocalizedNotificationMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListNotificationMessageTemplateLocalizedNotificationMessagesCustomPager{},
		Path:       fmt.Sprintf("%s/localizedNotificationMessages", id.ID()),
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
		Values *[]stable.LocalizedNotificationMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListNotificationMessageTemplateLocalizedNotificationMessagesComplete retrieves all the results into a single object
func (c NotificationMessageTemplateLocalizedNotificationMessageClient) ListNotificationMessageTemplateLocalizedNotificationMessagesComplete(ctx context.Context, id DeviceManagementNotificationMessageTemplateId) (ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteResult, error) {
	return c.ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteMatchingPredicate(ctx, id, LocalizedNotificationMessageOperationPredicate{})
}

// ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NotificationMessageTemplateLocalizedNotificationMessageClient) ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementNotificationMessageTemplateId, predicate LocalizedNotificationMessageOperationPredicate) (result ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteResult, err error) {
	items := make([]stable.LocalizedNotificationMessage, 0)

	resp, err := c.ListNotificationMessageTemplateLocalizedNotificationMessages(ctx, id)
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

	result = ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
