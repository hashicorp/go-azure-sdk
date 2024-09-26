package outboundshareduserprofiletenant

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

type CreateOutboundSharedUserProfileTenantOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TenantReference
}

type CreateOutboundSharedUserProfileTenantOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateOutboundSharedUserProfileTenantOperationOptions() CreateOutboundSharedUserProfileTenantOperationOptions {
	return CreateOutboundSharedUserProfileTenantOperationOptions{}
}

func (o CreateOutboundSharedUserProfileTenantOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOutboundSharedUserProfileTenantOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOutboundSharedUserProfileTenantOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOutboundSharedUserProfileTenant - Create new navigation property to tenants for directory
func (c OutboundSharedUserProfileTenantClient) CreateOutboundSharedUserProfileTenant(ctx context.Context, id beta.DirectoryOutboundSharedUserProfileId, input beta.TenantReference, options CreateOutboundSharedUserProfileTenantOperationOptions) (result CreateOutboundSharedUserProfileTenantOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/tenants", id.ID()),
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

	var model beta.TenantReference
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
