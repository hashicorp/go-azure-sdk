package teampermissiongrant

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

type ListTeamPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ResourceSpecificPermissionGrant
}

type ListTeamPermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ResourceSpecificPermissionGrant
}

type ListTeamPermissionGrantsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListTeamPermissionGrantsOperationOptions() ListTeamPermissionGrantsOperationOptions {
	return ListTeamPermissionGrantsOperationOptions{}
}

func (o ListTeamPermissionGrantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamPermissionGrantsOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamPermissionGrantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamPermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamPermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamPermissionGrants - Get permissionGrants from groups. A collection of permissions granted to apps to access
// the team.
func (c TeamPermissionGrantClient) ListTeamPermissionGrants(ctx context.Context, id stable.GroupId, options ListTeamPermissionGrantsOperationOptions) (result ListTeamPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamPermissionGrantsCustomPager{},
		Path:          fmt.Sprintf("%s/team/permissionGrants", id.ID()),
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
		Values *[]stable.ResourceSpecificPermissionGrant `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamPermissionGrantsComplete retrieves all the results into a single object
func (c TeamPermissionGrantClient) ListTeamPermissionGrantsComplete(ctx context.Context, id stable.GroupId, options ListTeamPermissionGrantsOperationOptions) (ListTeamPermissionGrantsCompleteResult, error) {
	return c.ListTeamPermissionGrantsCompleteMatchingPredicate(ctx, id, options, ResourceSpecificPermissionGrantOperationPredicate{})
}

// ListTeamPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamPermissionGrantClient) ListTeamPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id stable.GroupId, options ListTeamPermissionGrantsOperationOptions, predicate ResourceSpecificPermissionGrantOperationPredicate) (result ListTeamPermissionGrantsCompleteResult, err error) {
	items := make([]stable.ResourceSpecificPermissionGrant, 0)

	resp, err := c.ListTeamPermissionGrants(ctx, id, options)
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

	result = ListTeamPermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
