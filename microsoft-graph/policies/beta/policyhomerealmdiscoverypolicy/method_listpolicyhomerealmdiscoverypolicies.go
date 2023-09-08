package policyhomerealmdiscoverypolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPolicyHomeRealmDiscoveryPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.HomeRealmDiscoveryPolicyCollectionResponse
}

type ListPolicyHomeRealmDiscoveryPoliciesCompleteResult struct {
	Items []models.HomeRealmDiscoveryPolicyCollectionResponse
}

// ListPolicyHomeRealmDiscoveryPolicies ...
func (c PolicyHomeRealmDiscoveryPolicyClient) ListPolicyHomeRealmDiscoveryPolicies(ctx context.Context) (result ListPolicyHomeRealmDiscoveryPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/homeRealmDiscoveryPolicies",
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
		Values *[]models.HomeRealmDiscoveryPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyHomeRealmDiscoveryPoliciesComplete retrieves all the results into a single object
func (c PolicyHomeRealmDiscoveryPolicyClient) ListPolicyHomeRealmDiscoveryPoliciesComplete(ctx context.Context) (ListPolicyHomeRealmDiscoveryPoliciesCompleteResult, error) {
	return c.ListPolicyHomeRealmDiscoveryPoliciesCompleteMatchingPredicate(ctx, models.HomeRealmDiscoveryPolicyCollectionResponseOperationPredicate{})
}

// ListPolicyHomeRealmDiscoveryPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyHomeRealmDiscoveryPolicyClient) ListPolicyHomeRealmDiscoveryPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.HomeRealmDiscoveryPolicyCollectionResponseOperationPredicate) (result ListPolicyHomeRealmDiscoveryPoliciesCompleteResult, err error) {
	items := make([]models.HomeRealmDiscoveryPolicyCollectionResponse, 0)

	resp, err := c.ListPolicyHomeRealmDiscoveryPolicies(ctx)
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

	result = ListPolicyHomeRealmDiscoveryPoliciesCompleteResult{
		Items: items,
	}
	return
}
