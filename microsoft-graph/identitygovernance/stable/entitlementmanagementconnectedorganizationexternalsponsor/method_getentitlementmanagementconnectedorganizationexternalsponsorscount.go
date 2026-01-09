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

type GetEntitlementManagementConnectedOrganizationExternalSponsorsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetEntitlementManagementConnectedOrganizationExternalSponsorsCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetEntitlementManagementConnectedOrganizationExternalSponsorsCountOperationOptions() GetEntitlementManagementConnectedOrganizationExternalSponsorsCountOperationOptions {
	return GetEntitlementManagementConnectedOrganizationExternalSponsorsCountOperationOptions{}
}

func (o GetEntitlementManagementConnectedOrganizationExternalSponsorsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementConnectedOrganizationExternalSponsorsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetEntitlementManagementConnectedOrganizationExternalSponsorsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementConnectedOrganizationExternalSponsorsCount - Get the number of the resource
func (c EntitlementManagementConnectedOrganizationExternalSponsorClient) GetEntitlementManagementConnectedOrganizationExternalSponsorsCount(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementConnectedOrganizationId, options GetEntitlementManagementConnectedOrganizationExternalSponsorsCountOperationOptions) (result GetEntitlementManagementConnectedOrganizationExternalSponsorsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/externalSponsors/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
