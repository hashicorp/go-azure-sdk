package crosstenantaccesspolicypartneridentitysynchronization

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

type DeleteCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions() DeleteCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions {
	return DeleteCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions{}
}

func (o DeleteCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteCrossTenantAccessPolicyPartnerIdentitySynchronization - Delete crossTenantIdentitySyncPolicyPartner. Delete the
// user synchronization policy for a partner-specific configuration.
func (c CrossTenantAccessPolicyPartnerIdentitySynchronizationClient) DeleteCrossTenantAccessPolicyPartnerIdentitySynchronization(ctx context.Context, id stable.PolicyCrossTenantAccessPolicyPartnerId, options DeleteCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationOptions) (result DeleteCrossTenantAccessPolicyPartnerIdentitySynchronizationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/identitySynchronization", id.ID()),
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

	return
}
