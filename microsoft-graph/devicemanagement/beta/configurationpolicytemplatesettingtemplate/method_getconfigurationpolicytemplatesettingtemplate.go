package configurationpolicytemplatesettingtemplate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetConfigurationPolicyTemplateSettingTemplateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceManagementConfigurationSettingTemplate
}

type GetConfigurationPolicyTemplateSettingTemplateOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetConfigurationPolicyTemplateSettingTemplateOperationOptions() GetConfigurationPolicyTemplateSettingTemplateOperationOptions {
	return GetConfigurationPolicyTemplateSettingTemplateOperationOptions{}
}

func (o GetConfigurationPolicyTemplateSettingTemplateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetConfigurationPolicyTemplateSettingTemplateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetConfigurationPolicyTemplateSettingTemplateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetConfigurationPolicyTemplateSettingTemplate - Get settingTemplates from deviceManagement. Setting templates
func (c ConfigurationPolicyTemplateSettingTemplateClient) GetConfigurationPolicyTemplateSettingTemplate(ctx context.Context, id beta.DeviceManagementConfigurationPolicyTemplateIdSettingTemplateId, options GetConfigurationPolicyTemplateSettingTemplateOperationOptions) (result GetConfigurationPolicyTemplateSettingTemplateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.DeviceManagementConfigurationSettingTemplate
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
