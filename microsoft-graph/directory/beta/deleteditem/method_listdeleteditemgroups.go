package deleteditem

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

type ListDeletedItemGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Group
}

type ListDeletedItemGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Group
}

type ListDeletedItemGroupsOperationOptions struct {
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

func DefaultListDeletedItemGroupsOperationOptions() ListDeletedItemGroupsOperationOptions {
	return ListDeletedItemGroupsOperationOptions{}
}

func (o ListDeletedItemGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeletedItemGroupsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeletedItemGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeletedItemGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeletedItemGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeletedItemGroups - List deletedItems (directory objects). Retrieve a list of recently deleted directory objects
// from deleted items. The following types are supported: - administrativeUnit - application - certificateBasedAuthPki -
// [certificateAuthorityDetail](../resources/certificateauthoritydetail.md - externalUserProfile - group -
// pendingExternalUserProfile - servicePrincipal - user
func (c DeletedItemClient) ListDeletedItemGroups(ctx context.Context, options ListDeletedItemGroupsOperationOptions) (result ListDeletedItemGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeletedItemGroupsCustomPager{},
		Path:          "/directory/deletedItems/group",
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
		Values *[]beta.Group `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeletedItemGroupsComplete retrieves all the results into a single object
func (c DeletedItemClient) ListDeletedItemGroupsComplete(ctx context.Context, options ListDeletedItemGroupsOperationOptions) (ListDeletedItemGroupsCompleteResult, error) {
	return c.ListDeletedItemGroupsCompleteMatchingPredicate(ctx, options, GroupOperationPredicate{})
}

// ListDeletedItemGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeletedItemClient) ListDeletedItemGroupsCompleteMatchingPredicate(ctx context.Context, options ListDeletedItemGroupsOperationOptions, predicate GroupOperationPredicate) (result ListDeletedItemGroupsCompleteResult, err error) {
	items := make([]beta.Group, 0)

	resp, err := c.ListDeletedItemGroups(ctx, options)
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

	result = ListDeletedItemGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
