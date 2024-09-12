package crosstenantaccesspolicypartner

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCrossTenantAccessPolicyPartnerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.CrossTenantAccessPolicyConfigurationPartner
}

// CreateCrossTenantAccessPolicyPartner - Create crossTenantAccessPolicyConfigurationPartner. Create a new partner
// configuration in a cross-tenant access policy.
func (c CrossTenantAccessPolicyPartnerClient) CreateCrossTenantAccessPolicyPartner(ctx context.Context, input stable.CrossTenantAccessPolicyConfigurationPartner) (result CreateCrossTenantAccessPolicyPartnerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       "/policies/crossTenantAccessPolicy/partners",
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

	var model stable.CrossTenantAccessPolicyConfigurationPartner
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
