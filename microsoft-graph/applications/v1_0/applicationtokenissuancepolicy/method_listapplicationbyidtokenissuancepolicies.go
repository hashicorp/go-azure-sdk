package applicationtokenissuancepolicy

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

type ListApplicationByIdTokenIssuancePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TokenIssuancePolicyCollectionResponse
}

type ListApplicationByIdTokenIssuancePoliciesCompleteResult struct {
	Items []models.TokenIssuancePolicyCollectionResponse
}

// ListApplicationByIdTokenIssuancePolicies ...
func (c ApplicationTokenIssuancePolicyClient) ListApplicationByIdTokenIssuancePolicies(ctx context.Context, id ApplicationId) (result ListApplicationByIdTokenIssuancePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/tokenIssuancePolicies", id.ID()),
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
		Values *[]models.TokenIssuancePolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListApplicationByIdTokenIssuancePoliciesComplete retrieves all the results into a single object
func (c ApplicationTokenIssuancePolicyClient) ListApplicationByIdTokenIssuancePoliciesComplete(ctx context.Context, id ApplicationId) (ListApplicationByIdTokenIssuancePoliciesCompleteResult, error) {
	return c.ListApplicationByIdTokenIssuancePoliciesCompleteMatchingPredicate(ctx, id, models.TokenIssuancePolicyCollectionResponseOperationPredicate{})
}

// ListApplicationByIdTokenIssuancePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationTokenIssuancePolicyClient) ListApplicationByIdTokenIssuancePoliciesCompleteMatchingPredicate(ctx context.Context, id ApplicationId, predicate models.TokenIssuancePolicyCollectionResponseOperationPredicate) (result ListApplicationByIdTokenIssuancePoliciesCompleteResult, err error) {
	items := make([]models.TokenIssuancePolicyCollectionResponse, 0)

	resp, err := c.ListApplicationByIdTokenIssuancePolicies(ctx, id)
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

	result = ListApplicationByIdTokenIssuancePoliciesCompleteResult{
		Items: items,
	}
	return
}
