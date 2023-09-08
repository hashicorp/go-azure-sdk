package serviceprincipalhomerealmdiscoverypolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListServicePrincipalByIdHomeRealmDiscoveryPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.HomeRealmDiscoveryPolicyCollectionResponse
}

type ListServicePrincipalByIdHomeRealmDiscoveryPoliciesCompleteResult struct {
	Items []models.HomeRealmDiscoveryPolicyCollectionResponse
}

// ListServicePrincipalByIdHomeRealmDiscoveryPolicies ...
func (c ServicePrincipalHomeRealmDiscoveryPolicyClient) ListServicePrincipalByIdHomeRealmDiscoveryPolicies(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdHomeRealmDiscoveryPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/homeRealmDiscoveryPolicies", id.ID()),
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

// ListServicePrincipalByIdHomeRealmDiscoveryPoliciesComplete retrieves all the results into a single object
func (c ServicePrincipalHomeRealmDiscoveryPolicyClient) ListServicePrincipalByIdHomeRealmDiscoveryPoliciesComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdHomeRealmDiscoveryPoliciesCompleteResult, error) {
	return c.ListServicePrincipalByIdHomeRealmDiscoveryPoliciesCompleteMatchingPredicate(ctx, id, models.HomeRealmDiscoveryPolicyCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdHomeRealmDiscoveryPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalHomeRealmDiscoveryPolicyClient) ListServicePrincipalByIdHomeRealmDiscoveryPoliciesCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.HomeRealmDiscoveryPolicyCollectionResponseOperationPredicate) (result ListServicePrincipalByIdHomeRealmDiscoveryPoliciesCompleteResult, err error) {
	items := make([]models.HomeRealmDiscoveryPolicyCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdHomeRealmDiscoveryPolicies(ctx, id)
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

	result = ListServicePrincipalByIdHomeRealmDiscoveryPoliciesCompleteResult{
		Items: items,
	}
	return
}
