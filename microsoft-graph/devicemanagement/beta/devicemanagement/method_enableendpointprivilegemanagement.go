package devicemanagement

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnableEndpointPrivilegeManagementOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type EnableEndpointPrivilegeManagementOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultEnableEndpointPrivilegeManagementOperationOptions() EnableEndpointPrivilegeManagementOperationOptions {
	return EnableEndpointPrivilegeManagementOperationOptions{}
}

func (o EnableEndpointPrivilegeManagementOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EnableEndpointPrivilegeManagementOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o EnableEndpointPrivilegeManagementOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// EnableEndpointPrivilegeManagement - Invoke action enableEndpointPrivilegeManagement. DEPRECATED - DO NOT USE.
// (Triggers onboarding of tenant to Microsoft Managed Platform - Cloud (MMP-C)).
func (c DeviceManagementClient) EnableEndpointPrivilegeManagement(ctx context.Context, options EnableEndpointPrivilegeManagementOperationOptions) (result EnableEndpointPrivilegeManagementOperationResponse, err error) {
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
		Path:          "/deviceManagement/enableEndpointPrivilegeManagement",
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

	return
}
