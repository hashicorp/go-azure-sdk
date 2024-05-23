package changedatacapture

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByFactoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ChangeDataCaptureResource
}

type ListByFactoryCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ChangeDataCaptureResource
}

// ListByFactory ...
func (c ChangeDataCaptureClient) ListByFactory(ctx context.Context, id FactoryId) (result ListByFactoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/adfcdcs", id.ID()),
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
		Values *[]ChangeDataCaptureResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByFactoryComplete retrieves all the results into a single object
func (c ChangeDataCaptureClient) ListByFactoryComplete(ctx context.Context, id FactoryId) (ListByFactoryCompleteResult, error) {
	return c.ListByFactoryCompleteMatchingPredicate(ctx, id, ChangeDataCaptureResourceOperationPredicate{})
}

// ListByFactoryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChangeDataCaptureClient) ListByFactoryCompleteMatchingPredicate(ctx context.Context, id FactoryId, predicate ChangeDataCaptureResourceOperationPredicate) (result ListByFactoryCompleteResult, err error) {
	items := make([]ChangeDataCaptureResource, 0)

	resp, err := c.ListByFactory(ctx, id)
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

	result = ListByFactoryCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
