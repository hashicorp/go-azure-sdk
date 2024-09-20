package appmanagementpolicy

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

type RemoveAppManagementPolicyRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveAppManagementPolicyRefsOperationOptions struct {
	Id       *string
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultRemoveAppManagementPolicyRefsOperationOptions() RemoveAppManagementPolicyRefsOperationOptions {
	return RemoveAppManagementPolicyRefsOperationOptions{}
}

func (o RemoveAppManagementPolicyRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveAppManagementPolicyRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveAppManagementPolicyRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Id != nil {
		out.Append("@id", fmt.Sprintf("%v", *o.Id))
	}
	return &out
}

// RemoveAppManagementPolicyRefs - Remove appliesTo. Remove an appManagementPolicy policy object from an application or
// service principal object. When you remove the appManagementPolicy, the application or service principal adopts the
// tenant-wide tenantAppManagementPolicy setting.
func (c AppManagementPolicyClient) RemoveAppManagementPolicyRefs(ctx context.Context, id beta.ApplicationId, options RemoveAppManagementPolicyRefsOperationOptions) (result RemoveAppManagementPolicyRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/appManagementPolicies/$ref", id.ID()),
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
