package hypervsites

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteHealthSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SiteHealthSummary
}

type GetSiteHealthSummaryCompleteResult struct {
	Items []SiteHealthSummary
}

// GetSiteHealthSummary ...
func (c HyperVSitesClient) GetSiteHealthSummary(ctx context.Context, id HyperVSiteId) (result GetSiteHealthSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/healthSummary", id.ID()),
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
		Values *[]SiteHealthSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetSiteHealthSummaryComplete retrieves all the results into a single object
func (c HyperVSitesClient) GetSiteHealthSummaryComplete(ctx context.Context, id HyperVSiteId) (GetSiteHealthSummaryCompleteResult, error) {
	return c.GetSiteHealthSummaryCompleteMatchingPredicate(ctx, id, SiteHealthSummaryOperationPredicate{})
}

// GetSiteHealthSummaryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HyperVSitesClient) GetSiteHealthSummaryCompleteMatchingPredicate(ctx context.Context, id HyperVSiteId, predicate SiteHealthSummaryOperationPredicate) (result GetSiteHealthSummaryCompleteResult, err error) {
	items := make([]SiteHealthSummary, 0)

	resp, err := c.GetSiteHealthSummary(ctx, id)
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

	result = GetSiteHealthSummaryCompleteResult{
		Items: items,
	}
	return
}
