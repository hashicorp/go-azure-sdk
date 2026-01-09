package entitlementmanagementconnectedorganizationexternalsponsor

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

type AddEntitlementManagementConnectedOrganizationExternalSponsorRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddEntitlementManagementConnectedOrganizationExternalSponsorRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddEntitlementManagementConnectedOrganizationExternalSponsorRefOperationOptions() AddEntitlementManagementConnectedOrganizationExternalSponsorRefOperationOptions {
	return AddEntitlementManagementConnectedOrganizationExternalSponsorRefOperationOptions{}
}

func (o AddEntitlementManagementConnectedOrganizationExternalSponsorRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddEntitlementManagementConnectedOrganizationExternalSponsorRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddEntitlementManagementConnectedOrganizationExternalSponsorRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddEntitlementManagementConnectedOrganizationExternalSponsorRef - Add connected organization external sponsor. Add a
// user or a group to the connected organization's external sponsors. The external sponsors are a set of users who can
// approve requests on behalf of other users from that connected organization.
func (c EntitlementManagementConnectedOrganizationExternalSponsorClient) AddEntitlementManagementConnectedOrganizationExternalSponsorRef(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementConnectedOrganizationId, input beta.ReferenceCreate, options AddEntitlementManagementConnectedOrganizationExternalSponsorRefOperationOptions) (result AddEntitlementManagementConnectedOrganizationExternalSponsorRefOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/externalSponsors/$ref", id.ID()),
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
