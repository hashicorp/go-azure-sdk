package domain

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

type CreateForceDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateForceDelete - Invoke action forceDelete. Delete a domain using an asynchronous long-running operation. Before
// performing this operation, you must update or remove any references to Exchange as the provisioning service. The
// following actions are performed as part of this operation: After the domain deletion completes, API operations for
// the deleted domain return an HTTP 404 status code. To verify deletion of a domain, you can perform a get domain
// operation.
func (c DomainClient) CreateForceDelete(ctx context.Context, id stable.DomainId, input CreateForceDeleteRequest) (result CreateForceDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/forceDelete", id.ID()),
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
