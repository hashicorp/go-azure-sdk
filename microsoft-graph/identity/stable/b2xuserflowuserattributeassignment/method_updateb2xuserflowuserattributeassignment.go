package b2xuserflowuserattributeassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateB2xUserFlowUserAttributeAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateB2xUserFlowUserAttributeAssignmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateB2xUserFlowUserAttributeAssignmentOperationOptions() UpdateB2xUserFlowUserAttributeAssignmentOperationOptions {
	return UpdateB2xUserFlowUserAttributeAssignmentOperationOptions{}
}

func (o UpdateB2xUserFlowUserAttributeAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateB2xUserFlowUserAttributeAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateB2xUserFlowUserAttributeAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateB2xUserFlowUserAttributeAssignment - Update identityUserFlowAttributeAssignment. Update the properties of a
// identityUserFlowAttributeAssignment object.
func (c B2xUserFlowUserAttributeAssignmentClient) UpdateB2xUserFlowUserAttributeAssignment(ctx context.Context, id stable.IdentityB2xUserFlowIdUserAttributeAssignmentId, input stable.IdentityUserFlowAttributeAssignment, options UpdateB2xUserFlowUserAttributeAssignmentOperationOptions) (result UpdateB2xUserFlowUserAttributeAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
