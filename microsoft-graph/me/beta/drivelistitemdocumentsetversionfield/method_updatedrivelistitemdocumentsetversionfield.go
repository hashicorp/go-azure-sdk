package drivelistitemdocumentsetversionfield

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

type UpdateDriveListItemDocumentSetVersionFieldOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDriveListItemDocumentSetVersionFieldOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDriveListItemDocumentSetVersionFieldOperationOptions() UpdateDriveListItemDocumentSetVersionFieldOperationOptions {
	return UpdateDriveListItemDocumentSetVersionFieldOperationOptions{}
}

func (o UpdateDriveListItemDocumentSetVersionFieldOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDriveListItemDocumentSetVersionFieldOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDriveListItemDocumentSetVersionFieldOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDriveListItemDocumentSetVersionField - Update the navigation property fields in me
func (c DriveListItemDocumentSetVersionFieldClient) UpdateDriveListItemDocumentSetVersionField(ctx context.Context, id beta.MeDriveIdListItemIdDocumentSetVersionId, input beta.FieldValueSet, options UpdateDriveListItemDocumentSetVersionFieldOperationOptions) (result UpdateDriveListItemDocumentSetVersionFieldOperationResponse, err error) {
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
