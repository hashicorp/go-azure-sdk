package authenticationpasswordlessmicrosoftauthenticatormethod

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

type ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PasswordlessMicrosoftAuthenticatorAuthenticationMethod
}

type ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PasswordlessMicrosoftAuthenticatorAuthenticationMethod
}

type ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationPasswordlessMicrosoftAuthenticatorMethods ...
func (c AuthenticationPasswordlessMicrosoftAuthenticatorMethodClient) ListAuthenticationPasswordlessMicrosoftAuthenticatorMethods(ctx context.Context, id UserId) (result ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCustomPager{},
		Path:       fmt.Sprintf("%s/authentication/passwordlessMicrosoftAuthenticatorMethods", id.ID()),
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
		Values *[]beta.PasswordlessMicrosoftAuthenticatorAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsComplete retrieves all the results into a single object
func (c AuthenticationPasswordlessMicrosoftAuthenticatorMethodClient) ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsComplete(ctx context.Context, id UserId) (ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteResult, error) {
	return c.ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx, id, PasswordlessMicrosoftAuthenticatorAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationPasswordlessMicrosoftAuthenticatorMethodClient) ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate PasswordlessMicrosoftAuthenticatorAuthenticationMethodOperationPredicate) (result ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteResult, err error) {
	items := make([]beta.PasswordlessMicrosoftAuthenticatorAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationPasswordlessMicrosoftAuthenticatorMethods(ctx, id)
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

	result = ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
