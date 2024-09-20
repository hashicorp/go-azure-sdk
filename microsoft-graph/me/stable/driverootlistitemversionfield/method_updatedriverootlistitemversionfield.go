package driverootlistitemversionfield

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

type UpdateDriveRootListItemVersionFieldOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDriveRootListItemVersionFieldOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDriveRootListItemVersionFieldOperationOptions() UpdateDriveRootListItemVersionFieldOperationOptions {
	return UpdateDriveRootListItemVersionFieldOperationOptions{}
}

func (o UpdateDriveRootListItemVersionFieldOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDriveRootListItemVersionFieldOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDriveRootListItemVersionFieldOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDriveRootListItemVersionField - Update the navigation property fields in me
func (c DriveRootListItemVersionFieldClient) UpdateDriveRootListItemVersionField(ctx context.Context, id stable.MeDriveIdRootListItemVersionId, input stable.FieldValueSet, options UpdateDriveRootListItemVersionFieldOperationOptions) (result UpdateDriveRootListItemVersionFieldOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/fields", id.ID()),
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
