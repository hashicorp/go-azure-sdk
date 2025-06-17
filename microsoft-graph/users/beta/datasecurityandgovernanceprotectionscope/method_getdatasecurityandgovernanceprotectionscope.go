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

type GetDataSecurityAndGovernanceProtectionScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserProtectionScopeContainer
}

type GetDataSecurityAndGovernanceProtectionScopeOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetDataSecurityAndGovernanceProtectionScopeOperationOptions() GetDataSecurityAndGovernanceProtectionScopeOperationOptions {
	return GetDataSecurityAndGovernanceProtectionScopeOperationOptions{}
}

func (o GetDataSecurityAndGovernanceProtectionScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDataSecurityAndGovernanceProtectionScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDataSecurityAndGovernanceProtectionScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDataSecurityAndGovernanceProtectionScope - Get protectionScopes from users
func (c DataSecurityAndGovernanceProtectionScopeClient) GetDataSecurityAndGovernanceProtectionScope(ctx context.Context, id beta.UserId, options GetDataSecurityAndGovernanceProtectionScopeOperationOptions) (result GetDataSecurityAndGovernanceProtectionScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/dataSecurityAndGovernance/protectionScopes", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.UserProtectionScopeContainer
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
