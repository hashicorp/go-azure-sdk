package accessreviewdefinition

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

type ListAccessReviewDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewScheduleDefinition
}

type ListAccessReviewDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewScheduleDefinition
}

type ListAccessReviewDefinitionsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListAccessReviewDefinitionsOperationOptions() ListAccessReviewDefinitionsOperationOptions {
	return ListAccessReviewDefinitionsOperationOptions{}
}

func (o ListAccessReviewDefinitionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAccessReviewDefinitionsOperationOptions) ToOData() *odata.Query {
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

func (o ListAccessReviewDefinitionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAccessReviewDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitions - List definitions. Retrieve the accessReviewScheduleDefinition objects. A list of zero
// or more accessReviewScheduleDefinition objects are returned, including all of their nested properties, for each
// access review series created. This does not include the associated accessReviewInstance objects.
func (c AccessReviewDefinitionClient) ListAccessReviewDefinitions(ctx context.Context, options ListAccessReviewDefinitionsOperationOptions) (result ListAccessReviewDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAccessReviewDefinitionsCustomPager{},
		Path:          "/identityGovernance/accessReviews/definitions",
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
		Values *[]beta.AccessReviewScheduleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewDefinitionsComplete retrieves all the results into a single object
func (c AccessReviewDefinitionClient) ListAccessReviewDefinitionsComplete(ctx context.Context, options ListAccessReviewDefinitionsOperationOptions) (ListAccessReviewDefinitionsCompleteResult, error) {
	return c.ListAccessReviewDefinitionsCompleteMatchingPredicate(ctx, options, AccessReviewScheduleDefinitionOperationPredicate{})
}

// ListAccessReviewDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionClient) ListAccessReviewDefinitionsCompleteMatchingPredicate(ctx context.Context, options ListAccessReviewDefinitionsOperationOptions, predicate AccessReviewScheduleDefinitionOperationPredicate) (result ListAccessReviewDefinitionsCompleteResult, err error) {
	items := make([]beta.AccessReviewScheduleDefinition, 0)

	resp, err := c.ListAccessReviewDefinitions(ctx, options)
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

	result = ListAccessReviewDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
