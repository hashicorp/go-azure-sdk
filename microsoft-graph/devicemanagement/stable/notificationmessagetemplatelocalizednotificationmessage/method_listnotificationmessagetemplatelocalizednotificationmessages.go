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

type ListNotificationMessageTemplateLocalizedNotificationMessagesOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListNotificationMessageTemplateLocalizedNotificationMessagesOperationOptions() ListNotificationMessageTemplateLocalizedNotificationMessagesOperationOptions {
	return ListNotificationMessageTemplateLocalizedNotificationMessagesOperationOptions{}
}

func (o ListNotificationMessageTemplateLocalizedNotificationMessagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListNotificationMessageTemplateLocalizedNotificationMessagesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListNotificationMessageTemplateLocalizedNotificationMessagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListNotificationMessageTemplateLocalizedNotificationMessages - List localizedNotificationMessages. List properties
// and relationships of the localizedNotificationMessage objects.
func (c NotificationMessageTemplateLocalizedNotificationMessageClient) ListNotificationMessageTemplateLocalizedNotificationMessages(ctx context.Context, id stable.DeviceManagementNotificationMessageTemplateId, options ListNotificationMessageTemplateLocalizedNotificationMessagesOperationOptions) (result ListNotificationMessageTemplateLocalizedNotificationMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListNotificationMessageTemplateLocalizedNotificationMessagesCustomPager{},
		Path:          fmt.Sprintf("%s/localizedNotificationMessages", id.ID()),
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
func (c NotificationMessageTemplateLocalizedNotificationMessageClient) ListNotificationMessageTemplateLocalizedNotificationMessagesComplete(ctx context.Context, id stable.DeviceManagementNotificationMessageTemplateId, options ListNotificationMessageTemplateLocalizedNotificationMessagesOperationOptions) (ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteResult, error) {
	return c.ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteMatchingPredicate(ctx, id, options, LocalizedNotificationMessageOperationPredicate{})
}

// ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NotificationMessageTemplateLocalizedNotificationMessageClient) ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementNotificationMessageTemplateId, options ListNotificationMessageTemplateLocalizedNotificationMessagesOperationOptions, predicate LocalizedNotificationMessageOperationPredicate) (result ListNotificationMessageTemplateLocalizedNotificationMessagesCompleteResult, err error) {
	items := make([]stable.LocalizedNotificationMessage, 0)

	resp, err := c.ListNotificationMessageTemplateLocalizedNotificationMessages(ctx, id, options)
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
