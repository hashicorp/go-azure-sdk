package alerts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListExternalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Alert
}

type ListExternalCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Alert
}

// ListExternal ...
func (c AlertsClient) ListExternal(ctx context.Context, id ExternalCloudProviderTypeId) (result ListExternalOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/alerts", id.ID()),
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
		Values *[]Alert `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExternalComplete retrieves all the results into a single object
func (c AlertsClient) ListExternalComplete(ctx context.Context, id ExternalCloudProviderTypeId) (ListExternalCompleteResult, error) {
	return c.ListExternalCompleteMatchingPredicate(ctx, id, AlertOperationPredicate{})
}

// ListExternalCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AlertsClient) ListExternalCompleteMatchingPredicate(ctx context.Context, id ExternalCloudProviderTypeId, predicate AlertOperationPredicate) (result ListExternalCompleteResult, err error) {
	items := make([]Alert, 0)

	resp, err := c.ListExternal(ctx, id)
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

	result = ListExternalCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
