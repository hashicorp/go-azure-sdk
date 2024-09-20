package devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration

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

type ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceComplianceActionItem
}

type ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceComplianceActionItem
}

type ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationOptions struct {
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

func DefaultListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationOptions() ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationOptions {
	return ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationOptions{}
}

func (o ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicyScheduledActionsForRuleConfigurations - List deviceComplianceActionItems. List properties
// and relationships of the deviceComplianceActionItem objects.
func (c DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) ListDeviceCompliancePolicyScheduledActionsForRuleConfigurations(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId, options ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationOptions) (result ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsCustomPager{},
		Path:          fmt.Sprintf("%s/scheduledActionConfigurations", id.ID()),
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
		Values *[]stable.DeviceComplianceActionItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsComplete(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId, options ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationOptions) (ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsCompleteResult, error) {
	return c.ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsCompleteMatchingPredicate(ctx, id, options, DeviceComplianceActionItemOperationPredicate{})
}

// ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId, options ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsOperationOptions, predicate DeviceComplianceActionItemOperationPredicate) (result ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsCompleteResult, err error) {
	items := make([]stable.DeviceComplianceActionItem, 0)

	resp, err := c.ListDeviceCompliancePolicyScheduledActionsForRuleConfigurations(ctx, id, options)
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

	result = ListDeviceCompliancePolicyScheduledActionsForRuleConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
