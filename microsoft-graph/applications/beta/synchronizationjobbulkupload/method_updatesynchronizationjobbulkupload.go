package synchronizationjobbulkupload

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

type UpdateSynchronizationJobBulkUploadOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSynchronizationJobBulkUploadOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSynchronizationJobBulkUploadOperationOptions() UpdateSynchronizationJobBulkUploadOperationOptions {
	return UpdateSynchronizationJobBulkUploadOperationOptions{}
}

func (o UpdateSynchronizationJobBulkUploadOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSynchronizationJobBulkUploadOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSynchronizationJobBulkUploadOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSynchronizationJobBulkUpload - Update the navigation property bulkUpload in applications
func (c SynchronizationJobBulkUploadClient) UpdateSynchronizationJobBulkUpload(ctx context.Context, id beta.ApplicationIdSynchronizationJobId, input beta.BulkUpload, options UpdateSynchronizationJobBulkUploadOperationOptions) (result UpdateSynchronizationJobBulkUploadOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/bulkUpload", id.ID()),
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
