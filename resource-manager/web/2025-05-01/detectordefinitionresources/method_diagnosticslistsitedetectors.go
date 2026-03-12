package detectordefinitionresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsListSiteDetectorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DetectorDefinitionResource
}

type DiagnosticsListSiteDetectorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DetectorDefinitionResource
}

type DiagnosticsListSiteDetectorsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DiagnosticsListSiteDetectorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DiagnosticsListSiteDetectors ...
func (c DetectorDefinitionResourcesClient) DiagnosticsListSiteDetectors(ctx context.Context, id DiagnosticId) (result DiagnosticsListSiteDetectorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DiagnosticsListSiteDetectorsCustomPager{},
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

// DiagnosticsListSiteDetectorsComplete retrieves all the results into a single object
func (c DetectorDefinitionResourcesClient) DiagnosticsListSiteDetectorsComplete(ctx context.Context, id DiagnosticId) (DiagnosticsListSiteDetectorsCompleteResult, error) {
	return c.DiagnosticsListSiteDetectorsCompleteMatchingPredicate(ctx, id, DetectorDefinitionResourceOperationPredicate{})
}

// DiagnosticsListSiteDetectorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DetectorDefinitionResourcesClient) DiagnosticsListSiteDetectorsCompleteMatchingPredicate(ctx context.Context, id DiagnosticId, predicate DetectorDefinitionResourceOperationPredicate) (result DiagnosticsListSiteDetectorsCompleteResult, err error) {
	items := make([]DetectorDefinitionResource, 0)

	resp, err := c.DiagnosticsListSiteDetectors(ctx, id)
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

	result = DiagnosticsListSiteDetectorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
