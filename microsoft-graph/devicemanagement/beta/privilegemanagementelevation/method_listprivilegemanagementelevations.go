package privilegemanagementelevation

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

type ListPrivilegeManagementElevationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PrivilegeManagementElevation
}

type ListPrivilegeManagementElevationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PrivilegeManagementElevation
}

type ListPrivilegeManagementElevationsOperationOptions struct {
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

func DefaultListPrivilegeManagementElevationsOperationOptions() ListPrivilegeManagementElevationsOperationOptions {
	return ListPrivilegeManagementElevationsOperationOptions{}
}

func (o ListPrivilegeManagementElevationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPrivilegeManagementElevationsOperationOptions) ToOData() *odata.Query {
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

func (o ListPrivilegeManagementElevationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPrivilegeManagementElevationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegeManagementElevationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegeManagementElevations - Get privilegeManagementElevations from deviceManagement. The endpoint privilege
// management elevation event entity contains elevation details.
func (c PrivilegeManagementElevationClient) ListPrivilegeManagementElevations(ctx context.Context, options ListPrivilegeManagementElevationsOperationOptions) (result ListPrivilegeManagementElevationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPrivilegeManagementElevationsCustomPager{},
		Path:          "/deviceManagement/privilegeManagementElevations",
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
		Values *[]beta.PrivilegeManagementElevation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegeManagementElevationsComplete retrieves all the results into a single object
func (c PrivilegeManagementElevationClient) ListPrivilegeManagementElevationsComplete(ctx context.Context, options ListPrivilegeManagementElevationsOperationOptions) (ListPrivilegeManagementElevationsCompleteResult, error) {
	return c.ListPrivilegeManagementElevationsCompleteMatchingPredicate(ctx, options, PrivilegeManagementElevationOperationPredicate{})
}

// ListPrivilegeManagementElevationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegeManagementElevationClient) ListPrivilegeManagementElevationsCompleteMatchingPredicate(ctx context.Context, options ListPrivilegeManagementElevationsOperationOptions, predicate PrivilegeManagementElevationOperationPredicate) (result ListPrivilegeManagementElevationsCompleteResult, err error) {
	items := make([]beta.PrivilegeManagementElevation, 0)

	resp, err := c.ListPrivilegeManagementElevations(ctx, options)
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

	result = ListPrivilegeManagementElevationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
