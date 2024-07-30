package exchangeonpremisespolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListExchangeOnPremisesPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementExchangeOnPremisesPolicy
}

type ListExchangeOnPremisesPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementExchangeOnPremisesPolicy
}

type ListExchangeOnPremisesPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExchangeOnPremisesPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExchangeOnPremisesPolicies ...
func (c ExchangeOnPremisesPolicyClient) ListExchangeOnPremisesPolicies(ctx context.Context) (result ListExchangeOnPremisesPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListExchangeOnPremisesPoliciesCustomPager{},
		Path:       "/deviceManagement/exchangeOnPremisesPolicies",
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
		Values *[]beta.DeviceManagementExchangeOnPremisesPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExchangeOnPremisesPoliciesComplete retrieves all the results into a single object
func (c ExchangeOnPremisesPolicyClient) ListExchangeOnPremisesPoliciesComplete(ctx context.Context) (ListExchangeOnPremisesPoliciesCompleteResult, error) {
	return c.ListExchangeOnPremisesPoliciesCompleteMatchingPredicate(ctx, DeviceManagementExchangeOnPremisesPolicyOperationPredicate{})
}

// ListExchangeOnPremisesPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExchangeOnPremisesPolicyClient) ListExchangeOnPremisesPoliciesCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementExchangeOnPremisesPolicyOperationPredicate) (result ListExchangeOnPremisesPoliciesCompleteResult, err error) {
	items := make([]beta.DeviceManagementExchangeOnPremisesPolicy, 0)

	resp, err := c.ListExchangeOnPremisesPolicies(ctx)
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

	result = ListExchangeOnPremisesPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
