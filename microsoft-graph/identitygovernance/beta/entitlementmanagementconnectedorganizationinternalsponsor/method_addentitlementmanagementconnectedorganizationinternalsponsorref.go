package entitlementmanagementconnectedorganizationinternalsponsor

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

type AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultAddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationOptions() AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationOptions {
	return AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationOptions{}
}

func (o AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddEntitlementManagementConnectedOrganizationInternalSponsorRef - Add connected organization internal sponsor. Add a
// user or a group to the connected organization's internal sponsors. The internal sponsors are a set of users who can
// approve requests on behalf of other users from that connected organization.
func (c EntitlementManagementConnectedOrganizationInternalSponsorClient) AddEntitlementManagementConnectedOrganizationInternalSponsorRef(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementConnectedOrganizationId, input beta.ReferenceCreate, options AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationOptions) (result AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/internalSponsors/$ref", id.ID()),
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
