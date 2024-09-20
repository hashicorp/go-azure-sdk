package conditionalaccessauthenticationstrengthpolicycombinationconfiguration

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

type ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AuthenticationCombinationConfiguration
}

type ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AuthenticationCombinationConfiguration
}

type ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions struct {
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

func DefaultListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions() ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions {
	return ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions{}
}

func (o ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurations - List combinationConfigurations. Get the
// authenticationCombinationConfiguration objects for an authentication strength policy. The objects can be of one or
// more of the following derived types: * fido2combinationConfigurations * x509certificatecombinationconfiguration
// authenticationCombinationConfiguration objects are supported only for custom authentication strengths.
func (c ConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient) ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurations(ctx context.Context, id stable.IdentityConditionalAccessAuthenticationStrengthPolicyId, options ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions) (result ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsCustomPager{},
		Path:          fmt.Sprintf("%s/combinationConfigurations", id.ID()),
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

	temp := make([]stable.AuthenticationCombinationConfiguration, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalAuthenticationCombinationConfigurationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.AuthenticationCombinationConfiguration (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsComplete retrieves all the results into a single object
func (c ConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient) ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsComplete(ctx context.Context, id stable.IdentityConditionalAccessAuthenticationStrengthPolicyId, options ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions) (ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult, error) {
	return c.ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate(ctx, id, options, AuthenticationCombinationConfigurationOperationPredicate{})
}

// ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationClient) ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityConditionalAccessAuthenticationStrengthPolicyId, options ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions, predicate AuthenticationCombinationConfigurationOperationPredicate) (result ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult, err error) {
	items := make([]stable.AuthenticationCombinationConfiguration, 0)

	resp, err := c.ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurations(ctx, id, options)
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

	result = ListConditionalAccessAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
