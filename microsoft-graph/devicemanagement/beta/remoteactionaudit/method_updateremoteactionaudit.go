package remoteactionaudit

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRemoteActionAuditOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateRemoteActionAuditOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateRemoteActionAuditOperationOptions() UpdateRemoteActionAuditOperationOptions {
	return UpdateRemoteActionAuditOperationOptions{}
}

func (o UpdateRemoteActionAuditOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateRemoteActionAuditOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateRemoteActionAuditOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateRemoteActionAudit - Update the navigation property remoteActionAudits in deviceManagement
func (c RemoteActionAuditClient) UpdateRemoteActionAudit(ctx context.Context, id beta.DeviceManagementRemoteActionAuditId, input beta.RemoteActionAudit, options UpdateRemoteActionAuditOperationOptions) (result UpdateRemoteActionAuditOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
