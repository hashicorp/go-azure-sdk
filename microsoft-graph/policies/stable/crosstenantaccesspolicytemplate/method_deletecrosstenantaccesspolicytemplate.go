package crosstenantaccesspolicytemplate

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteCrossTenantAccessPolicyTemplateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteCrossTenantAccessPolicyTemplateOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteCrossTenantAccessPolicyTemplateOperationOptions() DeleteCrossTenantAccessPolicyTemplateOperationOptions {
	return DeleteCrossTenantAccessPolicyTemplateOperationOptions{}
}

func (o DeleteCrossTenantAccessPolicyTemplateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteCrossTenantAccessPolicyTemplateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteCrossTenantAccessPolicyTemplateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteCrossTenantAccessPolicyTemplate - Delete navigation property templates for policies
func (c CrossTenantAccessPolicyTemplateClient) DeleteCrossTenantAccessPolicyTemplate(ctx context.Context, options DeleteCrossTenantAccessPolicyTemplateOperationOptions) (result DeleteCrossTenantAccessPolicyTemplateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/policies/crossTenantAccessPolicy/templates",
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
