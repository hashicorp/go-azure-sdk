package externalidentitiespolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateExternalIdentitiesPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateExternalIdentitiesPolicyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateExternalIdentitiesPolicyOperationOptions() UpdateExternalIdentitiesPolicyOperationOptions {
	return UpdateExternalIdentitiesPolicyOperationOptions{}
}

func (o UpdateExternalIdentitiesPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateExternalIdentitiesPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateExternalIdentitiesPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateExternalIdentitiesPolicy - Update externalIdentitiesPolicy. Update the settings of the tenant-wide
// externalIdentitiesPolicy object that controls whether external users can leave a Microsoft Entra tenant via
// self-service controls.
func (c ExternalIdentitiesPolicyClient) UpdateExternalIdentitiesPolicy(ctx context.Context, input beta.ExternalIdentitiesPolicy, options UpdateExternalIdentitiesPolicyOperationOptions) (result UpdateExternalIdentitiesPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/policies/externalIdentitiesPolicy",
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
