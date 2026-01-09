package compliancepolicysetting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListCompliancePolicySettingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationSetting
}

type ListCompliancePolicySettingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationSetting
}

type ListCompliancePolicySettingsOperationOptions struct {
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

func DefaultListCompliancePolicySettingsOperationOptions() ListCompliancePolicySettingsOperationOptions {
	return ListCompliancePolicySettingsOperationOptions{}
}

func (o ListCompliancePolicySettingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCompliancePolicySettingsOperationOptions) ToOData() *odata.Query {
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

func (o ListCompliancePolicySettingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCompliancePolicySettingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCompliancePolicySettingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCompliancePolicySettings - Get settings from deviceManagement. Policy settings
func (c CompliancePolicySettingClient) ListCompliancePolicySettings(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, options ListCompliancePolicySettingsOperationOptions) (result ListCompliancePolicySettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCompliancePolicySettingsCustomPager{},
		Path:          fmt.Sprintf("%s/settings", id.ID()),
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
		Values *[]beta.DeviceManagementConfigurationSetting `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCompliancePolicySettingsComplete retrieves all the results into a single object
func (c CompliancePolicySettingClient) ListCompliancePolicySettingsComplete(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, options ListCompliancePolicySettingsOperationOptions) (ListCompliancePolicySettingsCompleteResult, error) {
	return c.ListCompliancePolicySettingsCompleteMatchingPredicate(ctx, id, options, DeviceManagementConfigurationSettingOperationPredicate{})
}

// ListCompliancePolicySettingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicySettingClient) ListCompliancePolicySettingsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, options ListCompliancePolicySettingsOperationOptions, predicate DeviceManagementConfigurationSettingOperationPredicate) (result ListCompliancePolicySettingsCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationSetting, 0)

	resp, err := c.ListCompliancePolicySettings(ctx, id, options)
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

	result = ListCompliancePolicySettingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
