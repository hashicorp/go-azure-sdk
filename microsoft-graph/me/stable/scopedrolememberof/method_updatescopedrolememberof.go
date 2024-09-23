package scopedrolememberof

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateScopedRoleMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateScopedRoleMemberOfOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateScopedRoleMemberOfOperationOptions() UpdateScopedRoleMemberOfOperationOptions {
	return UpdateScopedRoleMemberOfOperationOptions{}
}

func (o UpdateScopedRoleMemberOfOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateScopedRoleMemberOfOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateScopedRoleMemberOfOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateScopedRoleMemberOf - Update the navigation property scopedRoleMemberOf in me
func (c ScopedRoleMemberOfClient) UpdateScopedRoleMemberOf(ctx context.Context, id stable.MeScopedRoleMemberOfId, input stable.ScopedRoleMembership, options UpdateScopedRoleMemberOfOperationOptions) (result UpdateScopedRoleMemberOfOperationResponse, err error) {
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
