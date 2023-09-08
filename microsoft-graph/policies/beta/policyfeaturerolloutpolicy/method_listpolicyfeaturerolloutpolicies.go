package policyfeaturerolloutpolicy

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

type ListPolicyFeatureRolloutPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.FeatureRolloutPolicyCollectionResponse
}

type ListPolicyFeatureRolloutPoliciesCompleteResult struct {
	Items []models.FeatureRolloutPolicyCollectionResponse
}

// ListPolicyFeatureRolloutPolicies ...
func (c PolicyFeatureRolloutPolicyClient) ListPolicyFeatureRolloutPolicies(ctx context.Context) (result ListPolicyFeatureRolloutPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/featureRolloutPolicies",
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
		Values *[]models.FeatureRolloutPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyFeatureRolloutPoliciesComplete retrieves all the results into a single object
func (c PolicyFeatureRolloutPolicyClient) ListPolicyFeatureRolloutPoliciesComplete(ctx context.Context) (ListPolicyFeatureRolloutPoliciesCompleteResult, error) {
	return c.ListPolicyFeatureRolloutPoliciesCompleteMatchingPredicate(ctx, models.FeatureRolloutPolicyCollectionResponseOperationPredicate{})
}

// ListPolicyFeatureRolloutPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyFeatureRolloutPolicyClient) ListPolicyFeatureRolloutPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.FeatureRolloutPolicyCollectionResponseOperationPredicate) (result ListPolicyFeatureRolloutPoliciesCompleteResult, err error) {
	items := make([]models.FeatureRolloutPolicyCollectionResponse, 0)

	resp, err := c.ListPolicyFeatureRolloutPolicies(ctx)
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

	result = ListPolicyFeatureRolloutPoliciesCompleteResult{
		Items: items,
	}
	return
}
