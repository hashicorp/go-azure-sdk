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

type GetExchangeOnPremisesPolicyConditionalAccessSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.OnPremisesConditionalAccessSettings
}

type GetExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions() GetExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions {
	return GetExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions{}
}

func (o GetExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetExchangeOnPremisesPolicyConditionalAccessSetting - Get conditionalAccessSettings from deviceManagement. The
// Exchange on premises conditional access settings. On premises conditional access will require devices to be both
// enrolled and compliant for mail access
func (c ExchangeOnPremisesPolicyConditionalAccessSettingClient) GetExchangeOnPremisesPolicyConditionalAccessSetting(ctx context.Context, id beta.DeviceManagementExchangeOnPremisesPolicyId, options GetExchangeOnPremisesPolicyConditionalAccessSettingOperationOptions) (result GetExchangeOnPremisesPolicyConditionalAccessSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/conditionalAccessSettings", id.ID()),
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

	var model beta.OnPremisesConditionalAccessSettings
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
