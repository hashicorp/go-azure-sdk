package accessreviewdefinitioninstance

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

type ListAccessReviewDefinitionInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewInstance
}

type ListAccessReviewDefinitionInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewInstance
}

type ListAccessReviewDefinitionInstancesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListAccessReviewDefinitionInstancesOperationOptions() ListAccessReviewDefinitionInstancesOperationOptions {
	return ListAccessReviewDefinitionInstancesOperationOptions{}
}

func (o ListAccessReviewDefinitionInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAccessReviewDefinitionInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListAccessReviewDefinitionInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAccessReviewDefinitionInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitionInstances - List instances. Retrieve the accessReviewInstance objects for a specific
// accessReviewScheduleDefinition. A list of zero or more accessReviewInstance objects are returned, including all of
// their nested properties. Returned objects do not include associated accessReviewInstanceDecisionItems. To retrieve
// the decisions on the instance, use List accessReviewInstanceDecisionItem.
func (c AccessReviewDefinitionInstanceClient) ListAccessReviewDefinitionInstances(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionId, options ListAccessReviewDefinitionInstancesOperationOptions) (result ListAccessReviewDefinitionInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAccessReviewDefinitionInstancesCustomPager{},
		Path:          fmt.Sprintf("%s/instances", id.ID()),
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
		Values *[]beta.AccessReviewInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewDefinitionInstancesComplete retrieves all the results into a single object
func (c AccessReviewDefinitionInstanceClient) ListAccessReviewDefinitionInstancesComplete(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionId, options ListAccessReviewDefinitionInstancesOperationOptions) (ListAccessReviewDefinitionInstancesCompleteResult, error) {
	return c.ListAccessReviewDefinitionInstancesCompleteMatchingPredicate(ctx, id, options, AccessReviewInstanceOperationPredicate{})
}

// ListAccessReviewDefinitionInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionInstanceClient) ListAccessReviewDefinitionInstancesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionId, options ListAccessReviewDefinitionInstancesOperationOptions, predicate AccessReviewInstanceOperationPredicate) (result ListAccessReviewDefinitionInstancesCompleteResult, err error) {
	items := make([]beta.AccessReviewInstance, 0)

	resp, err := c.ListAccessReviewDefinitionInstances(ctx, id, options)
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

	result = ListAccessReviewDefinitionInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
