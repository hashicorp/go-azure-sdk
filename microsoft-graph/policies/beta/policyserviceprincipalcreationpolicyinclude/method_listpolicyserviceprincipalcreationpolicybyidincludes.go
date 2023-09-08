package policyserviceprincipalcreationpolicyinclude

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

type ListPolicyServicePrincipalCreationPolicyByIdIncludesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ServicePrincipalCreationConditionSetCollectionResponse
}

type ListPolicyServicePrincipalCreationPolicyByIdIncludesCompleteResult struct {
	Items []models.ServicePrincipalCreationConditionSetCollectionResponse
}

// ListPolicyServicePrincipalCreationPolicyByIdIncludes ...
func (c PolicyServicePrincipalCreationPolicyIncludeClient) ListPolicyServicePrincipalCreationPolicyByIdIncludes(ctx context.Context, id PolicyServicePrincipalCreationPolicyId) (result ListPolicyServicePrincipalCreationPolicyByIdIncludesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/includes", id.ID()),
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
		Values *[]models.ServicePrincipalCreationConditionSetCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPolicyServicePrincipalCreationPolicyByIdIncludesComplete retrieves all the results into a single object
func (c PolicyServicePrincipalCreationPolicyIncludeClient) ListPolicyServicePrincipalCreationPolicyByIdIncludesComplete(ctx context.Context, id PolicyServicePrincipalCreationPolicyId) (ListPolicyServicePrincipalCreationPolicyByIdIncludesCompleteResult, error) {
	return c.ListPolicyServicePrincipalCreationPolicyByIdIncludesCompleteMatchingPredicate(ctx, id, models.ServicePrincipalCreationConditionSetCollectionResponseOperationPredicate{})
}

// ListPolicyServicePrincipalCreationPolicyByIdIncludesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyServicePrincipalCreationPolicyIncludeClient) ListPolicyServicePrincipalCreationPolicyByIdIncludesCompleteMatchingPredicate(ctx context.Context, id PolicyServicePrincipalCreationPolicyId, predicate models.ServicePrincipalCreationConditionSetCollectionResponseOperationPredicate) (result ListPolicyServicePrincipalCreationPolicyByIdIncludesCompleteResult, err error) {
	items := make([]models.ServicePrincipalCreationConditionSetCollectionResponse, 0)

	resp, err := c.ListPolicyServicePrincipalCreationPolicyByIdIncludes(ctx, id)
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

	result = ListPolicyServicePrincipalCreationPolicyByIdIncludesCompleteResult{
		Items: items,
	}
	return
}
