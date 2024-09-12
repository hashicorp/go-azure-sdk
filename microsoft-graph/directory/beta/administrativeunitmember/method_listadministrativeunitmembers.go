package administrativeunitmember

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAdministrativeUnitMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListAdministrativeUnitMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListAdministrativeUnitMembersOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Expand           *odata.Expand
	Filter           *string
	OrderBy          *odata.OrderBy
	Search           *string
	Select           *[]string
	Skip             *int64
	Top              *int64
}

func DefaultListAdministrativeUnitMembersOperationOptions() ListAdministrativeUnitMembersOperationOptions {
	return ListAdministrativeUnitMembersOperationOptions{}
}

func (o ListAdministrativeUnitMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAdministrativeUnitMembersOperationOptions) ToOData() *odata.Query {
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

func (o ListAdministrativeUnitMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAdministrativeUnitMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAdministrativeUnitMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAdministrativeUnitMembers - Get members from directory. Users and groups that are members of this administrative
// unit. Supports $expand.
func (c AdministrativeUnitMemberClient) ListAdministrativeUnitMembers(ctx context.Context, id beta.DirectoryAdministrativeUnitId, options ListAdministrativeUnitMembersOperationOptions) (result ListAdministrativeUnitMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAdministrativeUnitMembersCustomPager{},
		Path:          fmt.Sprintf("%s/members", id.ID()),
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

	temp := make([]beta.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListAdministrativeUnitMembersComplete retrieves all the results into a single object
func (c AdministrativeUnitMemberClient) ListAdministrativeUnitMembersComplete(ctx context.Context, id beta.DirectoryAdministrativeUnitId, options ListAdministrativeUnitMembersOperationOptions) (ListAdministrativeUnitMembersCompleteResult, error) {
	return c.ListAdministrativeUnitMembersCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListAdministrativeUnitMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AdministrativeUnitMemberClient) ListAdministrativeUnitMembersCompleteMatchingPredicate(ctx context.Context, id beta.DirectoryAdministrativeUnitId, options ListAdministrativeUnitMembersOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListAdministrativeUnitMembersCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListAdministrativeUnitMembers(ctx, id, options)
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

	result = ListAdministrativeUnitMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
