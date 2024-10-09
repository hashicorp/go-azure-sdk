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

type SqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SqlPoolBlobAuditingPolicy
}

type SqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SqlPoolBlobAuditingPolicy
}

type SqlPoolBlobAuditingPoliciesListBySqlPoolCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SqlPoolBlobAuditingPoliciesListBySqlPoolCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SqlPoolBlobAuditingPoliciesListBySqlPool ...
func (c SqlPoolsBlobAuditingClient) SqlPoolBlobAuditingPoliciesListBySqlPool(ctx context.Context, id SqlPoolId) (result SqlPoolBlobAuditingPoliciesListBySqlPoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SqlPoolBlobAuditingPoliciesListBySqlPoolCustomPager{},
		Path:       fmt.Sprintf("%s/auditingSettings", id.ID()),
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
		Values *[]SqlPoolBlobAuditingPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlPoolBlobAuditingPoliciesListBySqlPoolComplete retrieves all the results into a single object
func (c SqlPoolsBlobAuditingClient) SqlPoolBlobAuditingPoliciesListBySqlPoolComplete(ctx context.Context, id SqlPoolId) (SqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult, error) {
	return c.SqlPoolBlobAuditingPoliciesListBySqlPoolCompleteMatchingPredicate(ctx, id, SqlPoolBlobAuditingPolicyOperationPredicate{})
}

// SqlPoolBlobAuditingPoliciesListBySqlPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsBlobAuditingClient) SqlPoolBlobAuditingPoliciesListBySqlPoolCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate SqlPoolBlobAuditingPolicyOperationPredicate) (result SqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult, err error) {
	items := make([]SqlPoolBlobAuditingPolicy, 0)

	resp, err := c.SqlPoolBlobAuditingPoliciesListBySqlPool(ctx, id)
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

	result = SqlPoolBlobAuditingPoliciesListBySqlPoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
