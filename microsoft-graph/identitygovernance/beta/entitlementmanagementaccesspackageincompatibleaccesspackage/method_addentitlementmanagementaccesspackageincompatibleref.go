package entitlementmanagementaccesspackageincompatibleaccesspackage

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

type AddEntitlementManagementAccessPackageIncompatibleRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddEntitlementManagementAccessPackageIncompatibleRefOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultAddEntitlementManagementAccessPackageIncompatibleRefOperationOptions() AddEntitlementManagementAccessPackageIncompatibleRefOperationOptions {
	return AddEntitlementManagementAccessPackageIncompatibleRefOperationOptions{}
}

func (o AddEntitlementManagementAccessPackageIncompatibleRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddEntitlementManagementAccessPackageIncompatibleRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddEntitlementManagementAccessPackageIncompatibleRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddEntitlementManagementAccessPackageIncompatibleRef - Add accessPackage to incompatibleAccessPackages. Add an
// accessPackage to the list of access packages marked as incompatible on an accessPackage.
func (c EntitlementManagementAccessPackageIncompatibleAccessPackageClient) AddEntitlementManagementAccessPackageIncompatibleRef(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, input beta.ReferenceCreate, options AddEntitlementManagementAccessPackageIncompatibleRefOperationOptions) (result AddEntitlementManagementAccessPackageIncompatibleRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/incompatibleAccessPackages/$ref", id.ID()),
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
