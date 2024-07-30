package appconsentappconsentrequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAppConsentAppConsentRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AppConsentRequest
}

type ListAppConsentAppConsentRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AppConsentRequest
}

type ListAppConsentAppConsentRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppConsentAppConsentRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppConsentAppConsentRequests ...
func (c AppConsentAppConsentRequestClient) ListAppConsentAppConsentRequests(ctx context.Context) (result ListAppConsentAppConsentRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppConsentAppConsentRequestsCustomPager{},
		Path:       "/identityGovernance/appConsent/appConsentRequests",
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
		Values *[]stable.AppConsentRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppConsentAppConsentRequestsComplete retrieves all the results into a single object
func (c AppConsentAppConsentRequestClient) ListAppConsentAppConsentRequestsComplete(ctx context.Context) (ListAppConsentAppConsentRequestsCompleteResult, error) {
	return c.ListAppConsentAppConsentRequestsCompleteMatchingPredicate(ctx, AppConsentRequestOperationPredicate{})
}

// ListAppConsentAppConsentRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppConsentAppConsentRequestClient) ListAppConsentAppConsentRequestsCompleteMatchingPredicate(ctx context.Context, predicate AppConsentRequestOperationPredicate) (result ListAppConsentAppConsentRequestsCompleteResult, err error) {
	items := make([]stable.AppConsentRequest, 0)

	resp, err := c.ListAppConsentAppConsentRequests(ctx)
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

	result = ListAppConsentAppConsentRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
