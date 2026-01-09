package rolescopetag

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRoleScopeTagsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RoleScopeTag
}

type ListRoleScopeTagsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RoleScopeTag
}

type ListRoleScopeTagsOperationOptions struct {
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

func DefaultListRoleScopeTagsOperationOptions() ListRoleScopeTagsOperationOptions {
	return ListRoleScopeTagsOperationOptions{}
}

func (o ListRoleScopeTagsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRoleScopeTagsOperationOptions) ToOData() *odata.Query {
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

func (o ListRoleScopeTagsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRoleScopeTagsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleScopeTagsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleScopeTags - Get roleScopeTags from deviceManagement. The Role Scope Tags.
func (c RoleScopeTagClient) ListRoleScopeTags(ctx context.Context, options ListRoleScopeTagsOperationOptions) (result ListRoleScopeTagsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRoleScopeTagsCustomPager{},
		Path:          "/deviceManagement/roleScopeTags",
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
		Values *[]beta.RoleScopeTag `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleScopeTagsComplete retrieves all the results into a single object
func (c RoleScopeTagClient) ListRoleScopeTagsComplete(ctx context.Context, options ListRoleScopeTagsOperationOptions) (ListRoleScopeTagsCompleteResult, error) {
	return c.ListRoleScopeTagsCompleteMatchingPredicate(ctx, options, RoleScopeTagOperationPredicate{})
}

// ListRoleScopeTagsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleScopeTagClient) ListRoleScopeTagsCompleteMatchingPredicate(ctx context.Context, options ListRoleScopeTagsOperationOptions, predicate RoleScopeTagOperationPredicate) (result ListRoleScopeTagsCompleteResult, err error) {
	items := make([]beta.RoleScopeTag, 0)

	resp, err := c.ListRoleScopeTags(ctx, options)
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

	result = ListRoleScopeTagsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
