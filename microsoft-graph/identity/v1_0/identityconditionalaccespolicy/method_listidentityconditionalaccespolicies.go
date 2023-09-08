package identityconditionalaccespolicy

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

type ListIdentityConditionalAccesPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ConditionalAccessPolicyCollectionResponse
}

type ListIdentityConditionalAccesPoliciesCompleteResult struct {
	Items []models.ConditionalAccessPolicyCollectionResponse
}

// ListIdentityConditionalAccesPolicies ...
func (c IdentityConditionalAccesPolicyClient) ListIdentityConditionalAccesPolicies(ctx context.Context) (result ListIdentityConditionalAccesPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/identity/conditionalAccess/policies",
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
		Values *[]models.ConditionalAccessPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIdentityConditionalAccesPoliciesComplete retrieves all the results into a single object
func (c IdentityConditionalAccesPolicyClient) ListIdentityConditionalAccesPoliciesComplete(ctx context.Context) (ListIdentityConditionalAccesPoliciesCompleteResult, error) {
	return c.ListIdentityConditionalAccesPoliciesCompleteMatchingPredicate(ctx, models.ConditionalAccessPolicyCollectionResponseOperationPredicate{})
}

// ListIdentityConditionalAccesPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IdentityConditionalAccesPolicyClient) ListIdentityConditionalAccesPoliciesCompleteMatchingPredicate(ctx context.Context, predicate models.ConditionalAccessPolicyCollectionResponseOperationPredicate) (result ListIdentityConditionalAccesPoliciesCompleteResult, err error) {
	items := make([]models.ConditionalAccessPolicyCollectionResponse, 0)

	resp, err := c.ListIdentityConditionalAccesPolicies(ctx)
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

	result = ListIdentityConditionalAccesPoliciesCompleteResult{
		Items: items,
	}
	return
}
