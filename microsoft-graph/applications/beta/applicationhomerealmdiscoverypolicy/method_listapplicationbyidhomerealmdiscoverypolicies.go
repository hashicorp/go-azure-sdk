package applicationhomerealmdiscoverypolicy

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

type ListApplicationByIdHomeRealmDiscoveryPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.HomeRealmDiscoveryPolicyCollectionResponse
}

type ListApplicationByIdHomeRealmDiscoveryPoliciesCompleteResult struct {
	Items []models.HomeRealmDiscoveryPolicyCollectionResponse
}

// ListApplicationByIdHomeRealmDiscoveryPolicies ...
func (c ApplicationHomeRealmDiscoveryPolicyClient) ListApplicationByIdHomeRealmDiscoveryPolicies(ctx context.Context, id ApplicationId) (result ListApplicationByIdHomeRealmDiscoveryPoliciesOperationResponse, err error) {
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

// ListApplicationByIdHomeRealmDiscoveryPoliciesComplete retrieves all the results into a single object
func (c ApplicationHomeRealmDiscoveryPolicyClient) ListApplicationByIdHomeRealmDiscoveryPoliciesComplete(ctx context.Context, id ApplicationId) (ListApplicationByIdHomeRealmDiscoveryPoliciesCompleteResult, error) {
	return c.ListApplicationByIdHomeRealmDiscoveryPoliciesCompleteMatchingPredicate(ctx, id, models.HomeRealmDiscoveryPolicyCollectionResponseOperationPredicate{})
}

// ListApplicationByIdHomeRealmDiscoveryPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationHomeRealmDiscoveryPolicyClient) ListApplicationByIdHomeRealmDiscoveryPoliciesCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate models.HomeRealmDiscoveryPolicyCollectionResponseOperationPredicate) (result ListApplicationByIdHomeRealmDiscoveryPoliciesCompleteResult, err error) {
	items := make([]models.HomeRealmDiscoveryPolicyCollectionResponse, 0)

	resp, err := c.ListApplicationByIdHomeRealmDiscoveryPolicies(ctx, id)
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

	result = ListApplicationByIdHomeRealmDiscoveryPoliciesCompleteResult{
		Items: items,
	}
	return
}
