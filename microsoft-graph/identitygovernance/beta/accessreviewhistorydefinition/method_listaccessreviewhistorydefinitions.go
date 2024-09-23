package accessreviewhistorydefinition

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

type ListAccessReviewHistoryDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewHistoryDefinition
}

type ListAccessReviewHistoryDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewHistoryDefinition
}

type ListAccessReviewHistoryDefinitionsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListAccessReviewHistoryDefinitionsOperationOptions() ListAccessReviewHistoryDefinitionsOperationOptions {
	return ListAccessReviewHistoryDefinitionsOperationOptions{}
}

func (o ListAccessReviewHistoryDefinitionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAccessReviewHistoryDefinitionsOperationOptions) ToOData() *odata.Query {
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

func (o ListAccessReviewHistoryDefinitionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAccessReviewHistoryDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewHistoryDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewHistoryDefinitions - List historyDefinitions. Retrieve the accessReviewHistoryDefinition objects
// created in the last 30 days, including all nested properties.
func (c AccessReviewHistoryDefinitionClient) ListAccessReviewHistoryDefinitions(ctx context.Context, options ListAccessReviewHistoryDefinitionsOperationOptions) (result ListAccessReviewHistoryDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAccessReviewHistoryDefinitionsCustomPager{},
		Path:          "/identityGovernance/accessReviews/historyDefinitions",
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
		Values *[]beta.AccessReviewHistoryDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewHistoryDefinitionsComplete retrieves all the results into a single object
func (c AccessReviewHistoryDefinitionClient) ListAccessReviewHistoryDefinitionsComplete(ctx context.Context, options ListAccessReviewHistoryDefinitionsOperationOptions) (ListAccessReviewHistoryDefinitionsCompleteResult, error) {
	return c.ListAccessReviewHistoryDefinitionsCompleteMatchingPredicate(ctx, options, AccessReviewHistoryDefinitionOperationPredicate{})
}

// ListAccessReviewHistoryDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewHistoryDefinitionClient) ListAccessReviewHistoryDefinitionsCompleteMatchingPredicate(ctx context.Context, options ListAccessReviewHistoryDefinitionsOperationOptions, predicate AccessReviewHistoryDefinitionOperationPredicate) (result ListAccessReviewHistoryDefinitionsCompleteResult, err error) {
	items := make([]beta.AccessReviewHistoryDefinition, 0)

	resp, err := c.ListAccessReviewHistoryDefinitions(ctx, options)
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

	result = ListAccessReviewHistoryDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
