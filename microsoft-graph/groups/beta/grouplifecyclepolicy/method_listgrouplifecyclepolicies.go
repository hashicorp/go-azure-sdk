package grouplifecyclepolicy

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

type ListGroupLifecyclePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupLifecyclePolicy
}

type ListGroupLifecyclePoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupLifecyclePolicy
}

type ListGroupLifecyclePoliciesOperationOptions struct {
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

func DefaultListGroupLifecyclePoliciesOperationOptions() ListGroupLifecyclePoliciesOperationOptions {
	return ListGroupLifecyclePoliciesOperationOptions{}
}

func (o ListGroupLifecyclePoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGroupLifecyclePoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListGroupLifecyclePoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGroupLifecyclePoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupLifecyclePoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupLifecyclePolicies - List groupLifecyclePolicies. Retrieves a list of groupLifecyclePolicy objects to which a
// group belongs.
func (c GroupLifecyclePolicyClient) ListGroupLifecyclePolicies(ctx context.Context, id beta.GroupId, options ListGroupLifecyclePoliciesOperationOptions) (result ListGroupLifecyclePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListGroupLifecyclePoliciesCustomPager{},
		Path:          fmt.Sprintf("%s/groupLifecyclePolicies", id.ID()),
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
		Values *[]beta.GroupLifecyclePolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupLifecyclePoliciesComplete retrieves all the results into a single object
func (c GroupLifecyclePolicyClient) ListGroupLifecyclePoliciesComplete(ctx context.Context, id beta.GroupId, options ListGroupLifecyclePoliciesOperationOptions) (ListGroupLifecyclePoliciesCompleteResult, error) {
	return c.ListGroupLifecyclePoliciesCompleteMatchingPredicate(ctx, id, options, GroupLifecyclePolicyOperationPredicate{})
}

// ListGroupLifecyclePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupLifecyclePolicyClient) ListGroupLifecyclePoliciesCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListGroupLifecyclePoliciesOperationOptions, predicate GroupLifecyclePolicyOperationPredicate) (result ListGroupLifecyclePoliciesCompleteResult, err error) {
	items := make([]beta.GroupLifecyclePolicy, 0)

	resp, err := c.ListGroupLifecyclePolicies(ctx, id, options)
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

	result = ListGroupLifecyclePoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
