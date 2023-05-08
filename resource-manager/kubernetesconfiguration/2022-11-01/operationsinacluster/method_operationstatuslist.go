package operationsinacluster

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatusListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]OperationStatusResult
}

type OperationStatusListCompleteResult struct {
	Items []OperationStatusResult
}

// OperationStatusList ...
func (c OperationsInAClusterClient) OperationStatusList(ctx context.Context, id ProviderId) (result OperationStatusListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.KubernetesConfiguration/operations", id.ID()),
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
		Values *[]OperationStatusResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// OperationStatusListComplete retrieves all the results into a single object
func (c OperationsInAClusterClient) OperationStatusListComplete(ctx context.Context, id ProviderId) (OperationStatusListCompleteResult, error) {
	return c.OperationStatusListCompleteMatchingPredicate(ctx, id, OperationStatusResultOperationPredicate{})
}

// OperationStatusListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OperationsInAClusterClient) OperationStatusListCompleteMatchingPredicate(ctx context.Context, id ProviderId, predicate OperationStatusResultOperationPredicate) (result OperationStatusListCompleteResult, err error) {
	items := make([]OperationStatusResult, 0)

	resp, err := c.OperationStatusList(ctx, id)
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

	result = OperationStatusListCompleteResult{
		Items: items,
	}
	return
}
