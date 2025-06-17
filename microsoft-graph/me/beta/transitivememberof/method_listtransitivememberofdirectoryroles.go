package transitivememberof

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

type ListTransitiveMemberOfDirectoryRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryRole
}

type ListTransitiveMemberOfDirectoryRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryRole
}

type ListTransitiveMemberOfDirectoryRolesOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Expand           *odata.Expand
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	RetryFunc        client.RequestRetryFunc
	Search           *string
	Select           *[]string
	Skip             *int64
	Top              *int64
}

func DefaultListTransitiveMemberOfDirectoryRolesOperationOptions() ListTransitiveMemberOfDirectoryRolesOperationOptions {
	return ListTransitiveMemberOfDirectoryRolesOperationOptions{}
}

func (o ListTransitiveMemberOfDirectoryRolesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTransitiveMemberOfDirectoryRolesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
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

func (o ListTransitiveMemberOfDirectoryRolesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTransitiveMemberOfDirectoryRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTransitiveMemberOfDirectoryRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTransitiveMemberOfDirectoryRoles - Get the items of type microsoft.graph.directoryRole in the
// microsoft.graph.directoryObject collection
func (c TransitiveMemberOfClient) ListTransitiveMemberOfDirectoryRoles(ctx context.Context, options ListTransitiveMemberOfDirectoryRolesOperationOptions) (result ListTransitiveMemberOfDirectoryRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTransitiveMemberOfDirectoryRolesCustomPager{},
		Path:          "/me/transitiveMemberOf/directoryRole",
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
		Values *[]beta.DirectoryRole `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTransitiveMemberOfDirectoryRolesComplete retrieves all the results into a single object
func (c TransitiveMemberOfClient) ListTransitiveMemberOfDirectoryRolesComplete(ctx context.Context, options ListTransitiveMemberOfDirectoryRolesOperationOptions) (ListTransitiveMemberOfDirectoryRolesCompleteResult, error) {
	return c.ListTransitiveMemberOfDirectoryRolesCompleteMatchingPredicate(ctx, options, DirectoryRoleOperationPredicate{})
}

// ListTransitiveMemberOfDirectoryRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TransitiveMemberOfClient) ListTransitiveMemberOfDirectoryRolesCompleteMatchingPredicate(ctx context.Context, options ListTransitiveMemberOfDirectoryRolesOperationOptions, predicate DirectoryRoleOperationPredicate) (result ListTransitiveMemberOfDirectoryRolesCompleteResult, err error) {
	items := make([]beta.DirectoryRole, 0)

	resp, err := c.ListTransitiveMemberOfDirectoryRoles(ctx, options)
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

	result = ListTransitiveMemberOfDirectoryRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
