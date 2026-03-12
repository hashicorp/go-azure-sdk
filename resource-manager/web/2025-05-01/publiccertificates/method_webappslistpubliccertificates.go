package publiccertificates

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

type WebAppsListPublicCertificatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PublicCertificate
}

type WebAppsListPublicCertificatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PublicCertificate
}

type WebAppsListPublicCertificatesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListPublicCertificatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListPublicCertificates ...
func (c PublicCertificatesClient) WebAppsListPublicCertificates(ctx context.Context, id commonids.AppServiceId) (result WebAppsListPublicCertificatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListPublicCertificatesCustomPager{},
		Path:       fmt.Sprintf("%s/publicCertificates", id.ID()),
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
		Values *[]PublicCertificate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListPublicCertificatesComplete retrieves all the results into a single object
func (c PublicCertificatesClient) WebAppsListPublicCertificatesComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListPublicCertificatesCompleteResult, error) {
	return c.WebAppsListPublicCertificatesCompleteMatchingPredicate(ctx, id, PublicCertificateOperationPredicate{})
}

// WebAppsListPublicCertificatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PublicCertificatesClient) WebAppsListPublicCertificatesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate PublicCertificateOperationPredicate) (result WebAppsListPublicCertificatesCompleteResult, err error) {
	items := make([]PublicCertificate, 0)

	resp, err := c.WebAppsListPublicCertificates(ctx, id)
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

	result = WebAppsListPublicCertificatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
