package onpremisessynchronization

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOnPremisesSynchronizationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateOnPremisesSynchronizationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateOnPremisesSynchronizationOperationOptions() UpdateOnPremisesSynchronizationOperationOptions {
	return UpdateOnPremisesSynchronizationOperationOptions{}
}

func (o UpdateOnPremisesSynchronizationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateOnPremisesSynchronizationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateOnPremisesSynchronizationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateOnPremisesSynchronization - Update onPremisesDirectorySynchronization. Update the properties of an
// onPremisesDirectorySynchronization object.
func (c OnPremisesSynchronizationClient) UpdateOnPremisesSynchronization(ctx context.Context, id beta.DirectoryOnPremisesSynchronizationId, input beta.OnPremisesDirectorySynchronization, options UpdateOnPremisesSynchronizationOperationOptions) (result UpdateOnPremisesSynchronizationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
