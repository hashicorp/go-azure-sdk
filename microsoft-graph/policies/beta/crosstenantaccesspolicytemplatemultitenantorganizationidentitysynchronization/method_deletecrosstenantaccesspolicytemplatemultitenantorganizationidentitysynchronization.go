package crosstenantaccesspolicytemplatemultitenantorganizationidentitysynchronization

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions() DeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions {
	return DeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions{}
}

func (o DeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronization - Delete navigation property
// multiTenantOrganizationIdentitySynchronization for policies
func (c CrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClient) DeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronization(ctx context.Context, options DeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationOptions) (result DeleteCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/policies/crossTenantAccessPolicy/templates/multiTenantOrganizationIdentitySynchronization",
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
