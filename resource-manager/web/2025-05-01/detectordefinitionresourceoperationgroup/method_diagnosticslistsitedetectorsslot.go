package detectordefinitionresourceoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsListSiteDetectorsSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DetectorDefinitionResource
}

type DiagnosticsListSiteDetectorsSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DetectorDefinitionResource
}

type DiagnosticsListSiteDetectorsSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DiagnosticsListSiteDetectorsSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DiagnosticsListSiteDetectorsSlot ...
func (c DetectorDefinitionResourceOperationGroupClient) DiagnosticsListSiteDetectorsSlot(ctx context.Context, id SlotDiagnosticId) (result DiagnosticsListSiteDetectorsSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DiagnosticsListSiteDetectorsSlotCustomPager{},
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
		Values *[]DetectorDefinitionResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DiagnosticsListSiteDetectorsSlotComplete retrieves all the results into a single object
func (c DetectorDefinitionResourceOperationGroupClient) DiagnosticsListSiteDetectorsSlotComplete(ctx context.Context, id SlotDiagnosticId) (DiagnosticsListSiteDetectorsSlotCompleteResult, error) {
	return c.DiagnosticsListSiteDetectorsSlotCompleteMatchingPredicate(ctx, id, DetectorDefinitionResourceOperationPredicate{})
}

// DiagnosticsListSiteDetectorsSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DetectorDefinitionResourceOperationGroupClient) DiagnosticsListSiteDetectorsSlotCompleteMatchingPredicate(ctx context.Context, id SlotDiagnosticId, predicate DetectorDefinitionResourceOperationPredicate) (result DiagnosticsListSiteDetectorsSlotCompleteResult, err error) {
	items := make([]DetectorDefinitionResource, 0)

	resp, err := c.DiagnosticsListSiteDetectorsSlot(ctx, id)
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

	result = DiagnosticsListSiteDetectorsSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
