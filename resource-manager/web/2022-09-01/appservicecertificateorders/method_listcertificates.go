package appservicecertificateorders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCertificatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AppServiceCertificateResource
}

type ListCertificatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AppServiceCertificateResource
}

// ListCertificates ...
func (c AppServiceCertificateOrdersClient) ListCertificates(ctx context.Context, id CertificateOrderId) (result ListCertificatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/certificates", id.ID()),
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
		Values *[]AppServiceCertificateResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCertificatesComplete retrieves all the results into a single object
func (c AppServiceCertificateOrdersClient) ListCertificatesComplete(ctx context.Context, id CertificateOrderId) (ListCertificatesCompleteResult, error) {
	return c.ListCertificatesCompleteMatchingPredicate(ctx, id, AppServiceCertificateResourceOperationPredicate{})
}

// ListCertificatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceCertificateOrdersClient) ListCertificatesCompleteMatchingPredicate(ctx context.Context, id CertificateOrderId, predicate AppServiceCertificateResourceOperationPredicate) (result ListCertificatesCompleteResult, err error) {
	items := make([]AppServiceCertificateResource, 0)

	resp, err := c.ListCertificates(ctx, id)
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

	result = ListCertificatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
