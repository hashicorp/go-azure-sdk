package intentcategory

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateIntentCategoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateIntentCategoryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateIntentCategoryOperationOptions() UpdateIntentCategoryOperationOptions {
	return UpdateIntentCategoryOperationOptions{}
}

func (o UpdateIntentCategoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateIntentCategoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateIntentCategoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateIntentCategory - Update the navigation property categories in deviceManagement
func (c IntentCategoryClient) UpdateIntentCategory(ctx context.Context, id beta.DeviceManagementIntentIdCategoryId, input beta.DeviceManagementIntentSettingCategory, options UpdateIntentCategoryOperationOptions) (result UpdateIntentCategoryOperationResponse, err error) {
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
