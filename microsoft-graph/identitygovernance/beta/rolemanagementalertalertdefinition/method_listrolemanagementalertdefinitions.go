package rolemanagementalertalertdefinition

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

type ListRoleManagementAlertDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleManagementAlertDefinition
}

type ListRoleManagementAlertDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleManagementAlertDefinition
}

type ListRoleManagementAlertDefinitionsOperationOptions struct {
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

func DefaultListRoleManagementAlertDefinitionsOperationOptions() ListRoleManagementAlertDefinitionsOperationOptions {
	return ListRoleManagementAlertDefinitionsOperationOptions{}
}

func (o ListRoleManagementAlertDefinitionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRoleManagementAlertDefinitionsOperationOptions) ToOData() *odata.Query {
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

func (o ListRoleManagementAlertDefinitionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRoleManagementAlertDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleManagementAlertDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleManagementAlertDefinitions - List alertDefinitions. Get a list of the unifiedRoleManagementAlertDefinition
// objects and their properties.
func (c RoleManagementAlertAlertDefinitionClient) ListRoleManagementAlertDefinitions(ctx context.Context, options ListRoleManagementAlertDefinitionsOperationOptions) (result ListRoleManagementAlertDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRoleManagementAlertDefinitionsCustomPager{},
		Path:          "/identityGovernance/roleManagementAlerts/alertDefinitions",
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
		Values *[]beta.UnifiedRoleManagementAlertDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementAlertDefinitionsComplete retrieves all the results into a single object
func (c RoleManagementAlertAlertDefinitionClient) ListRoleManagementAlertDefinitionsComplete(ctx context.Context, options ListRoleManagementAlertDefinitionsOperationOptions) (ListRoleManagementAlertDefinitionsCompleteResult, error) {
	return c.ListRoleManagementAlertDefinitionsCompleteMatchingPredicate(ctx, options, UnifiedRoleManagementAlertDefinitionOperationPredicate{})
}

// ListRoleManagementAlertDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementAlertAlertDefinitionClient) ListRoleManagementAlertDefinitionsCompleteMatchingPredicate(ctx context.Context, options ListRoleManagementAlertDefinitionsOperationOptions, predicate UnifiedRoleManagementAlertDefinitionOperationPredicate) (result ListRoleManagementAlertDefinitionsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleManagementAlertDefinition, 0)

	resp, err := c.ListRoleManagementAlertDefinitions(ctx, options)
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

	result = ListRoleManagementAlertDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
