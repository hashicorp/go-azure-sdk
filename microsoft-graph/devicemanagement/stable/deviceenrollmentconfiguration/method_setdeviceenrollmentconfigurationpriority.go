package deviceenrollmentconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetDeviceEnrollmentConfigurationPriorityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDeviceEnrollmentConfigurationPriorityOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetDeviceEnrollmentConfigurationPriorityOperationOptions() SetDeviceEnrollmentConfigurationPriorityOperationOptions {
	return SetDeviceEnrollmentConfigurationPriorityOperationOptions{}
}

func (o SetDeviceEnrollmentConfigurationPriorityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDeviceEnrollmentConfigurationPriorityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDeviceEnrollmentConfigurationPriorityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDeviceEnrollmentConfigurationPriority - Invoke action setPriority. Not yet documented
func (c DeviceEnrollmentConfigurationClient) SetDeviceEnrollmentConfigurationPriority(ctx context.Context, id stable.DeviceManagementDeviceEnrollmentConfigurationId, input SetDeviceEnrollmentConfigurationPriorityRequest, options SetDeviceEnrollmentConfigurationPriorityOperationOptions) (result SetDeviceEnrollmentConfigurationPriorityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/setPriority", id.ID()),
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
