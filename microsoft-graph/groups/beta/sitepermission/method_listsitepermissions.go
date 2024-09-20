package sitepermission

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

type ListSitePermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Permission
}

type ListSitePermissionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Permission
}

type ListSitePermissionsOperationOptions struct {
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

func DefaultListSitePermissionsOperationOptions() ListSitePermissionsOperationOptions {
	return ListSitePermissionsOperationOptions{}
}

func (o ListSitePermissionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSitePermissionsOperationOptions) ToOData() *odata.Query {
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

func (o ListSitePermissionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSitePermissionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSitePermissionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSitePermissions - Get permissions from groups. The permissions associated with the site. Nullable.
func (c SitePermissionClient) ListSitePermissions(ctx context.Context, id beta.GroupIdSiteId, options ListSitePermissionsOperationOptions) (result ListSitePermissionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSitePermissionsCustomPager{},
		Path:          fmt.Sprintf("%s/permissions", id.ID()),
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
		Values *[]beta.Permission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSitePermissionsComplete retrieves all the results into a single object
func (c SitePermissionClient) ListSitePermissionsComplete(ctx context.Context, id beta.GroupIdSiteId, options ListSitePermissionsOperationOptions) (ListSitePermissionsCompleteResult, error) {
	return c.ListSitePermissionsCompleteMatchingPredicate(ctx, id, options, PermissionOperationPredicate{})
}

// ListSitePermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitePermissionClient) ListSitePermissionsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteId, options ListSitePermissionsOperationOptions, predicate PermissionOperationPredicate) (result ListSitePermissionsCompleteResult, err error) {
	items := make([]beta.Permission, 0)

	resp, err := c.ListSitePermissions(ctx, id, options)
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

	result = ListSitePermissionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
