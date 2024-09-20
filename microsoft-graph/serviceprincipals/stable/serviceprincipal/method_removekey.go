package serviceprincipal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveKeyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveKeyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultRemoveKeyOperationOptions() RemoveKeyOperationOptions {
	return RemoveKeyOperationOptions{}
}

func (o RemoveKeyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveKeyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveKeyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveKey - Invoke action removeKey. Remove a key credential from a servicePrincipal. This method along with addKey
// can be used by a servicePrincipal to automate rolling its expiring keys. As part of the request validation for this
// method, a proof of possession of an existing key is verified before the action can be performed.
func (c ServicePrincipalClient) RemoveKey(ctx context.Context, id stable.ServicePrincipalId, input RemoveKeyRequest, options RemoveKeyOperationOptions) (result RemoveKeyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/removeKey", id.ID()),
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
