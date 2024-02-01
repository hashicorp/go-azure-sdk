package sqlpoolsusages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolUsagesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SqlPoolUsage
}

type SqlPoolUsagesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SqlPoolUsage
}

// SqlPoolUsagesList ...
func (c SqlPoolsUsagesClient) SqlPoolUsagesList(ctx context.Context, id SqlPoolId) (result SqlPoolUsagesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/usages", id.ID()),
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
		Values *[]SqlPoolUsage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlPoolUsagesListComplete retrieves all the results into a single object
func (c SqlPoolsUsagesClient) SqlPoolUsagesListComplete(ctx context.Context, id SqlPoolId) (SqlPoolUsagesListCompleteResult, error) {
	return c.SqlPoolUsagesListCompleteMatchingPredicate(ctx, id, SqlPoolUsageOperationPredicate{})
}

// SqlPoolUsagesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsUsagesClient) SqlPoolUsagesListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate SqlPoolUsageOperationPredicate) (result SqlPoolUsagesListCompleteResult, err error) {
	items := make([]SqlPoolUsage, 0)

	resp, err := c.SqlPoolUsagesList(ctx, id)
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

	result = SqlPoolUsagesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
