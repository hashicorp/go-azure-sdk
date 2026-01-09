package mobileappmanagementpolicyincludedgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddMobileAppManagementPolicyIncludedGroupRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddMobileAppManagementPolicyIncludedGroupRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddMobileAppManagementPolicyIncludedGroupRefOperationOptions() AddMobileAppManagementPolicyIncludedGroupRefOperationOptions {
	return AddMobileAppManagementPolicyIncludedGroupRefOperationOptions{}
}

func (o AddMobileAppManagementPolicyIncludedGroupRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddMobileAppManagementPolicyIncludedGroupRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddMobileAppManagementPolicyIncludedGroupRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddMobileAppManagementPolicyIncludedGroupRef - Add includedGroups. Add groups to be included in a mobile app
// management policy.
func (c MobileAppManagementPolicyIncludedGroupClient) AddMobileAppManagementPolicyIncludedGroupRef(ctx context.Context, id beta.PolicyMobileAppManagementPolicyId, input beta.ReferenceCreate, options AddMobileAppManagementPolicyIncludedGroupRefOperationOptions) (result AddMobileAppManagementPolicyIncludedGroupRefOperationResponse, err error) {
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
