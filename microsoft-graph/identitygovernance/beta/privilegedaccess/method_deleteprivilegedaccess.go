package privilegedaccess

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePrivilegedAccessOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeletePrivilegedAccessOperationOptions struct {
	IfMatch *string
}

func DefaultDeletePrivilegedAccessOperationOptions() DeletePrivilegedAccessOperationOptions {
	return DeletePrivilegedAccessOperationOptions{}
}

func (o DeletePrivilegedAccessOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeletePrivilegedAccessOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeletePrivilegedAccessOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeletePrivilegedAccess - Delete navigation property privilegedAccess for identityGovernance
func (c PrivilegedAccessClient) DeletePrivilegedAccess(ctx context.Context, options DeletePrivilegedAccessOperationOptions) (result DeletePrivilegedAccessOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/identityGovernance/privilegedAccess",
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
