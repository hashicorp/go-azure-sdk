package driveitemlistitemdocumentsetversionfield

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

type GetDriveItemListItemDocumentSetVersionFieldOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.FieldValueSet
}

type GetDriveItemListItemDocumentSetVersionFieldOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDriveItemListItemDocumentSetVersionFieldOperationOptions() GetDriveItemListItemDocumentSetVersionFieldOperationOptions {
	return GetDriveItemListItemDocumentSetVersionFieldOperationOptions{}
}

func (o GetDriveItemListItemDocumentSetVersionFieldOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveItemListItemDocumentSetVersionFieldOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDriveItemListItemDocumentSetVersionFieldOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveItemListItemDocumentSetVersionField - Get fields from users. A collection of the fields and values for this
// version of the list item.
func (c DriveItemListItemDocumentSetVersionFieldClient) GetDriveItemListItemDocumentSetVersionField(ctx context.Context, id stable.UserIdDriveIdItemIdListItemDocumentSetVersionId, options GetDriveItemListItemDocumentSetVersionFieldOperationOptions) (result GetDriveItemListItemDocumentSetVersionFieldOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/fields", id.ID()),
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

	var model stable.FieldValueSet
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
