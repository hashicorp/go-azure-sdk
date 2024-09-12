package entitlementmanagementconnectedorganizationinternalsponsor

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

type AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// AddEntitlementManagementConnectedOrganizationInternalSponsorRef - Add internalSponsors. Add a user or a group to the
// connected organization's internal sponsors. The internal sponsors are a set of users who can approve requests on
// behalf of other users from that connected organization.
func (c EntitlementManagementConnectedOrganizationInternalSponsorClient) AddEntitlementManagementConnectedOrganizationInternalSponsorRef(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementConnectedOrganizationId, input stable.ReferenceCreate) (result AddEntitlementManagementConnectedOrganizationInternalSponsorRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/internalSponsors/$ref", id.ID()),
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
