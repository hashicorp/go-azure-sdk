package certificateordersdiagnostics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAppServiceCertificateOrderDetectorResponseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DetectorResponse
}

type ListAppServiceCertificateOrderDetectorResponseCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DetectorResponse
}

type ListAppServiceCertificateOrderDetectorResponseCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListAppServiceCertificateOrderDetectorResponseCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppServiceCertificateOrderDetectorResponse ...
func (c CertificateOrdersDiagnosticsClient) ListAppServiceCertificateOrderDetectorResponse(ctx context.Context, id CertificateOrderId) (result ListAppServiceCertificateOrderDetectorResponseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppServiceCertificateOrderDetectorResponseCustomPager{},
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

// ListAppServiceCertificateOrderDetectorResponseComplete retrieves all the results into a single object
func (c CertificateOrdersDiagnosticsClient) ListAppServiceCertificateOrderDetectorResponseComplete(ctx context.Context, id CertificateOrderId) (ListAppServiceCertificateOrderDetectorResponseCompleteResult, error) {
	return c.ListAppServiceCertificateOrderDetectorResponseCompleteMatchingPredicate(ctx, id, DetectorResponseOperationPredicate{})
}

// ListAppServiceCertificateOrderDetectorResponseCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CertificateOrdersDiagnosticsClient) ListAppServiceCertificateOrderDetectorResponseCompleteMatchingPredicate(ctx context.Context, id CertificateOrderId, predicate DetectorResponseOperationPredicate) (result ListAppServiceCertificateOrderDetectorResponseCompleteResult, err error) {
	items := make([]DetectorResponse, 0)

	resp, err := c.ListAppServiceCertificateOrderDetectorResponse(ctx, id)
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

	result = ListAppServiceCertificateOrderDetectorResponseCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
