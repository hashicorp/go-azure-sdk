package driverootlistitemdocumentsetversionfield

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

type GetDriveRootListItemDocumentSetVersionFieldOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.FieldValueSet
}

type GetDriveRootListItemDocumentSetVersionFieldOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetDriveRootListItemDocumentSetVersionFieldOperationOptions() GetDriveRootListItemDocumentSetVersionFieldOperationOptions {
	return GetDriveRootListItemDocumentSetVersionFieldOperationOptions{}
}

func (o GetDriveRootListItemDocumentSetVersionFieldOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveRootListItemDocumentSetVersionFieldOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDriveRootListItemDocumentSetVersionFieldOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveRootListItemDocumentSetVersionField - Get fields from users. A collection of the fields and values for this
// version of the list item.
func (c DriveRootListItemDocumentSetVersionFieldClient) GetDriveRootListItemDocumentSetVersionField(ctx context.Context, id beta.UserIdDriveIdRootListItemDocumentSetVersionId, options GetDriveRootListItemDocumentSetVersionFieldOperationOptions) (result GetDriveRootListItemDocumentSetVersionFieldOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/fields", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.FieldValueSet
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
