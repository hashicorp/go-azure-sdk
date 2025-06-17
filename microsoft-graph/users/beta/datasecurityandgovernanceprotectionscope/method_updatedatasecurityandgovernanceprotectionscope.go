package datasecurityandgovernanceprotectionscope

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

type UpdateDataSecurityAndGovernanceProtectionScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDataSecurityAndGovernanceProtectionScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDataSecurityAndGovernanceProtectionScopeOperationOptions() UpdateDataSecurityAndGovernanceProtectionScopeOperationOptions {
	return UpdateDataSecurityAndGovernanceProtectionScopeOperationOptions{}
}

func (o UpdateDataSecurityAndGovernanceProtectionScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDataSecurityAndGovernanceProtectionScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDataSecurityAndGovernanceProtectionScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDataSecurityAndGovernanceProtectionScope - Update the navigation property protectionScopes in users
func (c DataSecurityAndGovernanceProtectionScopeClient) UpdateDataSecurityAndGovernanceProtectionScope(ctx context.Context, id beta.UserId, input beta.UserProtectionScopeContainer, options UpdateDataSecurityAndGovernanceProtectionScopeOperationOptions) (result UpdateDataSecurityAndGovernanceProtectionScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/dataSecurityAndGovernance/protectionScopes", id.ID()),
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
