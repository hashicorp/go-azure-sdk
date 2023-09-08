package policyactivitybasedtimeoutpolicy

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

type ListPolicyActivityBasedTimeoutPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ActivityBasedTimeoutPolicyCollectionResponse
}

type ListPolicyActivityBasedTimeoutPoliciesCompleteResult struct {
	Items []models.ActivityBasedTimeoutPolicyCollectionResponse
}

// ListPolicyActivityBasedTimeoutPolicies ...
func (c PolicyActivityBasedTimeoutPolicyClient) ListPolicyActivityBasedTimeoutPolicies(ctx context.Context) (result ListPolicyActivityBasedTimeoutPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/policies/activityBasedTimeoutPolicies",
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
		Values *[]models.ActivityBasedTimeoutPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyActivityBasedTimeoutPoliciesComplete retrieves all the results into a single object
func (c PolicyActivityBasedTimeoutPolicyClient) ListPolicyActivityBasedTimeoutPoliciesComplete(ctx context.Context) (ListPolicyActivityBasedTimeoutPoliciesCompleteResult, error) {
	return c.ListPolicyActivityBasedTimeoutPoliciesCompleteMatchingPredicate(ctx, models.ActivityBasedTimeoutPolicyCollectionResponseOperationPredicate{})
}

// ListPolicyActivityBasedTimeoutPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyActivityBasedTimeoutPolicyClient) ListPolicyActivityBasedTimeoutPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.ActivityBasedTimeoutPolicyCollectionResponseOperationPredicate) (result ListPolicyActivityBasedTimeoutPoliciesCompleteResult, err error) {
	items := make([]models.ActivityBasedTimeoutPolicyCollectionResponse, 0)

	resp, err := c.ListPolicyActivityBasedTimeoutPolicies(ctx)
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

	result = ListPolicyActivityBasedTimeoutPoliciesCompleteResult{
		Items: items,
	}
	return
}
