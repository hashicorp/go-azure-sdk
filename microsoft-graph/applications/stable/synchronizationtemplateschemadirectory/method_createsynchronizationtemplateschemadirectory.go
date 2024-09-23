package synchronizationtemplateschemadirectory

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

type CreateSynchronizationTemplateSchemaDirectoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DirectoryDefinition
}

type CreateSynchronizationTemplateSchemaDirectoryOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSynchronizationTemplateSchemaDirectoryOperationOptions() CreateSynchronizationTemplateSchemaDirectoryOperationOptions {
	return CreateSynchronizationTemplateSchemaDirectoryOperationOptions{}
}

func (o CreateSynchronizationTemplateSchemaDirectoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSynchronizationTemplateSchemaDirectoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSynchronizationTemplateSchemaDirectoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSynchronizationTemplateSchemaDirectory - Create new navigation property to directories for applications
func (c SynchronizationTemplateSchemaDirectoryClient) CreateSynchronizationTemplateSchemaDirectory(ctx context.Context, id stable.ApplicationIdSynchronizationTemplateId, input stable.DirectoryDefinition, options CreateSynchronizationTemplateSchemaDirectoryOperationOptions) (result CreateSynchronizationTemplateSchemaDirectoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/schema/directories", id.ID()),
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

	var model stable.DirectoryDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
