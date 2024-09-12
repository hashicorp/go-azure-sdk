package mobiledevicemanagementpolicyincludedgroup

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

type RemoveMobileDeviceManagementPolicyIncludedGroupRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions struct {
	Id      *string
	IfMatch *string
}

func DefaultRemoveMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions() RemoveMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions {
	return RemoveMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions{}
}

func (o RemoveMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RemoveMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Id != nil {
		out.Append("@id", fmt.Sprintf("%v", *o.Id))
	}
	return &out
}

// RemoveMobileDeviceManagementPolicyIncludedGroupRefs - Delete includedGroup. Delete a group from the list of groups
// included in a mobile device management policy.
func (c MobileDeviceManagementPolicyIncludedGroupClient) RemoveMobileDeviceManagementPolicyIncludedGroupRefs(ctx context.Context, id beta.PolicyMobileDeviceManagementPolicyId, options RemoveMobileDeviceManagementPolicyIncludedGroupRefsOperationOptions) (result RemoveMobileDeviceManagementPolicyIncludedGroupRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/includedGroups/$ref", id.ID()),
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
