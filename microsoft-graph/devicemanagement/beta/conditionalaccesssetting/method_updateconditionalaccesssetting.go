package conditionalaccesssetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateConditionalAccessSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateConditionalAccessSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateConditionalAccessSettingOperationOptions() UpdateConditionalAccessSettingOperationOptions {
	return UpdateConditionalAccessSettingOperationOptions{}
}

func (o UpdateConditionalAccessSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateConditionalAccessSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateConditionalAccessSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateConditionalAccessSetting - Update the navigation property conditionalAccessSettings in deviceManagement
func (c ConditionalAccessSettingClient) UpdateConditionalAccessSetting(ctx context.Context, input beta.OnPremisesConditionalAccessSettings, options UpdateConditionalAccessSettingOperationOptions) (result UpdateConditionalAccessSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/deviceManagement/conditionalAccessSettings",
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
