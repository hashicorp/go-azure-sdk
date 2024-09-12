package virtualendpointexternalpartnersetting

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

type DeleteVirtualEndpointExternalPartnerSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteVirtualEndpointExternalPartnerSettingOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteVirtualEndpointExternalPartnerSettingOperationOptions() DeleteVirtualEndpointExternalPartnerSettingOperationOptions {
	return DeleteVirtualEndpointExternalPartnerSettingOperationOptions{}
}

func (o DeleteVirtualEndpointExternalPartnerSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteVirtualEndpointExternalPartnerSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteVirtualEndpointExternalPartnerSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteVirtualEndpointExternalPartnerSetting - Delete navigation property externalPartnerSettings for deviceManagement
func (c VirtualEndpointExternalPartnerSettingClient) DeleteVirtualEndpointExternalPartnerSetting(ctx context.Context, id beta.DeviceManagementVirtualEndpointExternalPartnerSettingId, options DeleteVirtualEndpointExternalPartnerSettingOperationOptions) (result DeleteVirtualEndpointExternalPartnerSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
