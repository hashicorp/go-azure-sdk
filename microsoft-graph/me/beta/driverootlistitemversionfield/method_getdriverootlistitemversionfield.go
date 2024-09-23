package driverootlistitemversionfield

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

type GetDriveRootListItemVersionFieldOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.FieldValueSet
}

type GetDriveRootListItemVersionFieldOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetDriveRootListItemVersionFieldOperationOptions() GetDriveRootListItemVersionFieldOperationOptions {
	return GetDriveRootListItemVersionFieldOperationOptions{}
}

func (o GetDriveRootListItemVersionFieldOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveRootListItemVersionFieldOperationOptions) ToOData() *odata.Query {
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

func (o GetDriveRootListItemVersionFieldOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveRootListItemVersionField - Get fields from me. A collection of the fields and values for this version of the
// list item.
func (c DriveRootListItemVersionFieldClient) GetDriveRootListItemVersionField(ctx context.Context, id beta.MeDriveIdRootListItemVersionId, options GetDriveRootListItemVersionFieldOperationOptions) (result GetDriveRootListItemVersionFieldOperationResponse, err error) {
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
