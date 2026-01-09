package authenticationfido2method

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAuthenticationFido2MethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Fido2AuthenticationMethod
}

type ListAuthenticationFido2MethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Fido2AuthenticationMethod
}

type ListAuthenticationFido2MethodsOperationOptions struct {
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

func DefaultListAuthenticationFido2MethodsOperationOptions() ListAuthenticationFido2MethodsOperationOptions {
	return ListAuthenticationFido2MethodsOperationOptions{}
}

func (o ListAuthenticationFido2MethodsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationFido2MethodsOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationFido2MethodsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAuthenticationFido2MethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationFido2MethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationFido2Methods - List fido2AuthenticationMethod. Retrieve a list of a user's FIDO2 Security Key
// Authentication Method objects and their properties.
func (c AuthenticationFido2MethodClient) ListAuthenticationFido2Methods(ctx context.Context, options ListAuthenticationFido2MethodsOperationOptions) (result ListAuthenticationFido2MethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationFido2MethodsCustomPager{},
		Path:          "/me/authentication/fido2Methods",
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
		Values *[]beta.Fido2AuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationFido2MethodsComplete retrieves all the results into a single object
func (c AuthenticationFido2MethodClient) ListAuthenticationFido2MethodsComplete(ctx context.Context, options ListAuthenticationFido2MethodsOperationOptions) (ListAuthenticationFido2MethodsCompleteResult, error) {
	return c.ListAuthenticationFido2MethodsCompleteMatchingPredicate(ctx, options, Fido2AuthenticationMethodOperationPredicate{})
}

// ListAuthenticationFido2MethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationFido2MethodClient) ListAuthenticationFido2MethodsCompleteMatchingPredicate(ctx context.Context, options ListAuthenticationFido2MethodsOperationOptions, predicate Fido2AuthenticationMethodOperationPredicate) (result ListAuthenticationFido2MethodsCompleteResult, err error) {
	items := make([]beta.Fido2AuthenticationMethod, 0)

	resp, err := c.ListAuthenticationFido2Methods(ctx, options)
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

	result = ListAuthenticationFido2MethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
