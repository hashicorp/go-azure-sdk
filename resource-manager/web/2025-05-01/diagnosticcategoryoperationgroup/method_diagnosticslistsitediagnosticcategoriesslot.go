package diagnosticcategoryoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsListSiteDiagnosticCategoriesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DiagnosticCategory
}

type DiagnosticsListSiteDiagnosticCategoriesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DiagnosticCategory
}

type DiagnosticsListSiteDiagnosticCategoriesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DiagnosticsListSiteDiagnosticCategoriesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DiagnosticsListSiteDiagnosticCategoriesSlot ...
func (c DiagnosticCategoryOperationGroupClient) DiagnosticsListSiteDiagnosticCategoriesSlot(ctx context.Context, id SlotId) (result DiagnosticsListSiteDiagnosticCategoriesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DiagnosticsListSiteDiagnosticCategoriesSlotCustomPager{},
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

// DiagnosticsListSiteDiagnosticCategoriesSlotComplete retrieves all the results into a single object
func (c DiagnosticCategoryOperationGroupClient) DiagnosticsListSiteDiagnosticCategoriesSlotComplete(ctx context.Context, id SlotId) (DiagnosticsListSiteDiagnosticCategoriesSlotCompleteResult, error) {
	return c.DiagnosticsListSiteDiagnosticCategoriesSlotCompleteMatchingPredicate(ctx, id, DiagnosticCategoryOperationPredicate{})
}

// DiagnosticsListSiteDiagnosticCategoriesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticCategoryOperationGroupClient) DiagnosticsListSiteDiagnosticCategoriesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate DiagnosticCategoryOperationPredicate) (result DiagnosticsListSiteDiagnosticCategoriesSlotCompleteResult, err error) {
	items := make([]DiagnosticCategory, 0)

	resp, err := c.DiagnosticsListSiteDiagnosticCategoriesSlot(ctx, id)
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

	result = DiagnosticsListSiteDiagnosticCategoriesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
