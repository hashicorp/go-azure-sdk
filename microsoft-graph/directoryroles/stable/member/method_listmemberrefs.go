package member

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

type ListMemberRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListMemberRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListMemberRefsOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Filter           *string
	OrderBy          *odata.OrderBy
	Search           *string
	Skip             *int64
	Top              *int64
}

func DefaultListMemberRefsOperationOptions() ListMemberRefsOperationOptions {
	return ListMemberRefsOperationOptions{}
}

func (o ListMemberRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMemberRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListMemberRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMemberRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMemberRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMemberRefs - List members of a directory role. Retrieve the list of principals that are assigned to the directory
// role. You can use both the object ID and template ID of the directoryRole with this API. The template ID of a
// built-in role is immutable and can be seen in the role description on the Microsoft Entra admin center. For details,
// see Role template IDs.
func (c MemberClient) ListMemberRefs(ctx context.Context, id stable.DirectoryRoleId, options ListMemberRefsOperationOptions) (result ListMemberRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMemberRefsCustomPager{},
		Path:          fmt.Sprintf("%s/members/$ref", id.ID()),
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

// ListMemberRefsComplete retrieves all the results into a single object
func (c MemberClient) ListMemberRefsComplete(ctx context.Context, id stable.DirectoryRoleId, options ListMemberRefsOperationOptions) (ListMemberRefsCompleteResult, error) {
	return c.ListMemberRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListMemberRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MemberClient) ListMemberRefsCompleteMatchingPredicate(ctx context.Context, id stable.DirectoryRoleId, options ListMemberRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListMemberRefsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListMemberRefs(ctx, id, options)
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

	result = ListMemberRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
