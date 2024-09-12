package tokenissuancepolicy

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

type RemoveTokenIssuancePolicyRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveTokenIssuancePolicyRefOperationOptions struct {
	IfMatch *string
}

func DefaultRemoveTokenIssuancePolicyRefOperationOptions() RemoveTokenIssuancePolicyRefOperationOptions {
	return RemoveTokenIssuancePolicyRefOperationOptions{}
}

func (o RemoveTokenIssuancePolicyRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveTokenIssuancePolicyRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RemoveTokenIssuancePolicyRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveTokenIssuancePolicyRef - Remove tokenIssuancePolicy. Remove a tokenIssuancePolicy from an application.
func (c TokenIssuancePolicyClient) RemoveTokenIssuancePolicyRef(ctx context.Context, id stable.ApplicationIdTokenIssuancePolicyId, options RemoveTokenIssuancePolicyRefOperationOptions) (result RemoveTokenIssuancePolicyRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$ref", id.ID()),
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

	return
}
