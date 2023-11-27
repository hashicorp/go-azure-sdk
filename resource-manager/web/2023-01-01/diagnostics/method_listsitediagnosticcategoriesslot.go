package diagnostics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteDiagnosticCategoriesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DiagnosticCategory
}

type ListSiteDiagnosticCategoriesSlotCompleteResult struct {
	Items []DiagnosticCategory
}

// ListSiteDiagnosticCategoriesSlot ...
func (c DiagnosticsClient) ListSiteDiagnosticCategoriesSlot(ctx context.Context, id SlotId) (result ListSiteDiagnosticCategoriesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/diagnostics", id.ID()),
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
		Values *[]DiagnosticCategory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteDiagnosticCategoriesSlotComplete retrieves all the results into a single object
func (c DiagnosticsClient) ListSiteDiagnosticCategoriesSlotComplete(ctx context.Context, id SlotId) (ListSiteDiagnosticCategoriesSlotCompleteResult, error) {
	return c.ListSiteDiagnosticCategoriesSlotCompleteMatchingPredicate(ctx, id, DiagnosticCategoryOperationPredicate{})
}

// ListSiteDiagnosticCategoriesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticsClient) ListSiteDiagnosticCategoriesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate DiagnosticCategoryOperationPredicate) (result ListSiteDiagnosticCategoriesSlotCompleteResult, err error) {
	items := make([]DiagnosticCategory, 0)

	resp, err := c.ListSiteDiagnosticCategoriesSlot(ctx, id)
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

	result = ListSiteDiagnosticCategoriesSlotCompleteResult{
		Items: items,
	}
	return
}
