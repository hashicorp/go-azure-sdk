package configurationpolicytemplatesettingtemplate

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

type ListConfigurationPolicyTemplateSettingTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationSettingTemplate
}

type ListConfigurationPolicyTemplateSettingTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationSettingTemplate
}

type ListConfigurationPolicyTemplateSettingTemplatesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListConfigurationPolicyTemplateSettingTemplatesOperationOptions() ListConfigurationPolicyTemplateSettingTemplatesOperationOptions {
	return ListConfigurationPolicyTemplateSettingTemplatesOperationOptions{}
}

func (o ListConfigurationPolicyTemplateSettingTemplatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListConfigurationPolicyTemplateSettingTemplatesOperationOptions) ToOData() *odata.Query {
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

func (o ListConfigurationPolicyTemplateSettingTemplatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListConfigurationPolicyTemplateSettingTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConfigurationPolicyTemplateSettingTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigurationPolicyTemplateSettingTemplates - Get settingTemplates from deviceManagement. Setting templates
func (c ConfigurationPolicyTemplateSettingTemplateClient) ListConfigurationPolicyTemplateSettingTemplates(ctx context.Context, id beta.DeviceManagementConfigurationPolicyTemplateId, options ListConfigurationPolicyTemplateSettingTemplatesOperationOptions) (result ListConfigurationPolicyTemplateSettingTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListConfigurationPolicyTemplateSettingTemplatesCustomPager{},
		Path:          fmt.Sprintf("%s/settingTemplates", id.ID()),
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
		Values *[]beta.DeviceManagementConfigurationSettingTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConfigurationPolicyTemplateSettingTemplatesComplete retrieves all the results into a single object
func (c ConfigurationPolicyTemplateSettingTemplateClient) ListConfigurationPolicyTemplateSettingTemplatesComplete(ctx context.Context, id beta.DeviceManagementConfigurationPolicyTemplateId, options ListConfigurationPolicyTemplateSettingTemplatesOperationOptions) (ListConfigurationPolicyTemplateSettingTemplatesCompleteResult, error) {
	return c.ListConfigurationPolicyTemplateSettingTemplatesCompleteMatchingPredicate(ctx, id, options, DeviceManagementConfigurationSettingTemplateOperationPredicate{})
}

// ListConfigurationPolicyTemplateSettingTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigurationPolicyTemplateSettingTemplateClient) ListConfigurationPolicyTemplateSettingTemplatesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementConfigurationPolicyTemplateId, options ListConfigurationPolicyTemplateSettingTemplatesOperationOptions, predicate DeviceManagementConfigurationSettingTemplateOperationPredicate) (result ListConfigurationPolicyTemplateSettingTemplatesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationSettingTemplate, 0)

	resp, err := c.ListConfigurationPolicyTemplateSettingTemplates(ctx, id, options)
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

	result = ListConfigurationPolicyTemplateSettingTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
