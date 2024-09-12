package grouppolicymigrationreportsettingmapping

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetGroupPolicyMigrationReportSettingMappingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.GroupPolicySettingMapping
}

type GetGroupPolicyMigrationReportSettingMappingOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetGroupPolicyMigrationReportSettingMappingOperationOptions() GetGroupPolicyMigrationReportSettingMappingOperationOptions {
	return GetGroupPolicyMigrationReportSettingMappingOperationOptions{}
}

func (o GetGroupPolicyMigrationReportSettingMappingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetGroupPolicyMigrationReportSettingMappingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetGroupPolicyMigrationReportSettingMappingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetGroupPolicyMigrationReportSettingMapping - Get groupPolicySettingMappings from deviceManagement. A list of group
// policy settings to MDM/Intune mappings.
func (c GroupPolicyMigrationReportSettingMappingClient) GetGroupPolicyMigrationReportSettingMapping(ctx context.Context, id beta.DeviceManagementGroupPolicyMigrationReportIdGroupPolicySettingMappingId, options GetGroupPolicyMigrationReportSettingMappingOperationOptions) (result GetGroupPolicyMigrationReportSettingMappingOperationResponse, err error) {
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

	var model beta.GroupPolicySettingMapping
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
