package usercredentialusagedetail

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

type ListUserCredentialUsageDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserCredentialUsageDetails
}

type ListUserCredentialUsageDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserCredentialUsageDetails
}

type ListUserCredentialUsageDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserCredentialUsageDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserCredentialUsageDetails ...
func (c UserCredentialUsageDetailClient) ListUserCredentialUsageDetails(ctx context.Context) (result ListUserCredentialUsageDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserCredentialUsageDetailsCustomPager{},
		Path:       "/reports/userCredentialUsageDetails",
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
		Values *[]beta.UserCredentialUsageDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserCredentialUsageDetailsComplete retrieves all the results into a single object
func (c UserCredentialUsageDetailClient) ListUserCredentialUsageDetailsComplete(ctx context.Context) (ListUserCredentialUsageDetailsCompleteResult, error) {
	return c.ListUserCredentialUsageDetailsCompleteMatchingPredicate(ctx, UserCredentialUsageDetailsOperationPredicate{})
}

// ListUserCredentialUsageDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCredentialUsageDetailClient) ListUserCredentialUsageDetailsCompleteMatchingPredicate(ctx context.Context, predicate UserCredentialUsageDetailsOperationPredicate) (result ListUserCredentialUsageDetailsCompleteResult, err error) {
	items := make([]beta.UserCredentialUsageDetails, 0)

	resp, err := c.ListUserCredentialUsageDetails(ctx)
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

	result = ListUserCredentialUsageDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
