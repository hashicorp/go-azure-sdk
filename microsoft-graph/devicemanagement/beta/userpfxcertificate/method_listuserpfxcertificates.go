package userpfxcertificate

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

type ListUserPfxCertificatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserPFXCertificate
}

type ListUserPfxCertificatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserPFXCertificate
}

type ListUserPfxCertificatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserPfxCertificatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserPfxCertificates ...
func (c UserPfxCertificateClient) ListUserPfxCertificates(ctx context.Context) (result ListUserPfxCertificatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserPfxCertificatesCustomPager{},
		Path:       "/deviceManagement/userPfxCertificates",
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
		Values *[]beta.UserPFXCertificate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserPfxCertificatesComplete retrieves all the results into a single object
func (c UserPfxCertificateClient) ListUserPfxCertificatesComplete(ctx context.Context) (ListUserPfxCertificatesCompleteResult, error) {
	return c.ListUserPfxCertificatesCompleteMatchingPredicate(ctx, UserPFXCertificateOperationPredicate{})
}

// ListUserPfxCertificatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserPfxCertificateClient) ListUserPfxCertificatesCompleteMatchingPredicate(ctx context.Context, predicate UserPFXCertificateOperationPredicate) (result ListUserPfxCertificatesCompleteResult, err error) {
	items := make([]beta.UserPFXCertificate, 0)

	resp, err := c.ListUserPfxCertificates(ctx)
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

	result = ListUserPfxCertificatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
