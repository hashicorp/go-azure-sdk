package serviceprincipalappmanagementpolicy

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

type ListServicePrincipalByIdAppManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AppManagementPolicyCollectionResponse
}

type ListServicePrincipalByIdAppManagementPoliciesCompleteResult struct {
	Items []models.AppManagementPolicyCollectionResponse
}

// ListServicePrincipalByIdAppManagementPolicies ...
func (c ServicePrincipalAppManagementPolicyClient) ListServicePrincipalByIdAppManagementPolicies(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdAppManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/appManagementPolicies", id.ID()),
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
		Values *[]models.AppManagementPolicyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServicePrincipalByIdAppManagementPoliciesComplete retrieves all the results into a single object
func (c ServicePrincipalAppManagementPolicyClient) ListServicePrincipalByIdAppManagementPoliciesComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdAppManagementPoliciesCompleteResult, error) {
	return c.ListServicePrincipalByIdAppManagementPoliciesCompleteMatchingPredicate(ctx, id, models.AppManagementPolicyCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdAppManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalAppManagementPolicyClient) ListServicePrincipalByIdAppManagementPoliciesCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.AppManagementPolicyCollectionResponseOperationPredicate) (result ListServicePrincipalByIdAppManagementPoliciesCompleteResult, err error) {
	items := make([]models.AppManagementPolicyCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdAppManagementPolicies(ctx, id)
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

	result = ListServicePrincipalByIdAppManagementPoliciesCompleteResult{
		Items: items,
	}
	return
}
