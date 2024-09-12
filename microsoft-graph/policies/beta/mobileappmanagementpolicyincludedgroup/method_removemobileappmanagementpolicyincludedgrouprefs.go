package mobileappmanagementpolicyincludedgroup

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

type RemoveMobileAppManagementPolicyIncludedGroupRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveMobileAppManagementPolicyIncludedGroupRefsOperationOptions struct {
	Id      *string
	IfMatch *string
}

func DefaultRemoveMobileAppManagementPolicyIncludedGroupRefsOperationOptions() RemoveMobileAppManagementPolicyIncludedGroupRefsOperationOptions {
	return RemoveMobileAppManagementPolicyIncludedGroupRefsOperationOptions{}
}

func (o RemoveMobileAppManagementPolicyIncludedGroupRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveMobileAppManagementPolicyIncludedGroupRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RemoveMobileAppManagementPolicyIncludedGroupRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Id != nil {
		out.Append("@id", fmt.Sprintf("%v", *o.Id))
	}
	return &out
}

// RemoveMobileAppManagementPolicyIncludedGroupRefs - Delete includedGroup. Delete a group from the list of groups
// included in a mobile app management policy.
func (c MobileAppManagementPolicyIncludedGroupClient) RemoveMobileAppManagementPolicyIncludedGroupRefs(ctx context.Context, id beta.PolicyMobileAppManagementPolicyId, options RemoveMobileAppManagementPolicyIncludedGroupRefsOperationOptions) (result RemoveMobileAppManagementPolicyIncludedGroupRefsOperationResponse, err error) {
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
