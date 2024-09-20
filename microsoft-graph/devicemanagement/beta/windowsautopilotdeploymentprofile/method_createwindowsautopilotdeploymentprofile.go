package windowsautopilotdeploymentprofile

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsAutopilotDeploymentProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.WindowsAutopilotDeploymentProfile
}

type CreateWindowsAutopilotDeploymentProfileOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateWindowsAutopilotDeploymentProfileOperationOptions() CreateWindowsAutopilotDeploymentProfileOperationOptions {
	return CreateWindowsAutopilotDeploymentProfileOperationOptions{}
}

func (o CreateWindowsAutopilotDeploymentProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsAutopilotDeploymentProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsAutopilotDeploymentProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsAutopilotDeploymentProfile - Create new navigation property to windowsAutopilotDeploymentProfiles for
// deviceManagement
func (c WindowsAutopilotDeploymentProfileClient) CreateWindowsAutopilotDeploymentProfile(ctx context.Context, input beta.WindowsAutopilotDeploymentProfile, options CreateWindowsAutopilotDeploymentProfileOperationOptions) (result CreateWindowsAutopilotDeploymentProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/windowsAutopilotDeploymentProfiles",
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalWindowsAutopilotDeploymentProfileImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
