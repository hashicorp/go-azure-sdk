package calendarcalendarviewexceptionoccurrenceinstanceextension

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

type ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Extension
}

type ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Extension
}

type ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarCalendarViewExceptionOccurrenceInstanceExtensions ...
func (c CalendarCalendarViewExceptionOccurrenceInstanceExtensionClient) ListCalendarCalendarViewExceptionOccurrenceInstanceExtensions(ctx context.Context, id MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceId) (result ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsCustomPager{},
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
		Values *[]beta.Extension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsComplete retrieves all the results into a single object
func (c CalendarCalendarViewExceptionOccurrenceInstanceExtensionClient) ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsComplete(ctx context.Context, id MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceId) (ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsCompleteResult, error) {
	return c.ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsCompleteMatchingPredicate(ctx, id, ExtensionOperationPredicate{})
}

// ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarCalendarViewExceptionOccurrenceInstanceExtensionClient) ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsCompleteMatchingPredicate(ctx context.Context, id MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceId, predicate ExtensionOperationPredicate) (result ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsCompleteResult, err error) {
	items := make([]beta.Extension, 0)

	resp, err := c.ListCalendarCalendarViewExceptionOccurrenceInstanceExtensions(ctx, id)
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

	result = ListCalendarCalendarViewExceptionOccurrenceInstanceExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
