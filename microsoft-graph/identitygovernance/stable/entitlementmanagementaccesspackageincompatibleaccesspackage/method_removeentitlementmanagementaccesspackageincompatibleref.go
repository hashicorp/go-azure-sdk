package entitlementmanagementaccesspackageincompatibleaccesspackage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveEntitlementManagementAccessPackageIncompatibleRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveEntitlementManagementAccessPackageIncompatibleRefOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveEntitlementManagementAccessPackageIncompatibleRefOperationOptions() RemoveEntitlementManagementAccessPackageIncompatibleRefOperationOptions {
	return RemoveEntitlementManagementAccessPackageIncompatibleRefOperationOptions{}
}

func (o RemoveEntitlementManagementAccessPackageIncompatibleRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveEntitlementManagementAccessPackageIncompatibleRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveEntitlementManagementAccessPackageIncompatibleRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveEntitlementManagementAccessPackageIncompatibleRef - Remove accessPackage from incompatibleAccessPackages.
// Remove an access package from the list of access packages that have been marked as incompatible on an accessPackage.
func (c EntitlementManagementAccessPackageIncompatibleAccessPackageClient) RemoveEntitlementManagementAccessPackageIncompatibleRef(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageIdIncompatibleAccessPackageId, options RemoveEntitlementManagementAccessPackageIncompatibleRefOperationOptions) (result RemoveEntitlementManagementAccessPackageIncompatibleRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$ref", id.ID()),
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
