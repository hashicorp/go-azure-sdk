package calendareventinstanceextension

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

type ListCalendarEventInstanceExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Extension
}

type ListCalendarEventInstanceExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Extension
}

type ListCalendarEventInstanceExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarEventInstanceExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarEventInstanceExtensions ...
func (c CalendarEventInstanceExtensionClient) ListCalendarEventInstanceExtensions(ctx context.Context, id GroupIdCalendarEventIdInstanceId) (result ListCalendarEventInstanceExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarEventInstanceExtensionsCustomPager{},
		Path:       fmt.Sprintf("%s/extensions", id.ID()),
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
		Values *[]stable.Extension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCalendarEventInstanceExtensionsComplete retrieves all the results into a single object
func (c CalendarEventInstanceExtensionClient) ListCalendarEventInstanceExtensionsComplete(ctx context.Context, id GroupIdCalendarEventIdInstanceId) (ListCalendarEventInstanceExtensionsCompleteResult, error) {
	return c.ListCalendarEventInstanceExtensionsCompleteMatchingPredicate(ctx, id, ExtensionOperationPredicate{})
}

// ListCalendarEventInstanceExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarEventInstanceExtensionClient) ListCalendarEventInstanceExtensionsCompleteMatchingPredicate(ctx context.Context, id GroupIdCalendarEventIdInstanceId, predicate ExtensionOperationPredicate) (result ListCalendarEventInstanceExtensionsCompleteResult, err error) {
	items := make([]stable.Extension, 0)

	resp, err := c.ListCalendarEventInstanceExtensions(ctx, id)
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

	result = ListCalendarEventInstanceExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
