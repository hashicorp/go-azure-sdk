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

type CreateRemoteActionAuditOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.RemoteActionAudit
}

type CreateRemoteActionAuditOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateRemoteActionAuditOperationOptions() CreateRemoteActionAuditOperationOptions {
	return CreateRemoteActionAuditOperationOptions{}
}

func (o CreateRemoteActionAuditOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateRemoteActionAuditOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateRemoteActionAuditOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateRemoteActionAudit - Create new navigation property to remoteActionAudits for deviceManagement
func (c RemoteActionAuditClient) CreateRemoteActionAudit(ctx context.Context, input beta.RemoteActionAudit, options CreateRemoteActionAuditOperationOptions) (result CreateRemoteActionAuditOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/remoteActionAudits",
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

	var model beta.RemoteActionAudit
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
