package authenticationmethod

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAuthenticationMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AuthenticationMethod
}

type ListAuthenticationMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AuthenticationMethod
}

type ListAuthenticationMethodsOperationOptions struct {
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

func DefaultListAuthenticationMethodsOperationOptions() ListAuthenticationMethodsOperationOptions {
	return ListAuthenticationMethodsOperationOptions{}
}

func (o ListAuthenticationMethodsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationMethodsOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationMethodsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListAuthenticationMethods - Get methods from users. Represents all authentication methods registered to a user.
func (c AuthenticationMethodClient) ListAuthenticationMethods(ctx context.Context, id stable.UserId, options ListAuthenticationMethodsOperationOptions) (result ListAuthenticationMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationMethodsCustomPager{},
		Path:          fmt.Sprintf("%s/authentication/methods", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.AuthenticationMethod, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalAuthenticationMethodImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.AuthenticationMethod (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListAuthenticationMethodsComplete retrieves all the results into a single object
func (c AuthenticationMethodClient) ListAuthenticationMethodsComplete(ctx context.Context, id stable.UserId, options ListAuthenticationMethodsOperationOptions) (ListAuthenticationMethodsCompleteResult, error) {
	return c.ListAuthenticationMethodsCompleteMatchingPredicate(ctx, id, options, AuthenticationMethodOperationPredicate{})
}

// ListAuthenticationMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationMethodClient) ListAuthenticationMethodsCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListAuthenticationMethodsOperationOptions, predicate AuthenticationMethodOperationPredicate) (result ListAuthenticationMethodsCompleteResult, err error) {
	items := make([]stable.AuthenticationMethod, 0)

	resp, err := c.ListAuthenticationMethods(ctx, id, options)
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
