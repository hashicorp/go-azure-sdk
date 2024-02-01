package sqlpoolsblobauditing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ExtendedSqlPoolBlobAuditingPolicy
}

type ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ExtendedSqlPoolBlobAuditingPolicy
}

// ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool ...
func (c SqlPoolsBlobAuditingClient) ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool(ctx context.Context, id SqlPoolId) (result ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/extendedAuditingSettings", id.ID()),
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
		Values *[]ExtendedSqlPoolBlobAuditingPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolComplete retrieves all the results into a single object
func (c SqlPoolsBlobAuditingClient) ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolComplete(ctx context.Context, id SqlPoolId) (ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult, error) {
	return c.ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteMatchingPredicate(ctx, id, ExtendedSqlPoolBlobAuditingPolicyOperationPredicate{})
}

// ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsBlobAuditingClient) ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate ExtendedSqlPoolBlobAuditingPolicyOperationPredicate) (result ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult, err error) {
	items := make([]ExtendedSqlPoolBlobAuditingPolicy, 0)

	resp, err := c.ExtendedSqlPoolBlobAuditingPoliciesListBySqlPool(ctx, id)
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

	result = ExtendedSqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
