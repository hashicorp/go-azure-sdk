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

type ListAuthenticationMicrosoftAuthenticatorMethodsOperationOptions struct {
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

func DefaultListAuthenticationMicrosoftAuthenticatorMethodsOperationOptions() ListAuthenticationMicrosoftAuthenticatorMethodsOperationOptions {
	return ListAuthenticationMicrosoftAuthenticatorMethodsOperationOptions{}
}

func (o ListAuthenticationMicrosoftAuthenticatorMethodsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationMicrosoftAuthenticatorMethodsOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationMicrosoftAuthenticatorMethodsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListAuthenticationMicrosoftAuthenticatorMethods - Get microsoftAuthenticatorMethods from me. The details of the
// Microsoft Authenticator app registered to a user for authentication.
func (c AuthenticationMicrosoftAuthenticatorMethodClient) ListAuthenticationMicrosoftAuthenticatorMethods(ctx context.Context, options ListAuthenticationMicrosoftAuthenticatorMethodsOperationOptions) (result ListAuthenticationMicrosoftAuthenticatorMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationMicrosoftAuthenticatorMethodsCustomPager{},
		Path:          "/me/authentication/microsoftAuthenticatorMethods",
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
		Values *[]beta.MicrosoftAuthenticatorAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationMicrosoftAuthenticatorMethodsComplete retrieves all the results into a single object
func (c AuthenticationMicrosoftAuthenticatorMethodClient) ListAuthenticationMicrosoftAuthenticatorMethodsComplete(ctx context.Context, options ListAuthenticationMicrosoftAuthenticatorMethodsOperationOptions) (ListAuthenticationMicrosoftAuthenticatorMethodsCompleteResult, error) {
	return c.ListAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx, options, MicrosoftAuthenticatorAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationMicrosoftAuthenticatorMethodClient) ListAuthenticationMicrosoftAuthenticatorMethodsCompleteMatchingPredicate(ctx context.Context, options ListAuthenticationMicrosoftAuthenticatorMethodsOperationOptions, predicate MicrosoftAuthenticatorAuthenticationMethodOperationPredicate) (result ListAuthenticationMicrosoftAuthenticatorMethodsCompleteResult, err error) {
	items := make([]beta.MicrosoftAuthenticatorAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationMicrosoftAuthenticatorMethods(ctx, options)
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
