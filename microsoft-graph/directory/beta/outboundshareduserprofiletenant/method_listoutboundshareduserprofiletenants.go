package outboundshareduserprofiletenant

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

type ListOutboundSharedUserProfileTenantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TenantReference
}

type ListOutboundSharedUserProfileTenantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TenantReference
}

type ListOutboundSharedUserProfileTenantsOperationOptions struct {
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

func DefaultListOutboundSharedUserProfileTenantsOperationOptions() ListOutboundSharedUserProfileTenantsOperationOptions {
	return ListOutboundSharedUserProfileTenantsOperationOptions{}
}

func (o ListOutboundSharedUserProfileTenantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOutboundSharedUserProfileTenantsOperationOptions) ToOData() *odata.Query {
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

func (o ListOutboundSharedUserProfileTenantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOutboundSharedUserProfileTenantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutboundSharedUserProfileTenantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutboundSharedUserProfileTenants - Get tenantReferences. List the tenant references of an
// outboundSharedUserProfile.
func (c OutboundSharedUserProfileTenantClient) ListOutboundSharedUserProfileTenants(ctx context.Context, id beta.DirectoryOutboundSharedUserProfileId, options ListOutboundSharedUserProfileTenantsOperationOptions) (result ListOutboundSharedUserProfileTenantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOutboundSharedUserProfileTenantsCustomPager{},
		Path:          fmt.Sprintf("%s/tenants", id.ID()),
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
		Values *[]beta.TenantReference `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutboundSharedUserProfileTenantsComplete retrieves all the results into a single object
func (c OutboundSharedUserProfileTenantClient) ListOutboundSharedUserProfileTenantsComplete(ctx context.Context, id beta.DirectoryOutboundSharedUserProfileId, options ListOutboundSharedUserProfileTenantsOperationOptions) (ListOutboundSharedUserProfileTenantsCompleteResult, error) {
	return c.ListOutboundSharedUserProfileTenantsCompleteMatchingPredicate(ctx, id, options, TenantReferenceOperationPredicate{})
}

// ListOutboundSharedUserProfileTenantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutboundSharedUserProfileTenantClient) ListOutboundSharedUserProfileTenantsCompleteMatchingPredicate(ctx context.Context, id beta.DirectoryOutboundSharedUserProfileId, options ListOutboundSharedUserProfileTenantsOperationOptions, predicate TenantReferenceOperationPredicate) (result ListOutboundSharedUserProfileTenantsCompleteResult, err error) {
	items := make([]beta.TenantReference, 0)

	resp, err := c.ListOutboundSharedUserProfileTenants(ctx, id, options)
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

	result = ListOutboundSharedUserProfileTenantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
