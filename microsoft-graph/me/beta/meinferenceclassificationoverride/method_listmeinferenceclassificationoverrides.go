package meinferenceclassificationoverride

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

type ListMeInferenceClassificationOverridesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.InferenceClassificationOverrideCollectionResponse
}

type ListMeInferenceClassificationOverridesCompleteResult struct {
	Items []models.InferenceClassificationOverrideCollectionResponse
}

// ListMeInferenceClassificationOverrides ...
func (c MeInferenceClassificationOverrideClient) ListMeInferenceClassificationOverrides(ctx context.Context) (result ListMeInferenceClassificationOverridesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/inferenceClassification/overrides",
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

// ListMeInferenceClassificationOverridesComplete retrieves all the results into a single object
func (c MeInferenceClassificationOverrideClient) ListMeInferenceClassificationOverridesComplete(ctx context.Context) (ListMeInferenceClassificationOverridesCompleteResult, error) {
	return c.ListMeInferenceClassificationOverridesCompleteMatchingPredicate(ctx, models.InferenceClassificationOverrideCollectionResponseOperationPredicate{})
}

// ListMeInferenceClassificationOverridesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeInferenceClassificationOverrideClient) ListMeInferenceClassificationOverridesCompleteMatchingPredicate(ctx context.Context, predicate models.InferenceClassificationOverrideCollectionResponseOperationPredicate) (result ListMeInferenceClassificationOverridesCompleteResult, err error) {
	items := make([]models.InferenceClassificationOverrideCollectionResponse, 0)

	resp, err := c.ListMeInferenceClassificationOverrides(ctx)
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

	result = ListMeInferenceClassificationOverridesCompleteResult{
		Items: items,
	}
	return
}
