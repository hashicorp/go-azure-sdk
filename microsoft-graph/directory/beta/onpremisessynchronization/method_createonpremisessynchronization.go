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

type CreateOnPremisesSynchronizationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.OnPremisesDirectorySynchronization
}

type CreateOnPremisesSynchronizationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateOnPremisesSynchronizationOperationOptions() CreateOnPremisesSynchronizationOperationOptions {
	return CreateOnPremisesSynchronizationOperationOptions{}
}

func (o CreateOnPremisesSynchronizationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnPremisesSynchronizationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnPremisesSynchronizationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnPremisesSynchronization - Create new navigation property to onPremisesSynchronization for directory
func (c OnPremisesSynchronizationClient) CreateOnPremisesSynchronization(ctx context.Context, input beta.OnPremisesDirectorySynchronization, options CreateOnPremisesSynchronizationOperationOptions) (result CreateOnPremisesSynchronizationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/directory/onPremisesSynchronization",
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

	var model beta.OnPremisesDirectorySynchronization
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
