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

type CreateSynchronizationJobSchemaDirectoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DirectoryDefinition
}

type CreateSynchronizationJobSchemaDirectoryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateSynchronizationJobSchemaDirectoryOperationOptions() CreateSynchronizationJobSchemaDirectoryOperationOptions {
	return CreateSynchronizationJobSchemaDirectoryOperationOptions{}
}

func (o CreateSynchronizationJobSchemaDirectoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSynchronizationJobSchemaDirectoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSynchronizationJobSchemaDirectoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSynchronizationJobSchemaDirectory - Create new navigation property to directories for servicePrincipals
func (c SynchronizationJobSchemaDirectoryClient) CreateSynchronizationJobSchemaDirectory(ctx context.Context, id stable.ServicePrincipalIdSynchronizationJobId, input stable.DirectoryDefinition, options CreateSynchronizationJobSchemaDirectoryOperationOptions) (result CreateSynchronizationJobSchemaDirectoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/schema/directories", id.ID()),
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

	var model stable.DirectoryDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
