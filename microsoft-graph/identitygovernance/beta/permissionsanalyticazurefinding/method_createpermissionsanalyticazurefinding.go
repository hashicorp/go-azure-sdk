package permissionsanalyticazurefinding

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePermissionsAnalyticAzureFindingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.Finding
}

type CreatePermissionsAnalyticAzureFindingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreatePermissionsAnalyticAzureFindingOperationOptions() CreatePermissionsAnalyticAzureFindingOperationOptions {
	return CreatePermissionsAnalyticAzureFindingOperationOptions{}
}

func (o CreatePermissionsAnalyticAzureFindingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePermissionsAnalyticAzureFindingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePermissionsAnalyticAzureFindingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePermissionsAnalyticAzureFinding - Create new navigation property to findings for identityGovernance
func (c PermissionsAnalyticAzureFindingClient) CreatePermissionsAnalyticAzureFinding(ctx context.Context, input beta.Finding, options CreatePermissionsAnalyticAzureFindingOperationOptions) (result CreatePermissionsAnalyticAzureFindingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/identityGovernance/permissionsAnalytics/azure/findings",
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalFindingImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
