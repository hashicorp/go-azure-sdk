package synchronization

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

type AcquireSynchronizationAccessTokenOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AcquireSynchronizationAccessTokenOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAcquireSynchronizationAccessTokenOperationOptions() AcquireSynchronizationAccessTokenOperationOptions {
	return AcquireSynchronizationAccessTokenOperationOptions{}
}

func (o AcquireSynchronizationAccessTokenOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AcquireSynchronizationAccessTokenOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AcquireSynchronizationAccessTokenOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AcquireSynchronizationAccessToken - Invoke action acquireAccessToken. Acquire an OAuth access token to authorize the
// Microsoft Entra provisioning service to provision users into an application.
func (c SynchronizationClient) AcquireSynchronizationAccessToken(ctx context.Context, id stable.ServicePrincipalId, input AcquireSynchronizationAccessTokenRequest, options AcquireSynchronizationAccessTokenOperationOptions) (result AcquireSynchronizationAccessTokenOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/synchronization/acquireAccessToken", id.ID()),
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

	return
}
