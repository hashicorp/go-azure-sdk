package publiccertificateoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListPublicCertificatesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PublicCertificate
}

type WebAppsListPublicCertificatesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PublicCertificate
}

type WebAppsListPublicCertificatesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListPublicCertificatesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListPublicCertificatesSlot ...
func (c PublicCertificateOperationGroupClient) WebAppsListPublicCertificatesSlot(ctx context.Context, id SlotId) (result WebAppsListPublicCertificatesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListPublicCertificatesSlotCustomPager{},
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

// WebAppsListPublicCertificatesSlotComplete retrieves all the results into a single object
func (c PublicCertificateOperationGroupClient) WebAppsListPublicCertificatesSlotComplete(ctx context.Context, id SlotId) (WebAppsListPublicCertificatesSlotCompleteResult, error) {
	return c.WebAppsListPublicCertificatesSlotCompleteMatchingPredicate(ctx, id, PublicCertificateOperationPredicate{})
}

// WebAppsListPublicCertificatesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PublicCertificateOperationGroupClient) WebAppsListPublicCertificatesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate PublicCertificateOperationPredicate) (result WebAppsListPublicCertificatesSlotCompleteResult, err error) {
	items := make([]PublicCertificate, 0)

	resp, err := c.WebAppsListPublicCertificatesSlot(ctx, id)
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

	result = WebAppsListPublicCertificatesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
