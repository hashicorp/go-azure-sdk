package grouppolicymigrationreportgrouppolicysettingmapping

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

type ListGroupPolicyMigrationReportSettingMappingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicySettingMapping
}

type ListGroupPolicyMigrationReportSettingMappingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicySettingMapping
}

type ListGroupPolicyMigrationReportSettingMappingsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListGroupPolicyMigrationReportSettingMappingsOperationOptions() ListGroupPolicyMigrationReportSettingMappingsOperationOptions {
	return ListGroupPolicyMigrationReportSettingMappingsOperationOptions{}
}

func (o ListGroupPolicyMigrationReportSettingMappingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGroupPolicyMigrationReportSettingMappingsOperationOptions) ToOData() *odata.Query {
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

func (o ListGroupPolicyMigrationReportSettingMappingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGroupPolicyMigrationReportSettingMappingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyMigrationReportSettingMappingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyMigrationReportSettingMappings - Get groupPolicySettingMappings from deviceManagement. A list of group
// policy settings to MDM/Intune mappings.
func (c GroupPolicyMigrationReportGroupPolicySettingMappingClient) ListGroupPolicyMigrationReportSettingMappings(ctx context.Context, id beta.DeviceManagementGroupPolicyMigrationReportId, options ListGroupPolicyMigrationReportSettingMappingsOperationOptions) (result ListGroupPolicyMigrationReportSettingMappingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListGroupPolicyMigrationReportSettingMappingsCustomPager{},
		Path:          fmt.Sprintf("%s/groupPolicySettingMappings", id.ID()),
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
		Values *[]beta.GroupPolicySettingMapping `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyMigrationReportSettingMappingsComplete retrieves all the results into a single object
func (c GroupPolicyMigrationReportGroupPolicySettingMappingClient) ListGroupPolicyMigrationReportSettingMappingsComplete(ctx context.Context, id beta.DeviceManagementGroupPolicyMigrationReportId, options ListGroupPolicyMigrationReportSettingMappingsOperationOptions) (ListGroupPolicyMigrationReportSettingMappingsCompleteResult, error) {
	return c.ListGroupPolicyMigrationReportSettingMappingsCompleteMatchingPredicate(ctx, id, options, GroupPolicySettingMappingOperationPredicate{})
}

// ListGroupPolicyMigrationReportSettingMappingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyMigrationReportGroupPolicySettingMappingClient) ListGroupPolicyMigrationReportSettingMappingsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementGroupPolicyMigrationReportId, options ListGroupPolicyMigrationReportSettingMappingsOperationOptions, predicate GroupPolicySettingMappingOperationPredicate) (result ListGroupPolicyMigrationReportSettingMappingsCompleteResult, err error) {
	items := make([]beta.GroupPolicySettingMapping, 0)

	resp, err := c.ListGroupPolicyMigrationReportSettingMappings(ctx, id, options)
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

	result = ListGroupPolicyMigrationReportSettingMappingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
