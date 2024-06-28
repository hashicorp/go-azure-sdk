package transfers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerTransfersListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TransferDetails
}

type PartnerTransfersListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TransferDetails
}

type PartnerTransfersListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PartnerTransfersListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PartnerTransfersList ...
func (c TransfersClient) PartnerTransfersList(ctx context.Context, id BillingProfileCustomerId) (result PartnerTransfersListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &PartnerTransfersListCustomPager{},
		Path:       fmt.Sprintf("%s/transfers", id.ID()),
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
		Values *[]TransferDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PartnerTransfersListComplete retrieves all the results into a single object
func (c TransfersClient) PartnerTransfersListComplete(ctx context.Context, id BillingProfileCustomerId) (PartnerTransfersListCompleteResult, error) {
	return c.PartnerTransfersListCompleteMatchingPredicate(ctx, id, TransferDetailsOperationPredicate{})
}

// PartnerTransfersListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TransfersClient) PartnerTransfersListCompleteMatchingPredicate(ctx context.Context, id BillingProfileCustomerId, predicate TransferDetailsOperationPredicate) (result PartnerTransfersListCompleteResult, err error) {
	items := make([]TransferDetails, 0)

	resp, err := c.PartnerTransfersList(ctx, id)
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

	result = PartnerTransfersListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
