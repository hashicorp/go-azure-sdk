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

type RemoveMobileAppManagementPolicyIncludedGroupRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveMobileAppManagementPolicyIncludedGroupRefOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveMobileAppManagementPolicyIncludedGroupRefOperationOptions() RemoveMobileAppManagementPolicyIncludedGroupRefOperationOptions {
	return RemoveMobileAppManagementPolicyIncludedGroupRefOperationOptions{}
}

func (o RemoveMobileAppManagementPolicyIncludedGroupRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveMobileAppManagementPolicyIncludedGroupRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveMobileAppManagementPolicyIncludedGroupRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveMobileAppManagementPolicyIncludedGroupRef - Delete includedGroup. Delete a group from the list of groups
// included in a mobile app management policy.
func (c MobileAppManagementPolicyIncludedGroupClient) RemoveMobileAppManagementPolicyIncludedGroupRef(ctx context.Context, id beta.PolicyMobileAppManagementPolicyIdIncludedGroupId, options RemoveMobileAppManagementPolicyIncludedGroupRefOperationOptions) (result RemoveMobileAppManagementPolicyIncludedGroupRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$ref", id.ID()),
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

	return
}
