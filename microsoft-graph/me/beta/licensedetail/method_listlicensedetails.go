package licensedetail

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

type ListLicenseDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.LicenseDetails
}

type ListLicenseDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.LicenseDetails
}

type ListLicenseDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLicenseDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLicenseDetails ...
func (c LicenseDetailClient) ListLicenseDetails(ctx context.Context) (result ListLicenseDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListLicenseDetailsCustomPager{},
		Path:       "/me/licenseDetails",
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
		Values *[]beta.LicenseDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLicenseDetailsComplete retrieves all the results into a single object
func (c LicenseDetailClient) ListLicenseDetailsComplete(ctx context.Context) (ListLicenseDetailsCompleteResult, error) {
	return c.ListLicenseDetailsCompleteMatchingPredicate(ctx, LicenseDetailsOperationPredicate{})
}

// ListLicenseDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LicenseDetailClient) ListLicenseDetailsCompleteMatchingPredicate(ctx context.Context, predicate LicenseDetailsOperationPredicate) (result ListLicenseDetailsCompleteResult, err error) {
	items := make([]beta.LicenseDetails, 0)

	resp, err := c.ListLicenseDetails(ctx)
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

	result = ListLicenseDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
