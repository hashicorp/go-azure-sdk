package carttoclassassociation

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCartToClassAssociationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCartToClassAssociationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateCartToClassAssociationOperationOptions() UpdateCartToClassAssociationOperationOptions {
	return UpdateCartToClassAssociationOperationOptions{}
}

func (o UpdateCartToClassAssociationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCartToClassAssociationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCartToClassAssociationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCartToClassAssociation - Update the navigation property cartToClassAssociations in deviceManagement
func (c CartToClassAssociationClient) UpdateCartToClassAssociation(ctx context.Context, id beta.DeviceManagementCartToClassAssociationId, input beta.CartToClassAssociation, options UpdateCartToClassAssociationOperationOptions) (result UpdateCartToClassAssociationOperationResponse, err error) {
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
