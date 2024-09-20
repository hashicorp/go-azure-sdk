package appmanagementpolicy

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

type AddAppManagementPolicyRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddAppManagementPolicyRefOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultAddAppManagementPolicyRefOperationOptions() AddAppManagementPolicyRefOperationOptions {
	return AddAppManagementPolicyRefOperationOptions{}
}

func (o AddAppManagementPolicyRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddAppManagementPolicyRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddAppManagementPolicyRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddAppManagementPolicyRef - Assign appliesTo. Assign an appManagementPolicy policy object to an application or
// service principal object. The application or service principal adopts this policy over the tenant-wide
// tenantAppManagementPolicy setting. Only one policy object can be assigned to an application or service principal.
func (c AppManagementPolicyClient) AddAppManagementPolicyRef(ctx context.Context, id stable.ApplicationId, input stable.ReferenceCreate, options AddAppManagementPolicyRefOperationOptions) (result AddAppManagementPolicyRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/appManagementPolicies/$ref", id.ID()),
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
