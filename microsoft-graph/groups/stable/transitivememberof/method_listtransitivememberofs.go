package transitivememberof

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

type ListTransitiveMemberOfsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListTransitiveMemberOfsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListTransitiveMemberOfsOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Expand           *odata.Expand
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	Search           *string
	Select           *[]string
	Skip             *int64
	Top              *int64
}

func DefaultListTransitiveMemberOfsOperationOptions() ListTransitiveMemberOfsOperationOptions {
	return ListTransitiveMemberOfsOperationOptions{}
}

func (o ListTransitiveMemberOfsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTransitiveMemberOfsOperationOptions) ToOData() *odata.Query {
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

func (o ListTransitiveMemberOfsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTransitiveMemberOfsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTransitiveMemberOfsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTransitiveMemberOfs - List group transitive memberOf. Get groups that the group is a member of. This operation is
// transitive and will also include all groups that this group is a nested member of. Unlike getting a user's Microsoft
// 365 groups, this returns all types of groups, not just Microsoft 365 groups.
func (c TransitiveMemberOfClient) ListTransitiveMemberOfs(ctx context.Context, id stable.GroupId, options ListTransitiveMemberOfsOperationOptions) (result ListTransitiveMemberOfsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTransitiveMemberOfsCustomPager{},
		Path:          fmt.Sprintf("%s/transitiveMemberOf", id.ID()),
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

	temp := make([]stable.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListTransitiveMemberOfsComplete retrieves all the results into a single object
func (c TransitiveMemberOfClient) ListTransitiveMemberOfsComplete(ctx context.Context, id stable.GroupId, options ListTransitiveMemberOfsOperationOptions) (ListTransitiveMemberOfsCompleteResult, error) {
	return c.ListTransitiveMemberOfsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListTransitiveMemberOfsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TransitiveMemberOfClient) ListTransitiveMemberOfsCompleteMatchingPredicate(ctx context.Context, id stable.GroupId, options ListTransitiveMemberOfsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListTransitiveMemberOfsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListTransitiveMemberOfs(ctx, id, options)
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

	result = ListTransitiveMemberOfsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
