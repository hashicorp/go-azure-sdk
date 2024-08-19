package listtenantconfigurationviolationsoperations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTenantConfigurationViolationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Violation
}

type ListTenantConfigurationViolationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Violation
}

type ListTenantConfigurationViolationsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListTenantConfigurationViolationsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTenantConfigurationViolationsList ...
func (c ListTenantConfigurationViolationsOperationsClient) ListTenantConfigurationViolationsList(ctx context.Context) (result ListTenantConfigurationViolationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListTenantConfigurationViolationsListCustomPager{},
		Path:       "/providers/Microsoft.Portal/listTenantConfigurationViolations",
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
		Values *[]Violation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTenantConfigurationViolationsListComplete retrieves all the results into a single object
func (c ListTenantConfigurationViolationsOperationsClient) ListTenantConfigurationViolationsListComplete(ctx context.Context) (ListTenantConfigurationViolationsListCompleteResult, error) {
	return c.ListTenantConfigurationViolationsListCompleteMatchingPredicate(ctx, ViolationOperationPredicate{})
}

// ListTenantConfigurationViolationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ListTenantConfigurationViolationsOperationsClient) ListTenantConfigurationViolationsListCompleteMatchingPredicate(ctx context.Context, predicate ViolationOperationPredicate) (result ListTenantConfigurationViolationsListCompleteResult, err error) {
	items := make([]Violation, 0)

	resp, err := c.ListTenantConfigurationViolationsList(ctx)
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

	result = ListTenantConfigurationViolationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
