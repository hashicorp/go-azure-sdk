package scopedmember

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

type ListScopedMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ScopedRoleMembership
}

type ListScopedMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ScopedRoleMembership
}

type ListScopedMembersOperationOptions struct {
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

func DefaultListScopedMembersOperationOptions() ListScopedMembersOperationOptions {
	return ListScopedMembersOperationOptions{}
}

func (o ListScopedMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListScopedMembersOperationOptions) ToOData() *odata.Query {
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

func (o ListScopedMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListScopedMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListScopedMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListScopedMembers - List scopedMembers for a directory role. Retrieve a list of scopedRoleMembership objects for a
// directory role.
func (c ScopedMemberClient) ListScopedMembers(ctx context.Context, id beta.DirectoryRoleId, options ListScopedMembersOperationOptions) (result ListScopedMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListScopedMembersCustomPager{},
		Path:          fmt.Sprintf("%s/scopedMembers", id.ID()),
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
		Values *[]beta.ScopedRoleMembership `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListScopedMembersComplete retrieves all the results into a single object
func (c ScopedMemberClient) ListScopedMembersComplete(ctx context.Context, id beta.DirectoryRoleId, options ListScopedMembersOperationOptions) (ListScopedMembersCompleteResult, error) {
	return c.ListScopedMembersCompleteMatchingPredicate(ctx, id, options, ScopedRoleMembershipOperationPredicate{})
}

// ListScopedMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ScopedMemberClient) ListScopedMembersCompleteMatchingPredicate(ctx context.Context, id beta.DirectoryRoleId, options ListScopedMembersOperationOptions, predicate ScopedRoleMembershipOperationPredicate) (result ListScopedMembersCompleteResult, err error) {
	items := make([]beta.ScopedRoleMembership, 0)

	resp, err := c.ListScopedMembers(ctx, id, options)
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

	result = ListScopedMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
