package approleassignedresource

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

type ListAppRoleAssignedResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServicePrincipal
}

type ListAppRoleAssignedResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServicePrincipal
}

type ListAppRoleAssignedResourcesOperationOptions struct {
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

func DefaultListAppRoleAssignedResourcesOperationOptions() ListAppRoleAssignedResourcesOperationOptions {
	return ListAppRoleAssignedResourcesOperationOptions{}
}

func (o ListAppRoleAssignedResourcesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAppRoleAssignedResourcesOperationOptions) ToOData() *odata.Query {
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

func (o ListAppRoleAssignedResourcesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAppRoleAssignedResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppRoleAssignedResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppRoleAssignedResources - Get appRoleAssignedResources from users
func (c AppRoleAssignedResourceClient) ListAppRoleAssignedResources(ctx context.Context, id beta.UserId, options ListAppRoleAssignedResourcesOperationOptions) (result ListAppRoleAssignedResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAppRoleAssignedResourcesCustomPager{},
		Path:          fmt.Sprintf("%s/appRoleAssignedResources", id.ID()),
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
		Values *[]beta.ServicePrincipal `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppRoleAssignedResourcesComplete retrieves all the results into a single object
func (c AppRoleAssignedResourceClient) ListAppRoleAssignedResourcesComplete(ctx context.Context, id beta.UserId, options ListAppRoleAssignedResourcesOperationOptions) (ListAppRoleAssignedResourcesCompleteResult, error) {
	return c.ListAppRoleAssignedResourcesCompleteMatchingPredicate(ctx, id, options, ServicePrincipalOperationPredicate{})
}

// ListAppRoleAssignedResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppRoleAssignedResourceClient) ListAppRoleAssignedResourcesCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListAppRoleAssignedResourcesOperationOptions, predicate ServicePrincipalOperationPredicate) (result ListAppRoleAssignedResourcesCompleteResult, err error) {
	items := make([]beta.ServicePrincipal, 0)

	resp, err := c.ListAppRoleAssignedResources(ctx, id, options)
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

	result = ListAppRoleAssignedResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
