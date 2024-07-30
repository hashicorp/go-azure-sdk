package authenticationmethodspolicyauthenticationmethodconfiguration

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

type ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AuthenticationMethodConfiguration
}

type ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AuthenticationMethodConfiguration
}

type ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationMethodsPolicyAuthenticationMethodConfigurations ...
func (c AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient) ListAuthenticationMethodsPolicyAuthenticationMethodConfigurations(ctx context.Context) (result ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsCustomPager{},
		Path:       "/policies/authenticationMethodsPolicy/authenticationMethodConfigurations",
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
		Values *[]stable.AuthenticationMethodConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsComplete retrieves all the results into a single object
func (c AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient) ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsComplete(ctx context.Context) (ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsCompleteResult, error) {
	return c.ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsCompleteMatchingPredicate(ctx, AuthenticationMethodConfigurationOperationPredicate{})
}

// ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient) ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsCompleteMatchingPredicate(ctx context.Context, predicate AuthenticationMethodConfigurationOperationPredicate) (result ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsCompleteResult, err error) {
	items := make([]stable.AuthenticationMethodConfiguration, 0)

	resp, err := c.ListAuthenticationMethodsPolicyAuthenticationMethodConfigurations(ctx)
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

	result = ListAuthenticationMethodsPolicyAuthenticationMethodConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
