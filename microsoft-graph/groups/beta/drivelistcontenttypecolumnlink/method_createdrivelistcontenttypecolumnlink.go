package drivelistcontenttypecolumnlink

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

type CreateDriveListContentTypeColumnLinkOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ColumnLink
}

type CreateDriveListContentTypeColumnLinkOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDriveListContentTypeColumnLinkOperationOptions() CreateDriveListContentTypeColumnLinkOperationOptions {
	return CreateDriveListContentTypeColumnLinkOperationOptions{}
}

func (o CreateDriveListContentTypeColumnLinkOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDriveListContentTypeColumnLinkOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDriveListContentTypeColumnLinkOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDriveListContentTypeColumnLink - Create new navigation property to columnLinks for groups
func (c DriveListContentTypeColumnLinkClient) CreateDriveListContentTypeColumnLink(ctx context.Context, id beta.GroupIdDriveIdListContentTypeId, input beta.ColumnLink, options CreateDriveListContentTypeColumnLinkOperationOptions) (result CreateDriveListContentTypeColumnLinkOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/columnLinks", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.ColumnLink
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
