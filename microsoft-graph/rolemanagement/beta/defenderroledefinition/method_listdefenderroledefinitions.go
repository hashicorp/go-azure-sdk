package defenderroledefinition

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

type ListDefenderRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleDefinition
}

type ListDefenderRoleDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleDefinition
}

type ListDefenderRoleDefinitionsOperationOptions struct {
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

func DefaultListDefenderRoleDefinitionsOperationOptions() ListDefenderRoleDefinitionsOperationOptions {
	return ListDefenderRoleDefinitionsOperationOptions{}
}

func (o ListDefenderRoleDefinitionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDefenderRoleDefinitionsOperationOptions) ToOData() *odata.Query {
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

func (o ListDefenderRoleDefinitionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDefenderRoleDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDefenderRoleDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDefenderRoleDefinitions - Get roleDefinitions from roleManagement
func (c DefenderRoleDefinitionClient) ListDefenderRoleDefinitions(ctx context.Context, options ListDefenderRoleDefinitionsOperationOptions) (result ListDefenderRoleDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDefenderRoleDefinitionsCustomPager{},
		Path:          "/roleManagement/defender/roleDefinitions",
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
		Values *[]beta.UnifiedRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDefenderRoleDefinitionsComplete retrieves all the results into a single object
func (c DefenderRoleDefinitionClient) ListDefenderRoleDefinitionsComplete(ctx context.Context, options ListDefenderRoleDefinitionsOperationOptions) (ListDefenderRoleDefinitionsCompleteResult, error) {
	return c.ListDefenderRoleDefinitionsCompleteMatchingPredicate(ctx, options, UnifiedRoleDefinitionOperationPredicate{})
}

// ListDefenderRoleDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DefenderRoleDefinitionClient) ListDefenderRoleDefinitionsCompleteMatchingPredicate(ctx context.Context, options ListDefenderRoleDefinitionsOperationOptions, predicate UnifiedRoleDefinitionOperationPredicate) (result ListDefenderRoleDefinitionsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleDefinition, 0)

	resp, err := c.ListDefenderRoleDefinitions(ctx, options)
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

	result = ListDefenderRoleDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
