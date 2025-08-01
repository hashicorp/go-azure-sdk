package memberswithlicenseerror

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

type ListMembersWithLicenseErrorUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.User
}

type ListMembersWithLicenseErrorUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.User
}

type ListMembersWithLicenseErrorUsersOperationOptions struct {
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

func DefaultListMembersWithLicenseErrorUsersOperationOptions() ListMembersWithLicenseErrorUsersOperationOptions {
	return ListMembersWithLicenseErrorUsersOperationOptions{}
}

func (o ListMembersWithLicenseErrorUsersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMembersWithLicenseErrorUsersOperationOptions) ToOData() *odata.Query {
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

func (o ListMembersWithLicenseErrorUsersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMembersWithLicenseErrorUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMembersWithLicenseErrorUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMembersWithLicenseErrorUsers - Get the items of type microsoft.graph.user in the microsoft.graph.directoryObject
// collection
func (c MembersWithLicenseErrorClient) ListMembersWithLicenseErrorUsers(ctx context.Context, id stable.GroupId, options ListMembersWithLicenseErrorUsersOperationOptions) (result ListMembersWithLicenseErrorUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMembersWithLicenseErrorUsersCustomPager{},
		Path:          fmt.Sprintf("%s/membersWithLicenseErrors/user", id.ID()),
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
		Values *[]stable.User `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMembersWithLicenseErrorUsersComplete retrieves all the results into a single object
func (c MembersWithLicenseErrorClient) ListMembersWithLicenseErrorUsersComplete(ctx context.Context, id stable.GroupId, options ListMembersWithLicenseErrorUsersOperationOptions) (ListMembersWithLicenseErrorUsersCompleteResult, error) {
	return c.ListMembersWithLicenseErrorUsersCompleteMatchingPredicate(ctx, id, options, UserOperationPredicate{})
}

// ListMembersWithLicenseErrorUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MembersWithLicenseErrorClient) ListMembersWithLicenseErrorUsersCompleteMatchingPredicate(ctx context.Context, id stable.GroupId, options ListMembersWithLicenseErrorUsersOperationOptions, predicate UserOperationPredicate) (result ListMembersWithLicenseErrorUsersCompleteResult, err error) {
	items := make([]stable.User, 0)

	resp, err := c.ListMembersWithLicenseErrorUsers(ctx, id, options)
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

	result = ListMembersWithLicenseErrorUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
