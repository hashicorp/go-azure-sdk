package carttoclassassociation

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

type DeleteCartToClassAssociationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteCartToClassAssociationOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteCartToClassAssociationOperationOptions() DeleteCartToClassAssociationOperationOptions {
	return DeleteCartToClassAssociationOperationOptions{}
}

func (o DeleteCartToClassAssociationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteCartToClassAssociationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteCartToClassAssociationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteCartToClassAssociation - Delete navigation property cartToClassAssociations for deviceManagement
func (c CartToClassAssociationClient) DeleteCartToClassAssociation(ctx context.Context, id beta.DeviceManagementCartToClassAssociationId, options DeleteCartToClassAssociationOperationOptions) (result DeleteCartToClassAssociationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
