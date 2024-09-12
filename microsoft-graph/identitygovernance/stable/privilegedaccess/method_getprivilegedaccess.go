package privilegedaccess

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivilegedAccessOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PrivilegedAccessRoot
}

type GetPrivilegedAccessOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetPrivilegedAccessOperationOptions() GetPrivilegedAccessOperationOptions {
	return GetPrivilegedAccessOperationOptions{}
}

func (o GetPrivilegedAccessOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPrivilegedAccessOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccess - Get privilegedAccess from identityGovernance
func (c PrivilegedAccessClient) GetPrivilegedAccess(ctx context.Context, options GetPrivilegedAccessOperationOptions) (result GetPrivilegedAccessOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model stable.PrivilegedAccessRoot
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
