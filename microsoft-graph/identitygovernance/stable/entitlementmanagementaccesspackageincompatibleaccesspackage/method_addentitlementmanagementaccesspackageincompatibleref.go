package entitlementmanagementaccesspackageincompatibleaccesspackage

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

type AddEntitlementManagementAccessPackageIncompatibleRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// AddEntitlementManagementAccessPackageIncompatibleRef - Add accessPackage to incompatibleAccessPackages. Add an
// accessPackage to the list of access packages that have been marked as incompatible on an accessPackage.
func (c EntitlementManagementAccessPackageIncompatibleAccessPackageClient) AddEntitlementManagementAccessPackageIncompatibleRef(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageId, input stable.ReferenceCreate) (result AddEntitlementManagementAccessPackageIncompatibleRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/incompatibleAccessPackages/$ref", id.ID()),
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
