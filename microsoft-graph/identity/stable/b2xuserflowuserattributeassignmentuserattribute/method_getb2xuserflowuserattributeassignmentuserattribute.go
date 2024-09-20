package b2xuserflowuserattributeassignmentuserattribute

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetB2xUserFlowUserAttributeAssignmentUserAttributeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.IdentityUserFlowAttribute
}

type GetB2xUserFlowUserAttributeAssignmentUserAttributeOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetB2xUserFlowUserAttributeAssignmentUserAttributeOperationOptions() GetB2xUserFlowUserAttributeAssignmentUserAttributeOperationOptions {
	return GetB2xUserFlowUserAttributeAssignmentUserAttributeOperationOptions{}
}

func (o GetB2xUserFlowUserAttributeAssignmentUserAttributeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetB2xUserFlowUserAttributeAssignmentUserAttributeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetB2xUserFlowUserAttributeAssignmentUserAttributeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetB2xUserFlowUserAttributeAssignmentUserAttribute - Get userAttribute from identity. The user attribute that you
// want to add to your user flow.
func (c B2xUserFlowUserAttributeAssignmentUserAttributeClient) GetB2xUserFlowUserAttributeAssignmentUserAttribute(ctx context.Context, id stable.IdentityB2xUserFlowIdUserAttributeAssignmentId, options GetB2xUserFlowUserAttributeAssignmentUserAttributeOperationOptions) (result GetB2xUserFlowUserAttributeAssignmentUserAttributeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/userAttribute", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalIdentityUserFlowAttributeImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
