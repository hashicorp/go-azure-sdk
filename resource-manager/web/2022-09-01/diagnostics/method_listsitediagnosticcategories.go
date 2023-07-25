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

type ListSiteDiagnosticCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DiagnosticCategory
}

type ListSiteDiagnosticCategoriesCompleteResult struct {
	Items []DiagnosticCategory
}

// ListSiteDiagnosticCategories ...
func (c DiagnosticsClient) ListSiteDiagnosticCategories(ctx context.Context, id SiteId) (result ListSiteDiagnosticCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
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

// ListSiteDiagnosticCategoriesComplete retrieves all the results into a single object
func (c DiagnosticsClient) ListSiteDiagnosticCategoriesComplete(ctx context.Context, id SiteId) (ListSiteDiagnosticCategoriesCompleteResult, error) {
	return c.ListSiteDiagnosticCategoriesCompleteMatchingPredicate(ctx, id, DiagnosticCategoryOperationPredicate{})
}

// ListSiteDiagnosticCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticsClient) ListSiteDiagnosticCategoriesCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate DiagnosticCategoryOperationPredicate) (result ListSiteDiagnosticCategoriesCompleteResult, err error) {
	items := make([]DiagnosticCategory, 0)

	resp, err := c.ListSiteDiagnosticCategories(ctx, id)
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

	result = ListSiteDiagnosticCategoriesCompleteResult{
		Items: items,
	}
	return
}
