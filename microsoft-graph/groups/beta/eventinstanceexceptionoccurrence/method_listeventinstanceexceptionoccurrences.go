package eventinstanceexceptionoccurrence

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

type ListEventInstanceExceptionOccurrencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Event
}

type ListEventInstanceExceptionOccurrencesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Event
}

type ListEventInstanceExceptionOccurrencesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEventInstanceExceptionOccurrencesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEventInstanceExceptionOccurrences ...
func (c EventInstanceExceptionOccurrenceClient) ListEventInstanceExceptionOccurrences(ctx context.Context, id GroupIdEventIdInstanceId) (result ListEventInstanceExceptionOccurrencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEventInstanceExceptionOccurrencesCustomPager{},
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

// ListEventInstanceExceptionOccurrencesComplete retrieves all the results into a single object
func (c EventInstanceExceptionOccurrenceClient) ListEventInstanceExceptionOccurrencesComplete(ctx context.Context, id GroupIdEventIdInstanceId) (ListEventInstanceExceptionOccurrencesCompleteResult, error) {
	return c.ListEventInstanceExceptionOccurrencesCompleteMatchingPredicate(ctx, id, EventOperationPredicate{})
}

// ListEventInstanceExceptionOccurrencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EventInstanceExceptionOccurrenceClient) ListEventInstanceExceptionOccurrencesCompleteMatchingPredicate(ctx context.Context, id GroupIdEventIdInstanceId, predicate EventOperationPredicate) (result ListEventInstanceExceptionOccurrencesCompleteResult, err error) {
	items := make([]beta.Event, 0)

	resp, err := c.ListEventInstanceExceptionOccurrences(ctx, id)
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

	result = ListEventInstanceExceptionOccurrencesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
