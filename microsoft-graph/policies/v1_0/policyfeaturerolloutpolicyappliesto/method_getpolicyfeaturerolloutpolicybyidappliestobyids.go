package policyfeaturerolloutpolicyappliesto

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

type GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObject
}

type GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsCompleteResult struct {
	Items []models.DirectoryObject
}

// GetPolicyFeatureRolloutPolicyByIdAppliesToByIds ...
func (c PolicyFeatureRolloutPolicyAppliesToClient) GetPolicyFeatureRolloutPolicyByIdAppliesToByIds(ctx context.Context, id PolicyFeatureRolloutPolicyId) (result GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/appliesTo/getByIds", id.ID()),
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
		Values *[]models.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsComplete retrieves all the results into a single object
func (c PolicyFeatureRolloutPolicyAppliesToClient) GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsComplete(ctx context.Context, id PolicyFeatureRolloutPolicyId) (GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsCompleteResult, error) {
	return c.GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsCompleteMatchingPredicate(ctx, id, models.DirectoryObjectOperationPredicate{})
}

// GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyFeatureRolloutPolicyAppliesToClient) GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsCompleteMatchingPredicate(ctx context.Context, id PolicyFeatureRolloutPolicyId, predicate models.DirectoryObjectOperationPredicate) (result GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsCompleteResult, err error) {
	items := make([]models.DirectoryObject, 0)

	resp, err := c.GetPolicyFeatureRolloutPolicyByIdAppliesToByIds(ctx, id)
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

	result = GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsCompleteResult{
		Items: items,
	}
	return
}
