package templatemigratabletocategorysettingdefinition

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTemplateMigratableToCategorySettingDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTemplateMigratableToCategorySettingDefinitionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateTemplateMigratableToCategorySettingDefinitionOperationOptions() UpdateTemplateMigratableToCategorySettingDefinitionOperationOptions {
	return UpdateTemplateMigratableToCategorySettingDefinitionOperationOptions{}
}

func (o UpdateTemplateMigratableToCategorySettingDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTemplateMigratableToCategorySettingDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTemplateMigratableToCategorySettingDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTemplateMigratableToCategorySettingDefinition - Update the navigation property settingDefinitions in
// deviceManagement
func (c TemplateMigratableToCategorySettingDefinitionClient) UpdateTemplateMigratableToCategorySettingDefinition(ctx context.Context, id beta.DeviceManagementTemplateIdMigratableToIdCategoryIdSettingDefinitionId, input beta.DeviceManagementSettingDefinition, options UpdateTemplateMigratableToCategorySettingDefinitionOperationOptions) (result UpdateTemplateMigratableToCategorySettingDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
