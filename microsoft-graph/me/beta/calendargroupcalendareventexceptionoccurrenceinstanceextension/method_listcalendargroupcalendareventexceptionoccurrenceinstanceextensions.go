package calendargroupcalendareventexceptionoccurrenceinstanceextension

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

type ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Extension
}

type ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Extension
}

type ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensions ...
func (c CalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient) ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensions(ctx context.Context, id MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceId) (result ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsCustomPager{},
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

// ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsComplete retrieves all the results into a single object
func (c CalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient) ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsComplete(ctx context.Context, id MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceId) (ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsCompleteResult, error) {
	return c.ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsCompleteMatchingPredicate(ctx, id, ExtensionOperationPredicate{})
}

// ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient) ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsCompleteMatchingPredicate(ctx context.Context, id MeCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdInstanceId, predicate ExtensionOperationPredicate) (result ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsCompleteResult, err error) {
	items := make([]beta.Extension, 0)

	resp, err := c.ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensions(ctx, id)
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

	result = ListCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
