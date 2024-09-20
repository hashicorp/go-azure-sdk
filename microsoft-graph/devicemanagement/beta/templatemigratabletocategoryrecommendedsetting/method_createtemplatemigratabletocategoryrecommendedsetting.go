package templatemigratabletocategoryrecommendedsetting

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTemplateMigratableToCategoryRecommendedSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DeviceManagementSettingInstance
}

type CreateTemplateMigratableToCategoryRecommendedSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateTemplateMigratableToCategoryRecommendedSettingOperationOptions() CreateTemplateMigratableToCategoryRecommendedSettingOperationOptions {
	return CreateTemplateMigratableToCategoryRecommendedSettingOperationOptions{}
}

func (o CreateTemplateMigratableToCategoryRecommendedSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTemplateMigratableToCategoryRecommendedSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTemplateMigratableToCategoryRecommendedSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTemplateMigratableToCategoryRecommendedSetting - Create new navigation property to recommendedSettings for
// deviceManagement
func (c TemplateMigratableToCategoryRecommendedSettingClient) CreateTemplateMigratableToCategoryRecommendedSetting(ctx context.Context, id beta.DeviceManagementTemplateIdMigratableToIdCategoryId, input beta.DeviceManagementSettingInstance, options CreateTemplateMigratableToCategoryRecommendedSettingOperationOptions) (result CreateTemplateMigratableToCategoryRecommendedSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/recommendedSettings", id.ID()),
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
