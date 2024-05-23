package integrationruntimeobjectmetadata

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SsisObjectMetadata
}

type GetCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SsisObjectMetadata
}

// Get ...
func (c IntegrationRuntimeObjectMetadataClient) Get(ctx context.Context, id IntegrationRuntimeId, input GetSsisObjectMetadataRequest) (result GetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/getObjectMetadata", id.ID()),
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
		Values *[]SsisObjectMetadata `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetComplete retrieves all the results into a single object
func (c IntegrationRuntimeObjectMetadataClient) GetComplete(ctx context.Context, id IntegrationRuntimeId, input GetSsisObjectMetadataRequest) (GetCompleteResult, error) {
	return c.GetCompleteMatchingPredicate(ctx, id, input, SsisObjectMetadataOperationPredicate{})
}

// GetCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntegrationRuntimeObjectMetadataClient) GetCompleteMatchingPredicate(ctx context.Context, id IntegrationRuntimeId, input GetSsisObjectMetadataRequest, predicate SsisObjectMetadataOperationPredicate) (result GetCompleteResult, err error) {
	items := make([]SsisObjectMetadata, 0)

	resp, err := c.Get(ctx, id, input)
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

	result = GetCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
