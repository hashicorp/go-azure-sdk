package detectorresponses

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsListHostingEnvironmentDetectorResponsesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DetectorResponse
}

type DiagnosticsListHostingEnvironmentDetectorResponsesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DetectorResponse
}

type DiagnosticsListHostingEnvironmentDetectorResponsesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DiagnosticsListHostingEnvironmentDetectorResponsesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DiagnosticsListHostingEnvironmentDetectorResponses ...
func (c DetectorResponsesClient) DiagnosticsListHostingEnvironmentDetectorResponses(ctx context.Context, id commonids.AppServiceEnvironmentId) (result DiagnosticsListHostingEnvironmentDetectorResponsesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DiagnosticsListHostingEnvironmentDetectorResponsesCustomPager{},
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

// DiagnosticsListHostingEnvironmentDetectorResponsesComplete retrieves all the results into a single object
func (c DetectorResponsesClient) DiagnosticsListHostingEnvironmentDetectorResponsesComplete(ctx context.Context, id commonids.AppServiceEnvironmentId) (DiagnosticsListHostingEnvironmentDetectorResponsesCompleteResult, error) {
	return c.DiagnosticsListHostingEnvironmentDetectorResponsesCompleteMatchingPredicate(ctx, id, DetectorResponseOperationPredicate{})
}

// DiagnosticsListHostingEnvironmentDetectorResponsesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DetectorResponsesClient) DiagnosticsListHostingEnvironmentDetectorResponsesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, predicate DetectorResponseOperationPredicate) (result DiagnosticsListHostingEnvironmentDetectorResponsesCompleteResult, err error) {
	items := make([]DetectorResponse, 0)

	resp, err := c.DiagnosticsListHostingEnvironmentDetectorResponses(ctx, id)
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

	result = DiagnosticsListHostingEnvironmentDetectorResponsesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
