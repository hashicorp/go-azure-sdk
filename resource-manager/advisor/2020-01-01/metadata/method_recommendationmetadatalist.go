package metadata

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationMetadataListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]MetadataEntity
}

type RecommendationMetadataListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []MetadataEntity
}

// RecommendationMetadataList ...
func (c MetadataClient) RecommendationMetadataList(ctx context.Context) (result RecommendationMetadataListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/providers/Microsoft.Advisor/metadata",
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
		Values *[]MetadataEntity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RecommendationMetadataListComplete retrieves all the results into a single object
func (c MetadataClient) RecommendationMetadataListComplete(ctx context.Context) (RecommendationMetadataListCompleteResult, error) {
	return c.RecommendationMetadataListCompleteMatchingPredicate(ctx, MetadataEntityOperationPredicate{})
}

// RecommendationMetadataListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MetadataClient) RecommendationMetadataListCompleteMatchingPredicate(ctx context.Context, predicate MetadataEntityOperationPredicate) (result RecommendationMetadataListCompleteResult, err error) {
	items := make([]MetadataEntity, 0)

	resp, err := c.RecommendationMetadataList(ctx)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = RecommendationMetadataListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
