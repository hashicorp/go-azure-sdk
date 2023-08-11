package subscriptionfeatureregistrations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionFeatureRegistrationsListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubscriptionFeatureRegistration
}

type SubscriptionFeatureRegistrationsListBySubscriptionCompleteResult struct {
	Items []SubscriptionFeatureRegistration
}

// SubscriptionFeatureRegistrationsListBySubscription ...
func (c SubscriptionFeatureRegistrationsClient) SubscriptionFeatureRegistrationsListBySubscription(ctx context.Context, id FeatureProviderId) (result SubscriptionFeatureRegistrationsListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/subscriptionFeatureRegistrations", id.ID()),
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
		Values *[]SubscriptionFeatureRegistration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SubscriptionFeatureRegistrationsListBySubscriptionComplete retrieves all the results into a single object
func (c SubscriptionFeatureRegistrationsClient) SubscriptionFeatureRegistrationsListBySubscriptionComplete(ctx context.Context, id FeatureProviderId) (SubscriptionFeatureRegistrationsListBySubscriptionCompleteResult, error) {
	return c.SubscriptionFeatureRegistrationsListBySubscriptionCompleteMatchingPredicate(ctx, id, SubscriptionFeatureRegistrationOperationPredicate{})
}

// SubscriptionFeatureRegistrationsListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SubscriptionFeatureRegistrationsClient) SubscriptionFeatureRegistrationsListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id FeatureProviderId, predicate SubscriptionFeatureRegistrationOperationPredicate) (result SubscriptionFeatureRegistrationsListBySubscriptionCompleteResult, err error) {
	items := make([]SubscriptionFeatureRegistration, 0)

	resp, err := c.SubscriptionFeatureRegistrationsListBySubscription(ctx, id)
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

	result = SubscriptionFeatureRegistrationsListBySubscriptionCompleteResult{
		Items: items,
	}
	return
}
