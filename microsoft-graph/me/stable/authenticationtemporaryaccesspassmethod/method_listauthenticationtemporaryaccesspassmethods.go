package authenticationtemporaryaccesspassmethod

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

type ListAuthenticationTemporaryAccessPassMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TemporaryAccessPassAuthenticationMethod
}

type ListAuthenticationTemporaryAccessPassMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TemporaryAccessPassAuthenticationMethod
}

type ListAuthenticationTemporaryAccessPassMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationTemporaryAccessPassMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationTemporaryAccessPassMethods ...
func (c AuthenticationTemporaryAccessPassMethodClient) ListAuthenticationTemporaryAccessPassMethods(ctx context.Context) (result ListAuthenticationTemporaryAccessPassMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationTemporaryAccessPassMethodsCustomPager{},
		Path:       "/me/authentication/temporaryAccessPassMethods",
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
		Values *[]stable.TemporaryAccessPassAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationTemporaryAccessPassMethodsComplete retrieves all the results into a single object
func (c AuthenticationTemporaryAccessPassMethodClient) ListAuthenticationTemporaryAccessPassMethodsComplete(ctx context.Context) (ListAuthenticationTemporaryAccessPassMethodsCompleteResult, error) {
	return c.ListAuthenticationTemporaryAccessPassMethodsCompleteMatchingPredicate(ctx, TemporaryAccessPassAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationTemporaryAccessPassMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationTemporaryAccessPassMethodClient) ListAuthenticationTemporaryAccessPassMethodsCompleteMatchingPredicate(ctx context.Context, predicate TemporaryAccessPassAuthenticationMethodOperationPredicate) (result ListAuthenticationTemporaryAccessPassMethodsCompleteResult, err error) {
	items := make([]stable.TemporaryAccessPassAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationTemporaryAccessPassMethods(ctx)
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

	result = ListAuthenticationTemporaryAccessPassMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
