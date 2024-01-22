package billingpermissions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByInvoiceSectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingPermissionsProperties
}

type ListByInvoiceSectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingPermissionsProperties
}

// ListByInvoiceSections ...
func (c BillingPermissionsClient) ListByInvoiceSections(ctx context.Context, id InvoiceSectionId) (result ListByInvoiceSectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/billingPermissions", id.ID()),
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
		Values *[]BillingPermissionsProperties `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInvoiceSectionsComplete retrieves all the results into a single object
func (c BillingPermissionsClient) ListByInvoiceSectionsComplete(ctx context.Context, id InvoiceSectionId) (ListByInvoiceSectionsCompleteResult, error) {
	return c.ListByInvoiceSectionsCompleteMatchingPredicate(ctx, id, BillingPermissionsPropertiesOperationPredicate{})
}

// ListByInvoiceSectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingPermissionsClient) ListByInvoiceSectionsCompleteMatchingPredicate(ctx context.Context, id InvoiceSectionId, predicate BillingPermissionsPropertiesOperationPredicate) (result ListByInvoiceSectionsCompleteResult, err error) {
	items := make([]BillingPermissionsProperties, 0)

	resp, err := c.ListByInvoiceSections(ctx, id)
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

	result = ListByInvoiceSectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
