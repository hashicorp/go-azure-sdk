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

type ListAuthenticationTemporaryAccessPassMethodsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListAuthenticationTemporaryAccessPassMethodsOperationOptions() ListAuthenticationTemporaryAccessPassMethodsOperationOptions {
	return ListAuthenticationTemporaryAccessPassMethodsOperationOptions{}
}

func (o ListAuthenticationTemporaryAccessPassMethodsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationTemporaryAccessPassMethodsOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationTemporaryAccessPassMethodsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListAuthenticationTemporaryAccessPassMethods - List temporaryAccessPassMethods. Retrieve a list of a user's
// temporaryAccessPassAuthenticationMethod objects and their properties. This API will only return a single object in
// the collection as a user can have only one Temporary Access Pass method.
func (c AuthenticationTemporaryAccessPassMethodClient) ListAuthenticationTemporaryAccessPassMethods(ctx context.Context, id stable.UserId, options ListAuthenticationTemporaryAccessPassMethodsOperationOptions) (result ListAuthenticationTemporaryAccessPassMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationTemporaryAccessPassMethodsCustomPager{},
		Path:          fmt.Sprintf("%s/authentication/temporaryAccessPassMethods", id.ID()),
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
func (c AuthenticationTemporaryAccessPassMethodClient) ListAuthenticationTemporaryAccessPassMethodsComplete(ctx context.Context, id stable.UserId, options ListAuthenticationTemporaryAccessPassMethodsOperationOptions) (ListAuthenticationTemporaryAccessPassMethodsCompleteResult, error) {
	return c.ListAuthenticationTemporaryAccessPassMethodsCompleteMatchingPredicate(ctx, id, options, TemporaryAccessPassAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationTemporaryAccessPassMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationTemporaryAccessPassMethodClient) ListAuthenticationTemporaryAccessPassMethodsCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListAuthenticationTemporaryAccessPassMethodsOperationOptions, predicate TemporaryAccessPassAuthenticationMethodOperationPredicate) (result ListAuthenticationTemporaryAccessPassMethodsCompleteResult, err error) {
	items := make([]stable.TemporaryAccessPassAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationTemporaryAccessPassMethods(ctx, id, options)
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
