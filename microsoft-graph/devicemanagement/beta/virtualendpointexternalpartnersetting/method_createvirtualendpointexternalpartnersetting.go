package virtualendpointexternalpartnersetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointExternalPartnerSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCExternalPartnerSetting
}

type CreateVirtualEndpointExternalPartnerSettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateVirtualEndpointExternalPartnerSettingOperationOptions() CreateVirtualEndpointExternalPartnerSettingOperationOptions {
	return CreateVirtualEndpointExternalPartnerSettingOperationOptions{}
}

func (o CreateVirtualEndpointExternalPartnerSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointExternalPartnerSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointExternalPartnerSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointExternalPartnerSetting - Create cloudPcExternalPartnerSetting. Create a new
// cloudPcExternalPartnerSetting object.
func (c VirtualEndpointExternalPartnerSettingClient) CreateVirtualEndpointExternalPartnerSetting(ctx context.Context, input beta.CloudPCExternalPartnerSetting, options CreateVirtualEndpointExternalPartnerSettingOperationOptions) (result CreateVirtualEndpointExternalPartnerSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/externalPartnerSettings",
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

	var model beta.CloudPCExternalPartnerSetting
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
