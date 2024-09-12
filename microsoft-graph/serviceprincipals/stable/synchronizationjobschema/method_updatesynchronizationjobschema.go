package synchronizationjobschema

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

type UpdateSynchronizationJobSchemaOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateSynchronizationJobSchema - Update synchronizationSchema. Update the synchronization schema for a given job or
// template. This method fully replaces the current schema with the one provided in the request. To update the schema of
// a template, make the call on the application object. You must be the owner of the application.
func (c SynchronizationJobSchemaClient) UpdateSynchronizationJobSchema(ctx context.Context, id stable.ServicePrincipalIdSynchronizationJobId, input stable.SynchronizationSchema) (result UpdateSynchronizationJobSchemaOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPatch,
		Path:       fmt.Sprintf("%s/schema", id.ID()),
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
