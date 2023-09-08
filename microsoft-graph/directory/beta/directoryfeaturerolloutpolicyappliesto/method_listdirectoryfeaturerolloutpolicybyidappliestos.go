package directoryfeaturerolloutpolicyappliesto

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

type ListDirectoryFeatureRolloutPolicyByIdAppliesTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListDirectoryFeatureRolloutPolicyByIdAppliesTosCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListDirectoryFeatureRolloutPolicyByIdAppliesTos ...
func (c DirectoryFeatureRolloutPolicyAppliesToClient) ListDirectoryFeatureRolloutPolicyByIdAppliesTos(ctx context.Context, id DirectoryFeatureRolloutPolicyId) (result ListDirectoryFeatureRolloutPolicyByIdAppliesTosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/appliesTo", id.ID()),
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
		Values *[]models.DirectoryObjectCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryFeatureRolloutPolicyByIdAppliesTosComplete retrieves all the results into a single object
func (c DirectoryFeatureRolloutPolicyAppliesToClient) ListDirectoryFeatureRolloutPolicyByIdAppliesTosComplete(ctx context.Context, id DirectoryFeatureRolloutPolicyId) (ListDirectoryFeatureRolloutPolicyByIdAppliesTosCompleteResult, error) {
	return c.ListDirectoryFeatureRolloutPolicyByIdAppliesTosCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListDirectoryFeatureRolloutPolicyByIdAppliesTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryFeatureRolloutPolicyAppliesToClient) ListDirectoryFeatureRolloutPolicyByIdAppliesTosCompleteMatchingPredicate(ctx context.Context, id DirectoryFeatureRolloutPolicyId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListDirectoryFeatureRolloutPolicyByIdAppliesTosCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListDirectoryFeatureRolloutPolicyByIdAppliesTos(ctx, id)
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

	result = ListDirectoryFeatureRolloutPolicyByIdAppliesTosCompleteResult{
		Items: items,
	}
	return
}
