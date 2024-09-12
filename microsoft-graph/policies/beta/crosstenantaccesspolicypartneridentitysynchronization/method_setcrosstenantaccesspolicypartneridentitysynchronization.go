package crosstenantaccesspolicypartneridentitysynchronization

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

type SetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// SetCrossTenantAccessPolicyPartnerIdentitySynchronization - Update crossTenantIdentitySyncPolicyPartner. Update the
// user synchronization policy of a partner-specific configuration.
func (c CrossTenantAccessPolicyPartnerIdentitySynchronizationClient) SetCrossTenantAccessPolicyPartnerIdentitySynchronization(ctx context.Context, id beta.PolicyCrossTenantAccessPolicyPartnerId, input beta.CrossTenantIdentitySyncPolicyPartner) (result SetCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/identitySynchronization", id.ID()),
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
