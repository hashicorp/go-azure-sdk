package appconsentrequestsforapprovaluserconsentrequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAppConsentRequestsForApprovalUserConsentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAppConsentRequestsForApprovalUserConsentRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateAppConsentRequestsForApprovalUserConsentRequestOperationOptions() UpdateAppConsentRequestsForApprovalUserConsentRequestOperationOptions {
	return UpdateAppConsentRequestsForApprovalUserConsentRequestOperationOptions{}
}

func (o UpdateAppConsentRequestsForApprovalUserConsentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAppConsentRequestsForApprovalUserConsentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAppConsentRequestsForApprovalUserConsentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAppConsentRequestsForApprovalUserConsentRequest - Update the navigation property userConsentRequests in me
func (c AppConsentRequestsForApprovalUserConsentRequestClient) UpdateAppConsentRequestsForApprovalUserConsentRequest(ctx context.Context, id beta.MeAppConsentRequestsForApprovalIdUserConsentRequestId, input beta.UserConsentRequest, options UpdateAppConsentRequestsForApprovalUserConsentRequestOperationOptions) (result UpdateAppConsentRequestsForApprovalUserConsentRequestOperationResponse, err error) {
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
