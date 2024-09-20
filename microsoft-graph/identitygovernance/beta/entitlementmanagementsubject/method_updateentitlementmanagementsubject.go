package entitlementmanagementsubject

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementSubjectOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementSubjectOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementSubjectOperationOptions() UpdateEntitlementManagementSubjectOperationOptions {
	return UpdateEntitlementManagementSubjectOperationOptions{}
}

func (o UpdateEntitlementManagementSubjectOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementSubjectOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementSubjectOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementSubject - Update accessPackageSubject. Update an existing accessPackageSubject object to
// change the subject lifecycle.
func (c EntitlementManagementSubjectClient) UpdateEntitlementManagementSubject(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementSubjectId, input beta.AccessPackageSubject, options UpdateEntitlementManagementSubjectOperationOptions) (result UpdateEntitlementManagementSubjectOperationResponse, err error) {
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
