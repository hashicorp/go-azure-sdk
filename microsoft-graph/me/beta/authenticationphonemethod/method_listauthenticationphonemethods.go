package authenticationphonemethod

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

type ListAuthenticationPhoneMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PhoneAuthenticationMethod
}

type ListAuthenticationPhoneMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PhoneAuthenticationMethod
}

type ListAuthenticationPhoneMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationPhoneMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationPhoneMethods ...
func (c AuthenticationPhoneMethodClient) ListAuthenticationPhoneMethods(ctx context.Context) (result ListAuthenticationPhoneMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationPhoneMethodsCustomPager{},
		Path:       "/me/authentication/phoneMethods",
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
		Values *[]beta.PhoneAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationPhoneMethodsComplete retrieves all the results into a single object
func (c AuthenticationPhoneMethodClient) ListAuthenticationPhoneMethodsComplete(ctx context.Context) (ListAuthenticationPhoneMethodsCompleteResult, error) {
	return c.ListAuthenticationPhoneMethodsCompleteMatchingPredicate(ctx, PhoneAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationPhoneMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationPhoneMethodClient) ListAuthenticationPhoneMethodsCompleteMatchingPredicate(ctx context.Context, predicate PhoneAuthenticationMethodOperationPredicate) (result ListAuthenticationPhoneMethodsCompleteResult, err error) {
	items := make([]beta.PhoneAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationPhoneMethods(ctx)
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

	result = ListAuthenticationPhoneMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
