package conditionalaccesauthenticationstrengthauthenticationmethodmode

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

type ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AuthenticationMethodModeDetail
}

type ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AuthenticationMethodModeDetail
}

type ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccesAuthenticationStrengthAuthenticationMethodModes ...
func (c ConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient) ListConditionalAccesAuthenticationStrengthAuthenticationMethodModes(ctx context.Context) (result ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesCustomPager{},
		Path:       "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes",
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

// ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesComplete retrieves all the results into a single object
func (c ConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient) ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesComplete(ctx context.Context) (ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesCompleteResult, error) {
	return c.ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesCompleteMatchingPredicate(ctx, AuthenticationMethodModeDetailOperationPredicate{})
}

// ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccesAuthenticationStrengthAuthenticationMethodModeClient) ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesCompleteMatchingPredicate(ctx context.Context, predicate AuthenticationMethodModeDetailOperationPredicate) (result ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesCompleteResult, err error) {
	items := make([]stable.AuthenticationMethodModeDetail, 0)

	resp, err := c.ListConditionalAccesAuthenticationStrengthAuthenticationMethodModes(ctx)
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

	result = ListConditionalAccesAuthenticationStrengthAuthenticationMethodModesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
