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

type DiscoverSynchronizationTemplateSchemaDirectoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DirectoryDefinition
}

type DiscoverSynchronizationTemplateSchemaDirectoryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultDiscoverSynchronizationTemplateSchemaDirectoryOperationOptions() DiscoverSynchronizationTemplateSchemaDirectoryOperationOptions {
	return DiscoverSynchronizationTemplateSchemaDirectoryOperationOptions{}
}

func (o DiscoverSynchronizationTemplateSchemaDirectoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DiscoverSynchronizationTemplateSchemaDirectoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DiscoverSynchronizationTemplateSchemaDirectoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DiscoverSynchronizationTemplateSchemaDirectory - Invoke action discover. Discover the latest schema definition for
// provisioning to an application.
func (c SynchronizationTemplateSchemaDirectoryClient) DiscoverSynchronizationTemplateSchemaDirectory(ctx context.Context, id stable.ServicePrincipalIdSynchronizationTemplateIdSchemaDirectoryId, options DiscoverSynchronizationTemplateSchemaDirectoryOperationOptions) (result DiscoverSynchronizationTemplateSchemaDirectoryOperationResponse, err error) {
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
