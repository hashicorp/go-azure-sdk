package authenticationmicrosoftauthenticatormethod

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

type ListAuthenticationMicrosoftAuthenticatorMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MicrosoftAuthenticatorAuthenticationMethod
}

type ListAuthenticationMicrosoftAuthenticatorMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MicrosoftAuthenticatorAuthenticationMethod
}

type ListAuthenticationMicrosoftAuthenticatorMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationMicrosoftAuthenticatorMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationMicrosoftAuthenticatorMethods ...
func (c AuthenticationMicrosoftAuthenticatorMethodClient) ListAuthenticationMicrosoftAuthenticatorMethods(ctx context.Context, id UserId) (result ListAuthenticationMicrosoftAuthenticatorMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationMicrosoftAuthenticatorMethodsCustomPager{},
		Path:       fmt.Sprintf("%s/authentication/microsoftAuthenticatorMethods", id.ID()),
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
		Values *[]beta.MicrosoftAuthenticatorAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationMicrosoftAuthenticatorMethodsComplete retrieves all the results into a single object
func (c AuthenticationMicrosoftAuthenticatorMethodClient) ListAuthenticationMicrosoftAuthenticatorMethodsComplete(ctx context.Context, id UserId) (ListAuthenticationMicrosoftAuthenticatorMethodsCompleteResult, error) {
	return c.ListAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx, id, MicrosoftAuthenticatorAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationMicrosoftAuthenticatorMethodClient) ListAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate MicrosoftAuthenticatorAuthenticationMethodOperationPredicate) (result ListAuthenticationMicrosoftAuthenticatorMethodsCompleteResult, err error) {
	items := make([]beta.MicrosoftAuthenticatorAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationMicrosoftAuthenticatorMethods(ctx, id)
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

	result = ListAuthenticationMicrosoftAuthenticatorMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
