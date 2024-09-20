package drivelistcontenttypecolumn

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

type CreateDriveListContentTypeColumnOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ColumnDefinition
}

type CreateDriveListContentTypeColumnOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDriveListContentTypeColumnOperationOptions() CreateDriveListContentTypeColumnOperationOptions {
	return CreateDriveListContentTypeColumnOperationOptions{}
}

func (o CreateDriveListContentTypeColumnOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDriveListContentTypeColumnOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDriveListContentTypeColumnOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDriveListContentTypeColumn - Create new navigation property to columns for groups
func (c DriveListContentTypeColumnClient) CreateDriveListContentTypeColumn(ctx context.Context, id beta.GroupIdDriveIdListContentTypeId, input beta.ColumnDefinition, options CreateDriveListContentTypeColumnOperationOptions) (result CreateDriveListContentTypeColumnOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/columns", id.ID()),
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

	var model beta.ColumnDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
