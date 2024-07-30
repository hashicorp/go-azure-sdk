package eventexceptionoccurrence

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

type ListEventExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Event
}

type ListEventExceptionOccurrencesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Event
}

type ListEventExceptionOccurrencesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEventExceptionOccurrencesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEventExceptionOccurrences ...
func (c EventExceptionOccurrenceClient) ListEventExceptionOccurrences(ctx context.Context, id GroupIdEventId) (result ListEventExceptionOccurrencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEventExceptionOccurrencesCustomPager{},
		Path:       fmt.Sprintf("%s/exceptionOccurrences", id.ID()),
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
		Values *[]beta.Event `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEventExceptionOccurrencesComplete retrieves all the results into a single object
func (c EventExceptionOccurrenceClient) ListEventExceptionOccurrencesComplete(ctx context.Context, id GroupIdEventId) (ListEventExceptionOccurrencesCompleteResult, error) {
	return c.ListEventExceptionOccurrencesCompleteMatchingPredicate(ctx, id, EventOperationPredicate{})
}

// ListEventExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EventExceptionOccurrenceClient) ListEventExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id GroupIdEventId, predicate EventOperationPredicate) (result ListEventExceptionOccurrencesCompleteResult, err error) {
	items := make([]beta.Event, 0)

	resp, err := c.ListEventExceptionOccurrences(ctx, id)
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

	result = ListEventExceptionOccurrencesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
