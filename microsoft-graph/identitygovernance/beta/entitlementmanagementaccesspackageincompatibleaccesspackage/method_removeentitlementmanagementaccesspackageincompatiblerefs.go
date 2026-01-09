package entitlementmanagementaccesspackageincompatibleaccesspackage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveEntitlementManagementAccessPackageIncompatibleRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveEntitlementManagementAccessPackageIncompatibleRefsOperationOptions struct {
	Id        *string
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveEntitlementManagementAccessPackageIncompatibleRefsOperationOptions() RemoveEntitlementManagementAccessPackageIncompatibleRefsOperationOptions {
	return RemoveEntitlementManagementAccessPackageIncompatibleRefsOperationOptions{}
}

func (o RemoveEntitlementManagementAccessPackageIncompatibleRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveEntitlementManagementAccessPackageIncompatibleRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveEntitlementManagementAccessPackageIncompatibleRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Id != nil {
		out.Append("@id", fmt.Sprintf("%v", *o.Id))
	}
	return &out
}

// RemoveEntitlementManagementAccessPackageIncompatibleRefs - Remove accessPackage from incompatibleAccessPackages.
// Remove an access package from the list of access packages marked as incompatible on an accessPackage.
func (c EntitlementManagementAccessPackageIncompatibleAccessPackageClient) RemoveEntitlementManagementAccessPackageIncompatibleRefs(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options RemoveEntitlementManagementAccessPackageIncompatibleRefsOperationOptions) (result RemoveEntitlementManagementAccessPackageIncompatibleRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/incompatibleAccessPackages/$ref", id.ID()),
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
