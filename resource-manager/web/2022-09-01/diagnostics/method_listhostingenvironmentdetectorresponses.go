package diagnostics

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

type ListHostingEnvironmentDetectorResponsesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DetectorResponse
}

type ListHostingEnvironmentDetectorResponsesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DetectorResponse
}

type ListHostingEnvironmentDetectorResponsesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListHostingEnvironmentDetectorResponsesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListHostingEnvironmentDetectorResponses ...
func (c DiagnosticsClient) ListHostingEnvironmentDetectorResponses(ctx context.Context, id commonids.AppServiceEnvironmentId) (result ListHostingEnvironmentDetectorResponsesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListHostingEnvironmentDetectorResponsesCustomPager{},
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

// ListHostingEnvironmentDetectorResponsesComplete retrieves all the results into a single object
func (c DiagnosticsClient) ListHostingEnvironmentDetectorResponsesComplete(ctx context.Context, id commonids.AppServiceEnvironmentId) (ListHostingEnvironmentDetectorResponsesCompleteResult, error) {
	return c.ListHostingEnvironmentDetectorResponsesCompleteMatchingPredicate(ctx, id, DetectorResponseOperationPredicate{})
}

// ListHostingEnvironmentDetectorResponsesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticsClient) ListHostingEnvironmentDetectorResponsesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, predicate DetectorResponseOperationPredicate) (result ListHostingEnvironmentDetectorResponsesCompleteResult, err error) {
	items := make([]DetectorResponse, 0)

	resp, err := c.ListHostingEnvironmentDetectorResponses(ctx, id)
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

	result = ListHostingEnvironmentDetectorResponsesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
