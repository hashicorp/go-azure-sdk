package synchronizationjobbulkupload

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

type RemoveSynchronizationJobBulkUploadValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveSynchronizationJobBulkUploadValueOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveSynchronizationJobBulkUploadValueOperationOptions() RemoveSynchronizationJobBulkUploadValueOperationOptions {
	return RemoveSynchronizationJobBulkUploadValueOperationOptions{}
}

func (o RemoveSynchronizationJobBulkUploadValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveSynchronizationJobBulkUploadValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveSynchronizationJobBulkUploadValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveSynchronizationJobBulkUploadValue - Delete media content for the navigation property bulkUpload in
// servicePrincipals. The bulk upload operation for the job.
func (c SynchronizationJobBulkUploadClient) RemoveSynchronizationJobBulkUploadValue(ctx context.Context, id beta.ServicePrincipalIdSynchronizationJobId, options RemoveSynchronizationJobBulkUploadValueOperationOptions) (result RemoveSynchronizationJobBulkUploadValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/bulkUpload/$value", id.ID()),
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
