package diagnostics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSiteDetectorResponsesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DetectorResponse
}

type ListSiteDetectorResponsesCompleteResult struct {
	Items []DetectorResponse
}

// ListSiteDetectorResponses ...
func (c DiagnosticsClient) ListSiteDetectorResponses(ctx context.Context, id SiteId) (result ListSiteDetectorResponsesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
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

// ListSiteDetectorResponsesComplete retrieves all the results into a single object
func (c DiagnosticsClient) ListSiteDetectorResponsesComplete(ctx context.Context, id SiteId) (ListSiteDetectorResponsesCompleteResult, error) {
	return c.ListSiteDetectorResponsesCompleteMatchingPredicate(ctx, id, DetectorResponseOperationPredicate{})
}

// ListSiteDetectorResponsesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticsClient) ListSiteDetectorResponsesCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate DetectorResponseOperationPredicate) (result ListSiteDetectorResponsesCompleteResult, err error) {
	items := make([]DetectorResponse, 0)

	resp, err := c.ListSiteDetectorResponses(ctx, id)
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

	result = ListSiteDetectorResponsesCompleteResult{
		Items: items,
	}
	return
}
