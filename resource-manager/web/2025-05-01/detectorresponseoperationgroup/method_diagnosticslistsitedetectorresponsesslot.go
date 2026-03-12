package detectorresponseoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsListSiteDetectorResponsesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DetectorResponse
}

type DiagnosticsListSiteDetectorResponsesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DetectorResponse
}

type DiagnosticsListSiteDetectorResponsesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DiagnosticsListSiteDetectorResponsesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DiagnosticsListSiteDetectorResponsesSlot ...
func (c DetectorResponseOperationGroupClient) DiagnosticsListSiteDetectorResponsesSlot(ctx context.Context, id SlotId) (result DiagnosticsListSiteDetectorResponsesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DiagnosticsListSiteDetectorResponsesSlotCustomPager{},
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

// DiagnosticsListSiteDetectorResponsesSlotComplete retrieves all the results into a single object
func (c DetectorResponseOperationGroupClient) DiagnosticsListSiteDetectorResponsesSlotComplete(ctx context.Context, id SlotId) (DiagnosticsListSiteDetectorResponsesSlotCompleteResult, error) {
	return c.DiagnosticsListSiteDetectorResponsesSlotCompleteMatchingPredicate(ctx, id, DetectorResponseOperationPredicate{})
}

// DiagnosticsListSiteDetectorResponsesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DetectorResponseOperationGroupClient) DiagnosticsListSiteDetectorResponsesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate DetectorResponseOperationPredicate) (result DiagnosticsListSiteDetectorResponsesSlotCompleteResult, err error) {
	items := make([]DetectorResponse, 0)

	resp, err := c.DiagnosticsListSiteDetectorResponsesSlot(ctx, id)
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

	result = DiagnosticsListSiteDetectorResponsesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
