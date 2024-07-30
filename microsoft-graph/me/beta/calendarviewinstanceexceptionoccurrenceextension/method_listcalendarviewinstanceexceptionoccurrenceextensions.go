package calendarviewinstanceexceptionoccurrenceextension

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

type ListCalendarViewInstanceExceptionOccurrenceExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Extension
}

type ListCalendarViewInstanceExceptionOccurrenceExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Extension
}

type ListCalendarViewInstanceExceptionOccurrenceExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarViewInstanceExceptionOccurrenceExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarViewInstanceExceptionOccurrenceExtensions ...
func (c CalendarViewInstanceExceptionOccurrenceExtensionClient) ListCalendarViewInstanceExceptionOccurrenceExtensions(ctx context.Context, id MeCalendarViewIdInstanceIdExceptionOccurrenceId) (result ListCalendarViewInstanceExceptionOccurrenceExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarViewInstanceExceptionOccurrenceExtensionsCustomPager{},
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

// ListCalendarViewInstanceExceptionOccurrenceExtensionsComplete retrieves all the results into a single object
func (c CalendarViewInstanceExceptionOccurrenceExtensionClient) ListCalendarViewInstanceExceptionOccurrenceExtensionsComplete(ctx context.Context, id MeCalendarViewIdInstanceIdExceptionOccurrenceId) (ListCalendarViewInstanceExceptionOccurrenceExtensionsCompleteResult, error) {
	return c.ListCalendarViewInstanceExceptionOccurrenceExtensionsCompleteMatchingPredicate(ctx, id, ExtensionOperationPredicate{})
}

// ListCalendarViewInstanceExceptionOccurrenceExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarViewInstanceExceptionOccurrenceExtensionClient) ListCalendarViewInstanceExceptionOccurrenceExtensionsCompleteMatchingPredicate(ctx context.Context, id MeCalendarViewIdInstanceIdExceptionOccurrenceId, predicate ExtensionOperationPredicate) (result ListCalendarViewInstanceExceptionOccurrenceExtensionsCompleteResult, err error) {
	items := make([]beta.Extension, 0)

	resp, err := c.ListCalendarViewInstanceExceptionOccurrenceExtensions(ctx, id)
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

	result = ListCalendarViewInstanceExceptionOccurrenceExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
