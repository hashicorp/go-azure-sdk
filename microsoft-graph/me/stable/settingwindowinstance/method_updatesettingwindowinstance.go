package settingwindowinstance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSettingWindowInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSettingWindowInstanceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSettingWindowInstanceOperationOptions() UpdateSettingWindowInstanceOperationOptions {
	return UpdateSettingWindowInstanceOperationOptions{}
}

func (o UpdateSettingWindowInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSettingWindowInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSettingWindowInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSettingWindowInstance - Update the navigation property instances in me
func (c SettingWindowInstanceClient) UpdateSettingWindowInstance(ctx context.Context, id stable.MeSettingWindowIdInstanceId, input stable.WindowsSettingInstance, options UpdateSettingWindowInstanceOperationOptions) (result UpdateSettingWindowInstanceOperationResponse, err error) {
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
