package crosstenantaccesspolicypartner

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

type ListCrossTenantAccessPolicyPartnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CrossTenantAccessPolicyConfigurationPartner
}

type ListCrossTenantAccessPolicyPartnersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CrossTenantAccessPolicyConfigurationPartner
}

type ListCrossTenantAccessPolicyPartnersOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListCrossTenantAccessPolicyPartnersOperationOptions() ListCrossTenantAccessPolicyPartnersOperationOptions {
	return ListCrossTenantAccessPolicyPartnersOperationOptions{}
}

func (o ListCrossTenantAccessPolicyPartnersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCrossTenantAccessPolicyPartnersOperationOptions) ToOData() *odata.Query {
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

func (o ListCrossTenantAccessPolicyPartnersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCrossTenantAccessPolicyPartnersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCrossTenantAccessPolicyPartnersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCrossTenantAccessPolicyPartners - List partners. Get a list of all partner configurations within a cross-tenant
// access policy. You can also use the $expand parameter to list the user synchronization policy for all partner
// configurations.
func (c CrossTenantAccessPolicyPartnerClient) ListCrossTenantAccessPolicyPartners(ctx context.Context, options ListCrossTenantAccessPolicyPartnersOperationOptions) (result ListCrossTenantAccessPolicyPartnersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCrossTenantAccessPolicyPartnersCustomPager{},
		Path:          "/policies/crossTenantAccessPolicy/partners",
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
		Values *[]beta.CrossTenantAccessPolicyConfigurationPartner `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCrossTenantAccessPolicyPartnersComplete retrieves all the results into a single object
func (c CrossTenantAccessPolicyPartnerClient) ListCrossTenantAccessPolicyPartnersComplete(ctx context.Context, options ListCrossTenantAccessPolicyPartnersOperationOptions) (ListCrossTenantAccessPolicyPartnersCompleteResult, error) {
	return c.ListCrossTenantAccessPolicyPartnersCompleteMatchingPredicate(ctx, options, CrossTenantAccessPolicyConfigurationPartnerOperationPredicate{})
}

// ListCrossTenantAccessPolicyPartnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CrossTenantAccessPolicyPartnerClient) ListCrossTenantAccessPolicyPartnersCompleteMatchingPredicate(ctx context.Context, options ListCrossTenantAccessPolicyPartnersOperationOptions, predicate CrossTenantAccessPolicyConfigurationPartnerOperationPredicate) (result ListCrossTenantAccessPolicyPartnersCompleteResult, err error) {
	items := make([]beta.CrossTenantAccessPolicyConfigurationPartner, 0)

	resp, err := c.ListCrossTenantAccessPolicyPartners(ctx, options)
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

	result = ListCrossTenantAccessPolicyPartnersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
