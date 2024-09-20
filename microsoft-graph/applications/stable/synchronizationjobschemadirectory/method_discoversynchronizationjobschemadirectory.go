package synchronizationjobschemadirectory

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

type DiscoverSynchronizationJobSchemaDirectoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DirectoryDefinition
}

type DiscoverSynchronizationJobSchemaDirectoryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultDiscoverSynchronizationJobSchemaDirectoryOperationOptions() DiscoverSynchronizationJobSchemaDirectoryOperationOptions {
	return DiscoverSynchronizationJobSchemaDirectoryOperationOptions{}
}

func (o DiscoverSynchronizationJobSchemaDirectoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DiscoverSynchronizationJobSchemaDirectoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DiscoverSynchronizationJobSchemaDirectoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DiscoverSynchronizationJobSchemaDirectory - Invoke action discover. Discover the latest schema definition for
// provisioning to an application.
func (c SynchronizationJobSchemaDirectoryClient) DiscoverSynchronizationJobSchemaDirectory(ctx context.Context, id stable.ApplicationIdSynchronizationJobIdSchemaDirectoryId, options DiscoverSynchronizationJobSchemaDirectoryOperationOptions) (result DiscoverSynchronizationJobSchemaDirectoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/discover", id.ID()),
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

	var model stable.DirectoryDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
