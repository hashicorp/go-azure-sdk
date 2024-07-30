package calendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension

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

type ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Extension
}

type ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Extension
}

type ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensions ...
func (c CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient) ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensions(ctx context.Context, id MeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId) (result ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsCustomPager{},
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

// ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsComplete retrieves all the results into a single object
func (c CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient) ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsComplete(ctx context.Context, id MeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId) (ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsCompleteResult, error) {
	return c.ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsCompleteMatchingPredicate(ctx, id, ExtensionOperationPredicate{})
}

// ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient) ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsCompleteMatchingPredicate(ctx context.Context, id MeCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId, predicate ExtensionOperationPredicate) (result ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsCompleteResult, err error) {
	items := make([]beta.Extension, 0)

	resp, err := c.ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensions(ctx, id)
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

	result = ListCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
