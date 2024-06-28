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

type ListByEnrollmentAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingPermissionsProperties
}

type ListByEnrollmentAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingPermissionsProperties
}

type ListByEnrollmentAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByEnrollmentAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByEnrollmentAccount ...
func (c BillingPermissionsClient) ListByEnrollmentAccount(ctx context.Context, id EnrollmentAccountId) (result ListByEnrollmentAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByEnrollmentAccountCustomPager{},
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

// ListByEnrollmentAccountComplete retrieves all the results into a single object
func (c BillingPermissionsClient) ListByEnrollmentAccountComplete(ctx context.Context, id EnrollmentAccountId) (ListByEnrollmentAccountCompleteResult, error) {
	return c.ListByEnrollmentAccountCompleteMatchingPredicate(ctx, id, BillingPermissionsPropertiesOperationPredicate{})
}

// ListByEnrollmentAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingPermissionsClient) ListByEnrollmentAccountCompleteMatchingPredicate(ctx context.Context, id EnrollmentAccountId, predicate BillingPermissionsPropertiesOperationPredicate) (result ListByEnrollmentAccountCompleteResult, err error) {
	items := make([]BillingPermissionsProperties, 0)

	resp, err := c.ListByEnrollmentAccount(ctx, id)
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

	result = ListByEnrollmentAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
