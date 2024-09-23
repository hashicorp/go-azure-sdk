package conditionalaccessauthenticationstrengthauthenticationmethodmode

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

type ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AuthenticationMethodModeDetail
}

type ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AuthenticationMethodModeDetail
}

type ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationOptions struct {
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

func DefaultListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationOptions() ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationOptions {
	return ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationOptions{}
}

func (o ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationOptions) ToOData() *odata.Query {
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

func (o ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccessAuthenticationStrengthAuthenticationMethodModes - List authenticationMethodModes. Get a list of
// all supported authentication methods, or all supported authentication method combinations as a list of
// authenticationMethodModes objects and their properties.
func (c ConditionalAccessAuthenticationStrengthAuthenticationMethodModeClient) ListConditionalAccessAuthenticationStrengthAuthenticationMethodModes(ctx context.Context, options ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationOptions) (result ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesCustomPager{},
		Path:          "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes",
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
		Values *[]stable.AuthenticationMethodModeDetail `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesComplete retrieves all the results into a single object
func (c ConditionalAccessAuthenticationStrengthAuthenticationMethodModeClient) ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesComplete(ctx context.Context, options ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationOptions) (ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesCompleteResult, error) {
	return c.ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesCompleteMatchingPredicate(ctx, options, AuthenticationMethodModeDetailOperationPredicate{})
}

// ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccessAuthenticationStrengthAuthenticationMethodModeClient) ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesCompleteMatchingPredicate(ctx context.Context, options ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesOperationOptions, predicate AuthenticationMethodModeDetailOperationPredicate) (result ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesCompleteResult, err error) {
	items := make([]stable.AuthenticationMethodModeDetail, 0)

	resp, err := c.ListConditionalAccessAuthenticationStrengthAuthenticationMethodModes(ctx, options)
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

	result = ListConditionalAccessAuthenticationStrengthAuthenticationMethodModesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
