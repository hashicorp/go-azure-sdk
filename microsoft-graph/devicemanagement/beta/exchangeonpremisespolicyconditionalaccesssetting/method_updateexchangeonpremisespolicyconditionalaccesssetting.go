package exchangeonpremisespolicyconditionalaccesssetting

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

type UpdateExchangeOnPremisesPolicyConditionalAccessSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions() UpdateExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions {
	return UpdateExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions{}
}

func (o UpdateExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateExchangeOnPremisesPolicyConditionalAccessSetting - Update the navigation property conditionalAccessSettings in
// deviceManagement
func (c ExchangeOnPremisesPolicyConditionalAccessSettingClient) UpdateExchangeOnPremisesPolicyConditionalAccessSetting(ctx context.Context, id beta.DeviceManagementExchangeOnPremisesPolicyId, input beta.OnPremisesConditionalAccessSettings, options UpdateExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions) (result UpdateExchangeOnPremisesPolicyConditionalAccessSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/conditionalAccessSettings", id.ID()),
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

	return
}
