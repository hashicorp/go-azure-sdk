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

type ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationOptions() ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationOptions {
	return ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationOptions{}
}

func (o ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListAuthenticationPasswordlessMicrosoftAuthenticatorMethods - Get passwordlessMicrosoftAuthenticatorMethods from
// users. Represents the Microsoft Authenticator Passwordless Phone Sign-in methods registered to a user for
// authentication.
func (c AuthenticationPasswordlessMicrosoftAuthenticatorMethodClient) ListAuthenticationPasswordlessMicrosoftAuthenticatorMethods(ctx context.Context, id beta.UserId, options ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationOptions) (result ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCustomPager{},
		Path:          fmt.Sprintf("%s/authentication/passwordlessMicrosoftAuthenticatorMethods", id.ID()),
		RetryFunc:     options.RetryFunc,
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
func (c AuthenticationPasswordlessMicrosoftAuthenticatorMethodClient) ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsComplete(ctx context.Context, id beta.UserId, options ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationOptions) (ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteResult, error) {
	return c.ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx, id, options, PasswordlessMicrosoftAuthenticatorAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationPasswordlessMicrosoftAuthenticatorMethodClient) ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsOperationOptions, predicate PasswordlessMicrosoftAuthenticatorAuthenticationMethodOperationPredicate) (result ListAuthenticationPasswordlessMicrosoftAuthenticatorMethodsCompleteResult, err error) {
	items := make([]beta.PasswordlessMicrosoftAuthenticatorAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationPasswordlessMicrosoftAuthenticatorMethods(ctx, id, options)
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
