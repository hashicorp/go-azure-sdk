package virtualendpointusersettingassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualEndpointUserSettingAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudPCUserSettingAssignment
}

type GetVirtualEndpointUserSettingAssignmentOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetVirtualEndpointUserSettingAssignmentOperationOptions() GetVirtualEndpointUserSettingAssignmentOperationOptions {
	return GetVirtualEndpointUserSettingAssignmentOperationOptions{}
}

func (o GetVirtualEndpointUserSettingAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVirtualEndpointUserSettingAssignmentOperationOptions) ToOData() *odata.Query {
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

func (o GetVirtualEndpointUserSettingAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVirtualEndpointUserSettingAssignment - Get assignments from deviceManagement. Represents the set of Microsoft 365
// groups and security groups in Microsoft Entra ID that have cloudPCUserSetting assigned. Returned only on $expand. For
// an example, see Get cloudPcUserSettingample.
func (c VirtualEndpointUserSettingAssignmentClient) GetVirtualEndpointUserSettingAssignment(ctx context.Context, id beta.DeviceManagementVirtualEndpointUserSettingIdAssignmentId, options GetVirtualEndpointUserSettingAssignmentOperationOptions) (result GetVirtualEndpointUserSettingAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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

	var model beta.CloudPCUserSettingAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
