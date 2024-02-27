package billingroleassignments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddByInvoiceSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRoleAssignment
}

type AddByInvoiceSectionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleAssignment
}

// AddByInvoiceSection ...
func (c BillingRoleAssignmentsClient) AddByInvoiceSection(ctx context.Context, id InvoiceSectionId, input BillingRoleAssignmentPayload) (result AddByInvoiceSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/createBillingRoleAssignment", id.ID()),
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
		Values *[]BillingRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AddByInvoiceSectionComplete retrieves all the results into a single object
func (c BillingRoleAssignmentsClient) AddByInvoiceSectionComplete(ctx context.Context, id InvoiceSectionId, input BillingRoleAssignmentPayload) (AddByInvoiceSectionCompleteResult, error) {
	return c.AddByInvoiceSectionCompleteMatchingPredicate(ctx, id, input, BillingRoleAssignmentOperationPredicate{})
}

// AddByInvoiceSectionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingRoleAssignmentsClient) AddByInvoiceSectionCompleteMatchingPredicate(ctx context.Context, id InvoiceSectionId, input BillingRoleAssignmentPayload, predicate BillingRoleAssignmentOperationPredicate) (result AddByInvoiceSectionCompleteResult, err error) {
	items := make([]BillingRoleAssignment, 0)

	resp, err := c.AddByInvoiceSection(ctx, id, input)
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

	result = AddByInvoiceSectionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
