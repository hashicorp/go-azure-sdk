package synchronizationjobschemadirectory

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSynchronizationJobSchemaDirectoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSynchronizationJobSchemaDirectoryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSynchronizationJobSchemaDirectoryOperationOptions() UpdateSynchronizationJobSchemaDirectoryOperationOptions {
	return UpdateSynchronizationJobSchemaDirectoryOperationOptions{}
}

func (o UpdateSynchronizationJobSchemaDirectoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSynchronizationJobSchemaDirectoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSynchronizationJobSchemaDirectoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSynchronizationJobSchemaDirectory - Update the navigation property directories in applications
func (c SynchronizationJobSchemaDirectoryClient) UpdateSynchronizationJobSchemaDirectory(ctx context.Context, id beta.ApplicationIdSynchronizationJobIdSchemaDirectoryId, input beta.DirectoryDefinition, options UpdateSynchronizationJobSchemaDirectoryOperationOptions) (result UpdateSynchronizationJobSchemaDirectoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
