package cloudpcresourcenamespace

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

type ListCloudPCResourceNamespacesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRbacResourceNamespace
}

type ListCloudPCResourceNamespacesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRbacResourceNamespace
}

type ListCloudPCResourceNamespacesOperationOptions struct {
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

func DefaultListCloudPCResourceNamespacesOperationOptions() ListCloudPCResourceNamespacesOperationOptions {
	return ListCloudPCResourceNamespacesOperationOptions{}
}

func (o ListCloudPCResourceNamespacesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCloudPCResourceNamespacesOperationOptions) ToOData() *odata.Query {
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

func (o ListCloudPCResourceNamespacesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCloudPCResourceNamespacesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudPCResourceNamespacesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudPCResourceNamespaces - Get resourceNamespaces from roleManagement
func (c CloudPCResourceNamespaceClient) ListCloudPCResourceNamespaces(ctx context.Context, options ListCloudPCResourceNamespacesOperationOptions) (result ListCloudPCResourceNamespacesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCloudPCResourceNamespacesCustomPager{},
		Path:          "/roleManagement/cloudPC/resourceNamespaces",
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
		Values *[]beta.UnifiedRbacResourceNamespace `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCloudPCResourceNamespacesComplete retrieves all the results into a single object
func (c CloudPCResourceNamespaceClient) ListCloudPCResourceNamespacesComplete(ctx context.Context, options ListCloudPCResourceNamespacesOperationOptions) (ListCloudPCResourceNamespacesCompleteResult, error) {
	return c.ListCloudPCResourceNamespacesCompleteMatchingPredicate(ctx, options, UnifiedRbacResourceNamespaceOperationPredicate{})
}

// ListCloudPCResourceNamespacesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudPCResourceNamespaceClient) ListCloudPCResourceNamespacesCompleteMatchingPredicate(ctx context.Context, options ListCloudPCResourceNamespacesOperationOptions, predicate UnifiedRbacResourceNamespaceOperationPredicate) (result ListCloudPCResourceNamespacesCompleteResult, err error) {
	items := make([]beta.UnifiedRbacResourceNamespace, 0)

	resp, err := c.ListCloudPCResourceNamespaces(ctx, options)
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

	result = ListCloudPCResourceNamespacesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
