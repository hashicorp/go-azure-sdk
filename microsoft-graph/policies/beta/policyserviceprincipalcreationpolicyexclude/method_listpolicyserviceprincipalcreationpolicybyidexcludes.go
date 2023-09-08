package policyserviceprincipalcreationpolicyexclude

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

type ListPolicyServicePrincipalCreationPolicyByIdExcludesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ServicePrincipalCreationConditionSetCollectionResponse
}

type ListPolicyServicePrincipalCreationPolicyByIdExcludesCompleteResult struct {
	Items []models.ServicePrincipalCreationConditionSetCollectionResponse
}

// ListPolicyServicePrincipalCreationPolicyByIdExcludes ...
func (c PolicyServicePrincipalCreationPolicyExcludeClient) ListPolicyServicePrincipalCreationPolicyByIdExcludes(ctx context.Context, id PolicyServicePrincipalCreationPolicyId) (result ListPolicyServicePrincipalCreationPolicyByIdExcludesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/excludes", id.ID()),
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

// ListPolicyServicePrincipalCreationPolicyByIdExcludesComplete retrieves all the results into a single object
func (c PolicyServicePrincipalCreationPolicyExcludeClient) ListPolicyServicePrincipalCreationPolicyByIdExcludesComplete(ctx context.Context, id PolicyServicePrincipalCreationPolicyId) (ListPolicyServicePrincipalCreationPolicyByIdExcludesCompleteResult, error) {
	return c.ListPolicyServicePrincipalCreationPolicyByIdExcludesCompleteMatchingPredicate(ctx, id, models.ServicePrincipalCreationConditionSetCollectionResponseOperationPredicate{})
}

// ListPolicyServicePrincipalCreationPolicyByIdExcludesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyServicePrincipalCreationPolicyExcludeClient) ListPolicyServicePrincipalCreationPolicyByIdExcludesCompleteMatchingPredicate(ctx context.Context, id PolicyServicePrincipalCreationPolicyId, predicate models.ServicePrincipalCreationConditionSetCollectionResponseOperationPredicate) (result ListPolicyServicePrincipalCreationPolicyByIdExcludesCompleteResult, err error) {
	items := make([]models.ServicePrincipalCreationConditionSetCollectionResponse, 0)

	resp, err := c.ListPolicyServicePrincipalCreationPolicyByIdExcludes(ctx, id)
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

	result = ListPolicyServicePrincipalCreationPolicyByIdExcludesCompleteResult{
		Items: items,
	}
	return
}
