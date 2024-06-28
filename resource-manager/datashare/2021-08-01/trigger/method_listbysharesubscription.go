package trigger

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByShareSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Trigger
}

type ListByShareSubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Trigger
}

type ListByShareSubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByShareSubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByShareSubscription ...
func (c TriggerClient) ListByShareSubscription(ctx context.Context, id ShareSubscriptionId) (result ListByShareSubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByShareSubscriptionCustomPager{},
		Path:       fmt.Sprintf("%s/triggers", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]Trigger, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := unmarshalTriggerImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for Trigger (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListByShareSubscriptionComplete retrieves all the results into a single object
func (c TriggerClient) ListByShareSubscriptionComplete(ctx context.Context, id ShareSubscriptionId) (ListByShareSubscriptionCompleteResult, error) {
	return c.ListByShareSubscriptionCompleteMatchingPredicate(ctx, id, TriggerOperationPredicate{})
}

// ListByShareSubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TriggerClient) ListByShareSubscriptionCompleteMatchingPredicate(ctx context.Context, id ShareSubscriptionId, predicate TriggerOperationPredicate) (result ListByShareSubscriptionCompleteResult, err error) {
	items := make([]Trigger, 0)

	resp, err := c.ListByShareSubscription(ctx, id)
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

	result = ListByShareSubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
