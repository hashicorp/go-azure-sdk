package templatecategoryrecommendedsetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTemplateCategoryRecommendedSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTemplateCategoryRecommendedSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateTemplateCategoryRecommendedSettingOperationOptions() UpdateTemplateCategoryRecommendedSettingOperationOptions {
	return UpdateTemplateCategoryRecommendedSettingOperationOptions{}
}

func (o UpdateTemplateCategoryRecommendedSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTemplateCategoryRecommendedSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTemplateCategoryRecommendedSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTemplateCategoryRecommendedSetting - Update the navigation property recommendedSettings in deviceManagement
func (c TemplateCategoryRecommendedSettingClient) UpdateTemplateCategoryRecommendedSetting(ctx context.Context, id beta.DeviceManagementTemplateIdCategoryIdRecommendedSettingId, input beta.DeviceManagementSettingInstance, options UpdateTemplateCategoryRecommendedSettingOperationOptions) (result UpdateTemplateCategoryRecommendedSettingOperationResponse, err error) {
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
