package quotausagesforflexibleservers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaUsagesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]QuotaUsage
}

type QuotaUsagesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []QuotaUsage
}

// QuotaUsagesList ...
func (c QuotaUsagesForFlexibleServersClient) QuotaUsagesList(ctx context.Context, id LocationId) (result QuotaUsagesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/resourceType/flexibleServers/usages", id.ID()),
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
		Values *[]QuotaUsage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// QuotaUsagesListComplete retrieves all the results into a single object
func (c QuotaUsagesForFlexibleServersClient) QuotaUsagesListComplete(ctx context.Context, id LocationId) (QuotaUsagesListCompleteResult, error) {
	return c.QuotaUsagesListCompleteMatchingPredicate(ctx, id, QuotaUsageOperationPredicate{})
}

// QuotaUsagesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c QuotaUsagesForFlexibleServersClient) QuotaUsagesListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate QuotaUsageOperationPredicate) (result QuotaUsagesListCompleteResult, err error) {
	items := make([]QuotaUsage, 0)

	resp, err := c.QuotaUsagesList(ctx, id)
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

	result = QuotaUsagesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
