package entitlementmanagementaccesspackageincompatiblegroup

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

type RemoveEntitlementManagementAccessPackageIncompatibleGroupRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions struct {
	Id        *string
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions() RemoveEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions {
	return RemoveEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions{}
}

func (o RemoveEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Id != nil {
		out.Append("@id", fmt.Sprintf("%v", *o.Id))
	}
	return &out
}

// RemoveEntitlementManagementAccessPackageIncompatibleGroupRefs - Remove group from incompatibleGroups. Remove a group
// from the list of groups marked as incompatible on an accessPackage.
func (c EntitlementManagementAccessPackageIncompatibleGroupClient) RemoveEntitlementManagementAccessPackageIncompatibleGroupRefs(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options RemoveEntitlementManagementAccessPackageIncompatibleGroupRefsOperationOptions) (result RemoveEntitlementManagementAccessPackageIncompatibleGroupRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/incompatibleGroups/$ref", id.ID()),
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
