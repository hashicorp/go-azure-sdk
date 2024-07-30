package authenticationmethod

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

type ListAuthenticationMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AuthenticationMethod
}

type ListAuthenticationMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AuthenticationMethod
}

type ListAuthenticationMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationMethods ...
func (c AuthenticationMethodClient) ListAuthenticationMethods(ctx context.Context, id UserId) (result ListAuthenticationMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationMethodsCustomPager{},
		Path:       fmt.Sprintf("%s/authentication/methods", id.ID()),
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
		Values *[]beta.AuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationMethodsComplete retrieves all the results into a single object
func (c AuthenticationMethodClient) ListAuthenticationMethodsComplete(ctx context.Context, id UserId) (ListAuthenticationMethodsCompleteResult, error) {
	return c.ListAuthenticationMethodsCompleteMatchingPredicate(ctx, id, AuthenticationMethodOperationPredicate{})
}

// ListAuthenticationMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationMethodClient) ListAuthenticationMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate AuthenticationMethodOperationPredicate) (result ListAuthenticationMethodsCompleteResult, err error) {
	items := make([]beta.AuthenticationMethod, 0)

	resp, err := c.ListAuthenticationMethods(ctx, id)
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

	result = ListAuthenticationMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
