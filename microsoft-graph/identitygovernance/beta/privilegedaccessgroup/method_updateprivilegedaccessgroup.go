package privilegedaccessgroup

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePrivilegedAccessGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePrivilegedAccessGroupOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdatePrivilegedAccessGroupOperationOptions() UpdatePrivilegedAccessGroupOperationOptions {
	return UpdatePrivilegedAccessGroupOperationOptions{}
}

func (o UpdatePrivilegedAccessGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdatePrivilegedAccessGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePrivilegedAccessGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePrivilegedAccessGroup - Update the navigation property group in identityGovernance
func (c PrivilegedAccessGroupClient) UpdatePrivilegedAccessGroup(ctx context.Context, input beta.PrivilegedAccessGroup, options UpdatePrivilegedAccessGroupOperationOptions) (result UpdatePrivilegedAccessGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/identityGovernance/privilegedAccess/group",
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
