package configurationpolicytemplate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateConfigurationPolicyTemplateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateConfigurationPolicyTemplateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateConfigurationPolicyTemplateOperationOptions() UpdateConfigurationPolicyTemplateOperationOptions {
	return UpdateConfigurationPolicyTemplateOperationOptions{}
}

func (o UpdateConfigurationPolicyTemplateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateConfigurationPolicyTemplateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateConfigurationPolicyTemplateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateConfigurationPolicyTemplate - Update the navigation property configurationPolicyTemplates in deviceManagement
func (c ConfigurationPolicyTemplateClient) UpdateConfigurationPolicyTemplate(ctx context.Context, id beta.DeviceManagementConfigurationPolicyTemplateId, input beta.DeviceManagementConfigurationPolicyTemplate, options UpdateConfigurationPolicyTemplateOperationOptions) (result UpdateConfigurationPolicyTemplateOperationResponse, err error) {
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
