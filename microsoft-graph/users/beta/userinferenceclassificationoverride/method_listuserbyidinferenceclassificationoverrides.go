package userinferenceclassificationoverride

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

type ListUserByIdInferenceClassificationOverridesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.InferenceClassificationOverrideCollectionResponse
}

type ListUserByIdInferenceClassificationOverridesCompleteResult struct {
	Items []models.InferenceClassificationOverrideCollectionResponse
}

// ListUserByIdInferenceClassificationOverrides ...
func (c UserInferenceClassificationOverrideClient) ListUserByIdInferenceClassificationOverrides(ctx context.Context, id UserId) (result ListUserByIdInferenceClassificationOverridesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/inferenceClassification/overrides", id.ID()),
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
		Values *[]models.InferenceClassificationOverrideCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdInferenceClassificationOverridesComplete retrieves all the results into a single object
func (c UserInferenceClassificationOverrideClient) ListUserByIdInferenceClassificationOverridesComplete(ctx context.Context, id UserId) (ListUserByIdInferenceClassificationOverridesCompleteResult, error) {
	return c.ListUserByIdInferenceClassificationOverridesCompleteMatchingPredicate(ctx, id, models.InferenceClassificationOverrideCollectionResponseOperationPredicate{})
}

// ListUserByIdInferenceClassificationOverridesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInferenceClassificationOverrideClient) ListUserByIdInferenceClassificationOverridesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.InferenceClassificationOverrideCollectionResponseOperationPredicate) (result ListUserByIdInferenceClassificationOverridesCompleteResult, err error) {
	items := make([]models.InferenceClassificationOverrideCollectionResponse, 0)

	resp, err := c.ListUserByIdInferenceClassificationOverrides(ctx, id)
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

	result = ListUserByIdInferenceClassificationOverridesCompleteResult{
		Items: items,
	}
	return
}
