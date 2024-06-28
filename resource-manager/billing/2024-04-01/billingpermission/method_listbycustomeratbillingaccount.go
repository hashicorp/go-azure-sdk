package billingpermission

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByCustomerAtBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingPermission
}

type ListByCustomerAtBillingAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingPermission
}

type ListByCustomerAtBillingAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByCustomerAtBillingAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByCustomerAtBillingAccount ...
func (c BillingPermissionClient) ListByCustomerAtBillingAccount(ctx context.Context, id CustomerId) (result ListByCustomerAtBillingAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByCustomerAtBillingAccountCustomPager{},
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
		Values *[]BillingPermission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByCustomerAtBillingAccountComplete retrieves all the results into a single object
func (c BillingPermissionClient) ListByCustomerAtBillingAccountComplete(ctx context.Context, id CustomerId) (ListByCustomerAtBillingAccountCompleteResult, error) {
	return c.ListByCustomerAtBillingAccountCompleteMatchingPredicate(ctx, id, BillingPermissionOperationPredicate{})
}

// ListByCustomerAtBillingAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingPermissionClient) ListByCustomerAtBillingAccountCompleteMatchingPredicate(ctx context.Context, id CustomerId, predicate BillingPermissionOperationPredicate) (result ListByCustomerAtBillingAccountCompleteResult, err error) {
	items := make([]BillingPermission, 0)

	resp, err := c.ListByCustomerAtBillingAccount(ctx, id)
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

	result = ListByCustomerAtBillingAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
