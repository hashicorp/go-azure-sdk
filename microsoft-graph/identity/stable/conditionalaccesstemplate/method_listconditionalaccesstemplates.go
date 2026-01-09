package conditionalaccesstemplate

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListConditionalAccessTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConditionalAccessTemplate
}

type ListConditionalAccessTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConditionalAccessTemplate
}

type ListConditionalAccessTemplatesOperationOptions struct {
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

func DefaultListConditionalAccessTemplatesOperationOptions() ListConditionalAccessTemplatesOperationOptions {
	return ListConditionalAccessTemplatesOperationOptions{}
}

func (o ListConditionalAccessTemplatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListConditionalAccessTemplatesOperationOptions) ToOData() *odata.Query {
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

func (o ListConditionalAccessTemplatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListConditionalAccessTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccessTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccessTemplates - List conditionalAccessTemplates. Get a list of the conditionalAccessTemplate objects
// and their properties.
func (c ConditionalAccessTemplateClient) ListConditionalAccessTemplates(ctx context.Context, options ListConditionalAccessTemplatesOperationOptions) (result ListConditionalAccessTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListConditionalAccessTemplatesCustomPager{},
		Path:          "/identity/conditionalAccess/templates",
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
		Values *[]stable.ConditionalAccessTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConditionalAccessTemplatesComplete retrieves all the results into a single object
func (c ConditionalAccessTemplateClient) ListConditionalAccessTemplatesComplete(ctx context.Context, options ListConditionalAccessTemplatesOperationOptions) (ListConditionalAccessTemplatesCompleteResult, error) {
	return c.ListConditionalAccessTemplatesCompleteMatchingPredicate(ctx, options, ConditionalAccessTemplateOperationPredicate{})
}

// ListConditionalAccessTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccessTemplateClient) ListConditionalAccessTemplatesCompleteMatchingPredicate(ctx context.Context, options ListConditionalAccessTemplatesOperationOptions, predicate ConditionalAccessTemplateOperationPredicate) (result ListConditionalAccessTemplatesCompleteResult, err error) {
	items := make([]stable.ConditionalAccessTemplate, 0)

	resp, err := c.ListConditionalAccessTemplates(ctx, options)
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

	result = ListConditionalAccessTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
