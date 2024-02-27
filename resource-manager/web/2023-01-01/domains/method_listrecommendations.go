package domains

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRecommendationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]NameIdentifier
}

type ListRecommendationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []NameIdentifier
}

// ListRecommendations ...
func (c DomainsClient) ListRecommendations(ctx context.Context, id commonids.SubscriptionId, input DomainRecommendationSearchParameters) (result ListRecommendationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.DomainRegistration/listDomainRecommendations", id.ID()),
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
		Values *[]NameIdentifier `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRecommendationsComplete retrieves all the results into a single object
func (c DomainsClient) ListRecommendationsComplete(ctx context.Context, id commonids.SubscriptionId, input DomainRecommendationSearchParameters) (ListRecommendationsCompleteResult, error) {
	return c.ListRecommendationsCompleteMatchingPredicate(ctx, id, input, NameIdentifierOperationPredicate{})
}

// ListRecommendationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DomainsClient) ListRecommendationsCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, input DomainRecommendationSearchParameters, predicate NameIdentifierOperationPredicate) (result ListRecommendationsCompleteResult, err error) {
	items := make([]NameIdentifier, 0)

	resp, err := c.ListRecommendations(ctx, id, input)
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

	result = ListRecommendationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
