package entitlementmanagementconnectedorganizationinternalsponsor

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

type AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
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

// AddEntitlementManagementConnectedOrganizationInternalSponsorRef - Add internalSponsors. Add a user or a group to the
// connected organization's internal sponsors. The internal sponsors are a set of users who can approve requests on
// behalf of other users from that connected organization.
func (c EntitlementManagementConnectedOrganizationInternalSponsorClient) AddEntitlementManagementConnectedOrganizationInternalSponsorRef(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementConnectedOrganizationId, input stable.ReferenceCreate, options AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationOptions) (result AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/internalSponsors/$ref", id.ID()),
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
