package notificationmessagetemplate

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

type ListNotificationMessageTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.NotificationMessageTemplate
}

type ListNotificationMessageTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.NotificationMessageTemplate
}

type ListNotificationMessageTemplatesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListNotificationMessageTemplatesOperationOptions() ListNotificationMessageTemplatesOperationOptions {
	return ListNotificationMessageTemplatesOperationOptions{}
}

func (o ListNotificationMessageTemplatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListNotificationMessageTemplatesOperationOptions) ToOData() *odata.Query {
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

func (o ListNotificationMessageTemplatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListNotificationMessageTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListNotificationMessageTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListNotificationMessageTemplates - List notificationMessageTemplates. List properties and relationships of the
// notificationMessageTemplate objects.
func (c NotificationMessageTemplateClient) ListNotificationMessageTemplates(ctx context.Context, options ListNotificationMessageTemplatesOperationOptions) (result ListNotificationMessageTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListNotificationMessageTemplatesCustomPager{},
		Path:          "/deviceManagement/notificationMessageTemplates",
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
		Values *[]stable.NotificationMessageTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListNotificationMessageTemplatesComplete retrieves all the results into a single object
func (c NotificationMessageTemplateClient) ListNotificationMessageTemplatesComplete(ctx context.Context, options ListNotificationMessageTemplatesOperationOptions) (ListNotificationMessageTemplatesCompleteResult, error) {
	return c.ListNotificationMessageTemplatesCompleteMatchingPredicate(ctx, options, NotificationMessageTemplateOperationPredicate{})
}

// ListNotificationMessageTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NotificationMessageTemplateClient) ListNotificationMessageTemplatesCompleteMatchingPredicate(ctx context.Context, options ListNotificationMessageTemplatesOperationOptions, predicate NotificationMessageTemplateOperationPredicate) (result ListNotificationMessageTemplatesCompleteResult, err error) {
	items := make([]stable.NotificationMessageTemplate, 0)

	resp, err := c.ListNotificationMessageTemplates(ctx, options)
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

	result = ListNotificationMessageTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
