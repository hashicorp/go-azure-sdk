package identitysecuritydefaultsenforcementpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateIdentitySecurityDefaultsEnforcementPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateIdentitySecurityDefaultsEnforcementPolicy - Update identitySecurityDefaultsEnforcementPolicy. Update the
// properties of an identitySecurityDefaultsEnforcementPolicy object.
func (c IdentitySecurityDefaultsEnforcementPolicyClient) UpdateIdentitySecurityDefaultsEnforcementPolicy(ctx context.Context, input stable.IdentitySecurityDefaultsEnforcementPolicy) (result UpdateIdentitySecurityDefaultsEnforcementPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPatch,
		Path:       "/policies/identitySecurityDefaultsEnforcementPolicy",
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
