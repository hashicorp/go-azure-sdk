package calendargroupcalendareventinstanceexceptionoccurrenceextension

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

type ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Extension
}

type ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Extension
}

type ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensions ...
func (c CalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient) ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensions(ctx context.Context, id MeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceId) (result ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsCustomPager{},
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

// ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsComplete retrieves all the results into a single object
func (c CalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient) ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsComplete(ctx context.Context, id MeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceId) (ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsCompleteResult, error) {
	return c.ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsCompleteMatchingPredicate(ctx, id, ExtensionOperationPredicate{})
}

// ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient) ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsCompleteMatchingPredicate(ctx context.Context, id MeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceId, predicate ExtensionOperationPredicate) (result ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsCompleteResult, err error) {
	items := make([]beta.Extension, 0)

	resp, err := c.ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensions(ctx, id)
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

	result = ListCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
