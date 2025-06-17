package ownedobject

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOwnedObjectGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Group
}

type ListOwnedObjectGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Group
}

type ListOwnedObjectGroupsOperationOptions struct {
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

func DefaultListOwnedObjectGroupsOperationOptions() ListOwnedObjectGroupsOperationOptions {
	return ListOwnedObjectGroupsOperationOptions{}
}

func (o ListOwnedObjectGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOwnedObjectGroupsOperationOptions) ToOData() *odata.Query {
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

func (o ListOwnedObjectGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOwnedObjectGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOwnedObjectGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOwnedObjectGroups - Get the items of type microsoft.graph.group in the microsoft.graph.directoryObject collection
func (c OwnedObjectClient) ListOwnedObjectGroups(ctx context.Context, id stable.UserId, options ListOwnedObjectGroupsOperationOptions) (result ListOwnedObjectGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOwnedObjectGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/ownedObjects/group", id.ID()),
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
		Values *[]stable.Group `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOwnedObjectGroupsComplete retrieves all the results into a single object
func (c OwnedObjectClient) ListOwnedObjectGroupsComplete(ctx context.Context, id stable.UserId, options ListOwnedObjectGroupsOperationOptions) (ListOwnedObjectGroupsCompleteResult, error) {
	return c.ListOwnedObjectGroupsCompleteMatchingPredicate(ctx, id, options, GroupOperationPredicate{})
}

// ListOwnedObjectGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OwnedObjectClient) ListOwnedObjectGroupsCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListOwnedObjectGroupsOperationOptions, predicate GroupOperationPredicate) (result ListOwnedObjectGroupsCompleteResult, err error) {
	items := make([]stable.Group, 0)

	resp, err := c.ListOwnedObjectGroups(ctx, id, options)
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

	result = ListOwnedObjectGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
