package batchmanagements

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchAccountListDetectorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DetectorResponse
}

type BatchAccountListDetectorsCompleteResult struct {
	Items []DetectorResponse
}

// BatchAccountListDetectors ...
func (c BatchManagementsClient) BatchAccountListDetectors(ctx context.Context, id BatchAccountId) (result BatchAccountListDetectorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/detectors", id.ID()),
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
		Values *[]DetectorResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BatchAccountListDetectorsComplete retrieves all the results into a single object
func (c BatchManagementsClient) BatchAccountListDetectorsComplete(ctx context.Context, id BatchAccountId) (BatchAccountListDetectorsCompleteResult, error) {
	return c.BatchAccountListDetectorsCompleteMatchingPredicate(ctx, id, DetectorResponseOperationPredicate{})
}

// BatchAccountListDetectorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BatchManagementsClient) BatchAccountListDetectorsCompleteMatchingPredicate(ctx context.Context, id BatchAccountId, predicate DetectorResponseOperationPredicate) (result BatchAccountListDetectorsCompleteResult, err error) {
	items := make([]DetectorResponse, 0)

	resp, err := c.BatchAccountListDetectors(ctx, id)
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

	result = BatchAccountListDetectorsCompleteResult{
		Items: items,
	}
	return
}
