package diagnosticcategories

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

type DiagnosticsListSiteDiagnosticCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DiagnosticCategory
}

type DiagnosticsListSiteDiagnosticCategoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DiagnosticCategory
}

type DiagnosticsListSiteDiagnosticCategoriesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DiagnosticsListSiteDiagnosticCategoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DiagnosticsListSiteDiagnosticCategories ...
func (c DiagnosticCategoriesClient) DiagnosticsListSiteDiagnosticCategories(ctx context.Context, id commonids.AppServiceId) (result DiagnosticsListSiteDiagnosticCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DiagnosticsListSiteDiagnosticCategoriesCustomPager{},
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

// DiagnosticsListSiteDiagnosticCategoriesComplete retrieves all the results into a single object
func (c DiagnosticCategoriesClient) DiagnosticsListSiteDiagnosticCategoriesComplete(ctx context.Context, id commonids.AppServiceId) (DiagnosticsListSiteDiagnosticCategoriesCompleteResult, error) {
	return c.DiagnosticsListSiteDiagnosticCategoriesCompleteMatchingPredicate(ctx, id, DiagnosticCategoryOperationPredicate{})
}

// DiagnosticsListSiteDiagnosticCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticCategoriesClient) DiagnosticsListSiteDiagnosticCategoriesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate DiagnosticCategoryOperationPredicate) (result DiagnosticsListSiteDiagnosticCategoriesCompleteResult, err error) {
	items := make([]DiagnosticCategory, 0)

	resp, err := c.DiagnosticsListSiteDiagnosticCategories(ctx, id)
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

	result = DiagnosticsListSiteDiagnosticCategoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
