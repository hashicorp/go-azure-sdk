package sitecertificates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Certificate
}

type ListSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Certificate
}

type ListSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSlot ...
func (c SiteCertificatesClient) ListSlot(ctx context.Context, id SlotId) (result ListSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSlotCustomPager{},
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
		Values *[]Certificate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSlotComplete retrieves all the results into a single object
func (c SiteCertificatesClient) ListSlotComplete(ctx context.Context, id SlotId) (ListSlotCompleteResult, error) {
	return c.ListSlotCompleteMatchingPredicate(ctx, id, CertificateOperationPredicate{})
}

// ListSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteCertificatesClient) ListSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate CertificateOperationPredicate) (result ListSlotCompleteResult, err error) {
	items := make([]Certificate, 0)

	resp, err := c.ListSlot(ctx, id)
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

	result = ListSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
