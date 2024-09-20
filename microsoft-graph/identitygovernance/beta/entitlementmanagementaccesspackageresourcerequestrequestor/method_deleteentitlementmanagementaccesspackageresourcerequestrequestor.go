package entitlementmanagementaccesspackageresourcerequestrequestor

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

type DeleteEntitlementManagementAccessPackageResourceRequestRequestorOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementAccessPackageResourceRequestRequestorOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteEntitlementManagementAccessPackageResourceRequestRequestorOperationOptions() DeleteEntitlementManagementAccessPackageResourceRequestRequestorOperationOptions {
	return DeleteEntitlementManagementAccessPackageResourceRequestRequestorOperationOptions{}
}

func (o DeleteEntitlementManagementAccessPackageResourceRequestRequestorOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageResourceRequestRequestorOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageResourceRequestRequestorOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementAccessPackageResourceRequestRequestor - Delete navigation property requestor for
// identityGovernance
func (c EntitlementManagementAccessPackageResourceRequestRequestorClient) DeleteEntitlementManagementAccessPackageResourceRequestRequestor(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageResourceRequestId, options DeleteEntitlementManagementAccessPackageResourceRequestRequestorOperationOptions) (result DeleteEntitlementManagementAccessPackageResourceRequestRequestorOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/requestor", id.ID()),
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
