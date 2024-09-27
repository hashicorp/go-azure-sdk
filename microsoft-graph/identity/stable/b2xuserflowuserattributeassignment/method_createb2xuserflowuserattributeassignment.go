package b2xuserflowuserattributeassignment

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

type CreateB2xUserFlowUserAttributeAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityUserFlowAttributeAssignment
}

type CreateB2xUserFlowUserAttributeAssignmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateB2xUserFlowUserAttributeAssignmentOperationOptions() CreateB2xUserFlowUserAttributeAssignmentOperationOptions {
	return CreateB2xUserFlowUserAttributeAssignmentOperationOptions{}
}

func (o CreateB2xUserFlowUserAttributeAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateB2xUserFlowUserAttributeAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateB2xUserFlowUserAttributeAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateB2xUserFlowUserAttributeAssignment - Create userAttributeAssignments. Create a new
// identityUserFlowAttributeAssignment object in a b2xIdentityUserFlow.
func (c B2xUserFlowUserAttributeAssignmentClient) CreateB2xUserFlowUserAttributeAssignment(ctx context.Context, id stable.IdentityB2xUserFlowId, input stable.IdentityUserFlowAttributeAssignment, options CreateB2xUserFlowUserAttributeAssignmentOperationOptions) (result CreateB2xUserFlowUserAttributeAssignmentOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/userAttributeAssignments", id.ID()),
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

	var model stable.IdentityUserFlowAttributeAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
