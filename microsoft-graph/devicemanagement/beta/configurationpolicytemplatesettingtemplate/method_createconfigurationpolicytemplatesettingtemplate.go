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

type CreateConfigurationPolicyTemplateSettingTemplateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceManagementConfigurationSettingTemplate
}

type CreateConfigurationPolicyTemplateSettingTemplateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateConfigurationPolicyTemplateSettingTemplateOperationOptions() CreateConfigurationPolicyTemplateSettingTemplateOperationOptions {
	return CreateConfigurationPolicyTemplateSettingTemplateOperationOptions{}
}

func (o CreateConfigurationPolicyTemplateSettingTemplateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateConfigurationPolicyTemplateSettingTemplateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateConfigurationPolicyTemplateSettingTemplateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateConfigurationPolicyTemplateSettingTemplate - Create new navigation property to settingTemplates for
// deviceManagement
func (c ConfigurationPolicyTemplateSettingTemplateClient) CreateConfigurationPolicyTemplateSettingTemplate(ctx context.Context, id beta.DeviceManagementConfigurationPolicyTemplateId, input beta.DeviceManagementConfigurationSettingTemplate, options CreateConfigurationPolicyTemplateSettingTemplateOperationOptions) (result CreateConfigurationPolicyTemplateSettingTemplateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/settingTemplates", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
