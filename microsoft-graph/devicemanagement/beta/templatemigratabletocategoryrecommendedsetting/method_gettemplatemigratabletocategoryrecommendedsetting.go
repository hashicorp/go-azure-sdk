package templatemigratabletocategoryrecommendedsetting

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTemplateMigratableToCategoryRecommendedSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DeviceManagementSettingInstance
}

type GetTemplateMigratableToCategoryRecommendedSettingOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetTemplateMigratableToCategoryRecommendedSettingOperationOptions() GetTemplateMigratableToCategoryRecommendedSettingOperationOptions {
	return GetTemplateMigratableToCategoryRecommendedSettingOperationOptions{}
}

func (o GetTemplateMigratableToCategoryRecommendedSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTemplateMigratableToCategoryRecommendedSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTemplateMigratableToCategoryRecommendedSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTemplateMigratableToCategoryRecommendedSetting - Get recommendedSettings from deviceManagement. The settings this
// category contains
func (c TemplateMigratableToCategoryRecommendedSettingClient) GetTemplateMigratableToCategoryRecommendedSetting(ctx context.Context, id beta.DeviceManagementTemplateIdMigratableToIdCategoryIdRecommendedSettingId, options GetTemplateMigratableToCategoryRecommendedSettingOperationOptions) (result GetTemplateMigratableToCategoryRecommendedSettingOperationResponse, err error) {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalDeviceManagementSettingInstanceImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
