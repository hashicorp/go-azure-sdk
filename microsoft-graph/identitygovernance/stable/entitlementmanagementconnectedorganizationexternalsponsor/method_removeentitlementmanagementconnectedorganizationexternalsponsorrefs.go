package entitlementmanagementconnectedorganizationexternalsponsor

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

type RemoveEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions struct {
	Id        *string
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions() RemoveEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions {
	return RemoveEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions{}
}

func (o RemoveEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Id != nil {
		out.Append("@id", fmt.Sprintf("%v", *o.Id))
	}
	return &out
}

// RemoveEntitlementManagementConnectedOrganizationExternalSponsorRefs - Remove externalSponsors. Remove a user or a
// group from the connected organization's external sponsors. The external sponsors are a set of users who can approve
// requests on behalf of other users from that connected organization.
func (c EntitlementManagementConnectedOrganizationExternalSponsorClient) RemoveEntitlementManagementConnectedOrganizationExternalSponsorRefs(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementConnectedOrganizationId, options RemoveEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationOptions) (result RemoveEntitlementManagementConnectedOrganizationExternalSponsorRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/externalSponsors/$ref", id.ID()),
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
