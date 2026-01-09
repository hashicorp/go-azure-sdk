package entitlementmanagementaccesspackageincompatiblegroup

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

type AddEntitlementManagementAccessPackageIncompatibleGroupRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddEntitlementManagementAccessPackageIncompatibleGroupRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddEntitlementManagementAccessPackageIncompatibleGroupRefOperationOptions() AddEntitlementManagementAccessPackageIncompatibleGroupRefOperationOptions {
	return AddEntitlementManagementAccessPackageIncompatibleGroupRefOperationOptions{}
}

func (o AddEntitlementManagementAccessPackageIncompatibleGroupRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddEntitlementManagementAccessPackageIncompatibleGroupRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddEntitlementManagementAccessPackageIncompatibleGroupRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddEntitlementManagementAccessPackageIncompatibleGroupRef - Add group to incompatibleGroups. Add a group to the list
// of groups that have been marked as incompatible on an accessPackage.
func (c EntitlementManagementAccessPackageIncompatibleGroupClient) AddEntitlementManagementAccessPackageIncompatibleGroupRef(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageId, input stable.ReferenceCreate, options AddEntitlementManagementAccessPackageIncompatibleGroupRefOperationOptions) (result AddEntitlementManagementAccessPackageIncompatibleGroupRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/incompatibleGroups/$ref", id.ID()),
		RetryFunc:     options.RetryFunc,
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
