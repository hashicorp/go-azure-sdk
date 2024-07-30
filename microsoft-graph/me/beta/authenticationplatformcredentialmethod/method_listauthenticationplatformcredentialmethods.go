package authenticationplatformcredentialmethod

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

type ListAuthenticationPlatformCredentialMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PlatformCredentialAuthenticationMethod
}

type ListAuthenticationPlatformCredentialMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PlatformCredentialAuthenticationMethod
}

type ListAuthenticationPlatformCredentialMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationPlatformCredentialMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationPlatformCredentialMethods ...
func (c AuthenticationPlatformCredentialMethodClient) ListAuthenticationPlatformCredentialMethods(ctx context.Context) (result ListAuthenticationPlatformCredentialMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationPlatformCredentialMethodsCustomPager{},
		Path:       "/me/authentication/platformCredentialMethods",
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
		Values *[]beta.PlatformCredentialAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationPlatformCredentialMethodsComplete retrieves all the results into a single object
func (c AuthenticationPlatformCredentialMethodClient) ListAuthenticationPlatformCredentialMethodsComplete(ctx context.Context) (ListAuthenticationPlatformCredentialMethodsCompleteResult, error) {
	return c.ListAuthenticationPlatformCredentialMethodsCompleteMatchingPredicate(ctx, PlatformCredentialAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationPlatformCredentialMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationPlatformCredentialMethodClient) ListAuthenticationPlatformCredentialMethodsCompleteMatchingPredicate(ctx context.Context, predicate PlatformCredentialAuthenticationMethodOperationPredicate) (result ListAuthenticationPlatformCredentialMethodsCompleteResult, err error) {
	items := make([]beta.PlatformCredentialAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationPlatformCredentialMethods(ctx)
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

	result = ListAuthenticationPlatformCredentialMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
