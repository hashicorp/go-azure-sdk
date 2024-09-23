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

type AddMobileDeviceManagementPolicyIncludedGroupRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddMobileDeviceManagementPolicyIncludedGroupRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddMobileDeviceManagementPolicyIncludedGroupRefOperationOptions() AddMobileDeviceManagementPolicyIncludedGroupRefOperationOptions {
	return AddMobileDeviceManagementPolicyIncludedGroupRefOperationOptions{}
}

func (o AddMobileDeviceManagementPolicyIncludedGroupRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddMobileDeviceManagementPolicyIncludedGroupRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddMobileDeviceManagementPolicyIncludedGroupRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddMobileDeviceManagementPolicyIncludedGroupRef - Add includedGroups. Add groups to be included in a mobile app
// management policy.
func (c MobileDeviceManagementPolicyIncludedGroupClient) AddMobileDeviceManagementPolicyIncludedGroupRef(ctx context.Context, id beta.PolicyMobileDeviceManagementPolicyId, input beta.ReferenceCreate, options AddMobileDeviceManagementPolicyIncludedGroupRefOperationOptions) (result AddMobileDeviceManagementPolicyIncludedGroupRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/includedGroups/$ref", id.ID()),
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
