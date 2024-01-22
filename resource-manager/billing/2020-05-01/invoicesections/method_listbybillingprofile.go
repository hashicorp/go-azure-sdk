package invoicesections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]InvoiceSection
}

type ListByBillingProfileCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []InvoiceSection
}

// ListByBillingProfile ...
func (c InvoiceSectionsClient) ListByBillingProfile(ctx context.Context, id BillingProfileId) (result ListByBillingProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/invoiceSections", id.ID()),
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
		Values *[]InvoiceSection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBillingProfileComplete retrieves all the results into a single object
func (c InvoiceSectionsClient) ListByBillingProfileComplete(ctx context.Context, id BillingProfileId) (ListByBillingProfileCompleteResult, error) {
	return c.ListByBillingProfileCompleteMatchingPredicate(ctx, id, InvoiceSectionOperationPredicate{})
}

// ListByBillingProfileCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InvoiceSectionsClient) ListByBillingProfileCompleteMatchingPredicate(ctx context.Context, id BillingProfileId, predicate InvoiceSectionOperationPredicate) (result ListByBillingProfileCompleteResult, err error) {
	items := make([]InvoiceSection, 0)

	resp, err := c.ListByBillingProfile(ctx, id)
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

	result = ListByBillingProfileCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
