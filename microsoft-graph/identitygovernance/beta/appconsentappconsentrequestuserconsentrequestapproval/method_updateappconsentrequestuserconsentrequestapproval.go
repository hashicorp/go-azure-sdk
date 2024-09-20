package appconsentappconsentrequestuserconsentrequestapproval

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

type UpdateAppConsentRequestUserConsentRequestApprovalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAppConsentRequestUserConsentRequestApprovalOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateAppConsentRequestUserConsentRequestApprovalOperationOptions() UpdateAppConsentRequestUserConsentRequestApprovalOperationOptions {
	return UpdateAppConsentRequestUserConsentRequestApprovalOperationOptions{}
}

func (o UpdateAppConsentRequestUserConsentRequestApprovalOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAppConsentRequestUserConsentRequestApprovalOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAppConsentRequestUserConsentRequestApprovalOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAppConsentRequestUserConsentRequestApproval - Update the navigation property approval in identityGovernance
func (c AppConsentAppConsentRequestUserConsentRequestApprovalClient) UpdateAppConsentRequestUserConsentRequestApproval(ctx context.Context, id beta.IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId, input beta.Approval, options UpdateAppConsentRequestUserConsentRequestApprovalOperationOptions) (result UpdateAppConsentRequestUserConsentRequestApprovalOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/approval", id.ID()),
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
