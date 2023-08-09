package attestations

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

type AttestationsListForSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Attestation
}

type AttestationsListForSubscriptionCompleteResult struct {
	Items []Attestation
}

type AttestationsListForSubscriptionOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultAttestationsListForSubscriptionOperationOptions() AttestationsListForSubscriptionOperationOptions {
	return AttestationsListForSubscriptionOperationOptions{}
}

func (o AttestationsListForSubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AttestationsListForSubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o AttestationsListForSubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// AttestationsListForSubscription ...
func (c AttestationsClient) AttestationsListForSubscription(ctx context.Context, id commonids.SubscriptionId, options AttestationsListForSubscriptionOperationOptions) (result AttestationsListForSubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.PolicyInsights/attestations", id.ID()),
		OptionsObject: options,
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
		Values *[]Attestation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AttestationsListForSubscriptionComplete retrieves all the results into a single object
func (c AttestationsClient) AttestationsListForSubscriptionComplete(ctx context.Context, id commonids.SubscriptionId, options AttestationsListForSubscriptionOperationOptions) (AttestationsListForSubscriptionCompleteResult, error) {
	return c.AttestationsListForSubscriptionCompleteMatchingPredicate(ctx, id, options, AttestationOperationPredicate{})
}

// AttestationsListForSubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AttestationsClient) AttestationsListForSubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options AttestationsListForSubscriptionOperationOptions, predicate AttestationOperationPredicate) (result AttestationsListForSubscriptionCompleteResult, err error) {
	items := make([]Attestation, 0)

	resp, err := c.AttestationsListForSubscription(ctx, id, options)
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

	result = AttestationsListForSubscriptionCompleteResult{
		Items: items,
	}
	return
}
