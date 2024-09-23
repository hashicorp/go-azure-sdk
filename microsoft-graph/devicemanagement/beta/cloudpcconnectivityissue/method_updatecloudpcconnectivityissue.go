package cloudpcconnectivityissue

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCloudPCConnectivityIssueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCloudPCConnectivityIssueOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateCloudPCConnectivityIssueOperationOptions() UpdateCloudPCConnectivityIssueOperationOptions {
	return UpdateCloudPCConnectivityIssueOperationOptions{}
}

func (o UpdateCloudPCConnectivityIssueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCloudPCConnectivityIssueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCloudPCConnectivityIssueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCloudPCConnectivityIssue - Update the navigation property cloudPCConnectivityIssues in deviceManagement
func (c CloudPCConnectivityIssueClient) UpdateCloudPCConnectivityIssue(ctx context.Context, id beta.DeviceManagementCloudPCConnectivityIssueId, input beta.CloudPCConnectivityIssue, options UpdateCloudPCConnectivityIssueOperationOptions) (result UpdateCloudPCConnectivityIssueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
