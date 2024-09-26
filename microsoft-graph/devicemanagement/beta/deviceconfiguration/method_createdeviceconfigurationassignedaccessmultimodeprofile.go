package deviceconfiguration

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

type CreateDeviceConfigurationAssignedAccessMultiModeProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateDeviceConfigurationAssignedAccessMultiModeProfileOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDeviceConfigurationAssignedAccessMultiModeProfileOperationOptions() CreateDeviceConfigurationAssignedAccessMultiModeProfileOperationOptions {
	return CreateDeviceConfigurationAssignedAccessMultiModeProfileOperationOptions{}
}

func (o CreateDeviceConfigurationAssignedAccessMultiModeProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceConfigurationAssignedAccessMultiModeProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceConfigurationAssignedAccessMultiModeProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceConfigurationAssignedAccessMultiModeProfile - Invoke action assignedAccessMultiModeProfiles
func (c DeviceConfigurationClient) CreateDeviceConfigurationAssignedAccessMultiModeProfile(ctx context.Context, id beta.DeviceManagementDeviceConfigurationId, input CreateDeviceConfigurationAssignedAccessMultiModeProfileRequest, options CreateDeviceConfigurationAssignedAccessMultiModeProfileOperationOptions) (result CreateDeviceConfigurationAssignedAccessMultiModeProfileOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/assignedAccessMultiModeProfiles", id.ID()),
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
