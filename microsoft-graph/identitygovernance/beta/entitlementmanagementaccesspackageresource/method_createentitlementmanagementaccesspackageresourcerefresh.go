package entitlementmanagementaccesspackageresource

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

type CreateEntitlementManagementAccessPackageResourceRefreshOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateEntitlementManagementAccessPackageResourceRefreshOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateEntitlementManagementAccessPackageResourceRefreshOperationOptions() CreateEntitlementManagementAccessPackageResourceRefreshOperationOptions {
	return CreateEntitlementManagementAccessPackageResourceRefreshOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageResourceRefreshOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRefreshOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRefreshOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageResourceRefresh - Invoke action refresh. In Microsoft Entra entitlement
// management, refresh the accessPackageResource object to fetch the latest details for displayName, description, and
// resourceType from the origin system. For the AadApplication originSystem, this operation also updates the displayName
// and description for the accessPackageResourceRole.
func (c EntitlementManagementAccessPackageResourceClient) CreateEntitlementManagementAccessPackageResourceRefresh(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageResourceId, options CreateEntitlementManagementAccessPackageResourceRefreshOperationOptions) (result CreateEntitlementManagementAccessPackageResourceRefreshOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/refresh", id.ID()),
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
