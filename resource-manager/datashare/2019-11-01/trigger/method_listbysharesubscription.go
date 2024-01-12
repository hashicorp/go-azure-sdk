package trigger

import (
	"context"
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

// ListByShareSubscription ...
func (c TriggerClient) ListByShareSubscription(ctx context.Context, id ShareSubscriptionId) (result ListByShareSubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]Trigger `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

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
