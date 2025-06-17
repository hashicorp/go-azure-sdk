package datasecurityandgovernance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDataSecurityAndGovernanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDataSecurityAndGovernanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDataSecurityAndGovernanceOperationOptions() UpdateDataSecurityAndGovernanceOperationOptions {
	return UpdateDataSecurityAndGovernanceOperationOptions{}
}

func (o UpdateDataSecurityAndGovernanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDataSecurityAndGovernanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDataSecurityAndGovernanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDataSecurityAndGovernance - Update the navigation property dataSecurityAndGovernance in users
func (c DataSecurityAndGovernanceClient) UpdateDataSecurityAndGovernance(ctx context.Context, id beta.UserId, input beta.UserDataSecurityAndGovernance, options UpdateDataSecurityAndGovernanceOperationOptions) (result UpdateDataSecurityAndGovernanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/dataSecurityAndGovernance", id.ID()),
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
