package certificateconnectordetail

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCertificateConnectorDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CertificateConnectorDetails
}

type ListCertificateConnectorDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CertificateConnectorDetails
}

type ListCertificateConnectorDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCertificateConnectorDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCertificateConnectorDetails ...
func (c CertificateConnectorDetailClient) ListCertificateConnectorDetails(ctx context.Context) (result ListCertificateConnectorDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCertificateConnectorDetailsCustomPager{},
		Path:       "/deviceManagement/certificateConnectorDetails",
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
		Values *[]beta.CertificateConnectorDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCertificateConnectorDetailsComplete retrieves all the results into a single object
func (c CertificateConnectorDetailClient) ListCertificateConnectorDetailsComplete(ctx context.Context) (ListCertificateConnectorDetailsCompleteResult, error) {
	return c.ListCertificateConnectorDetailsCompleteMatchingPredicate(ctx, CertificateConnectorDetailsOperationPredicate{})
}

// ListCertificateConnectorDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CertificateConnectorDetailClient) ListCertificateConnectorDetailsCompleteMatchingPredicate(ctx context.Context, predicate CertificateConnectorDetailsOperationPredicate) (result ListCertificateConnectorDetailsCompleteResult, err error) {
	items := make([]beta.CertificateConnectorDetails, 0)

	resp, err := c.ListCertificateConnectorDetails(ctx)
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

	result = ListCertificateConnectorDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
