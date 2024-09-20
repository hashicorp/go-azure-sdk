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

type CreateCartToClassAssociationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CartToClassAssociation
}

type CreateCartToClassAssociationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateCartToClassAssociationOperationOptions() CreateCartToClassAssociationOperationOptions {
	return CreateCartToClassAssociationOperationOptions{}
}

func (o CreateCartToClassAssociationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCartToClassAssociationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCartToClassAssociationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCartToClassAssociation - Create new navigation property to cartToClassAssociations for deviceManagement
func (c CartToClassAssociationClient) CreateCartToClassAssociation(ctx context.Context, input beta.CartToClassAssociation, options CreateCartToClassAssociationOperationOptions) (result CreateCartToClassAssociationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/cartToClassAssociations",
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

	var model beta.CartToClassAssociation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
