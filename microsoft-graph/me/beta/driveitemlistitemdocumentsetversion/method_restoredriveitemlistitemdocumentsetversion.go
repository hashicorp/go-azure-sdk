package driveitemlistitemdocumentsetversion

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreDriveItemListItemDocumentSetVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RestoreDriveItemListItemDocumentSetVersionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRestoreDriveItemListItemDocumentSetVersionOperationOptions() RestoreDriveItemListItemDocumentSetVersionOperationOptions {
	return RestoreDriveItemListItemDocumentSetVersionOperationOptions{}
}

func (o RestoreDriveItemListItemDocumentSetVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RestoreDriveItemListItemDocumentSetVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RestoreDriveItemListItemDocumentSetVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RestoreDriveItemListItemDocumentSetVersion - Invoke action restore. Restore a document set version.
func (c DriveItemListItemDocumentSetVersionClient) RestoreDriveItemListItemDocumentSetVersion(ctx context.Context, id beta.MeDriveIdItemIdListItemDocumentSetVersionId, options RestoreDriveItemListItemDocumentSetVersionOperationOptions) (result RestoreDriveItemListItemDocumentSetVersionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/restore", id.ID()),
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

	return
}
