package entitlementmanagementaccesspackageassignmentaccesspackage

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

type MoveEntitlementManagementAccessPackageAssignmentToCatalogOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type MoveEntitlementManagementAccessPackageAssignmentToCatalogOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultMoveEntitlementManagementAccessPackageAssignmentToCatalogOperationOptions() MoveEntitlementManagementAccessPackageAssignmentToCatalogOperationOptions {
	return MoveEntitlementManagementAccessPackageAssignmentToCatalogOperationOptions{}
}

func (o MoveEntitlementManagementAccessPackageAssignmentToCatalogOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o MoveEntitlementManagementAccessPackageAssignmentToCatalogOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o MoveEntitlementManagementAccessPackageAssignmentToCatalogOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// MoveEntitlementManagementAccessPackageAssignmentToCatalog - Invoke action moveToCatalog. In Microsoft Entra
// entitlement management, this action moves the accessPackage to a specified target accessPackageCatalog. The resources
// in the access package must be present in the target catalog.
func (c EntitlementManagementAccessPackageAssignmentAccessPackageClient) MoveEntitlementManagementAccessPackageAssignmentToCatalog(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, input MoveEntitlementManagementAccessPackageAssignmentToCatalogRequest, options MoveEntitlementManagementAccessPackageAssignmentToCatalogOperationOptions) (result MoveEntitlementManagementAccessPackageAssignmentToCatalogOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackage/moveToCatalog", id.ID()),
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
