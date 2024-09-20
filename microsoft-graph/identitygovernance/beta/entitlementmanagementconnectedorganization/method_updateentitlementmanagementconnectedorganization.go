package entitlementmanagementconnectedorganization

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementConnectedOrganizationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementConnectedOrganizationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementConnectedOrganizationOperationOptions() UpdateEntitlementManagementConnectedOrganizationOperationOptions {
	return UpdateEntitlementManagementConnectedOrganizationOperationOptions{}
}

func (o UpdateEntitlementManagementConnectedOrganizationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementConnectedOrganizationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementConnectedOrganizationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementConnectedOrganization - Update connectedOrganization. Update a connectedOrganization
// object to change one or more of its properties.
func (c EntitlementManagementConnectedOrganizationClient) UpdateEntitlementManagementConnectedOrganization(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementConnectedOrganizationId, input beta.ConnectedOrganization, options UpdateEntitlementManagementConnectedOrganizationOperationOptions) (result UpdateEntitlementManagementConnectedOrganizationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
