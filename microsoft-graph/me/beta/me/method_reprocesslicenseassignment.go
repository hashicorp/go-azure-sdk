package me

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReprocessLicenseAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.User
}

type ReprocessLicenseAssignmentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultReprocessLicenseAssignmentOperationOptions() ReprocessLicenseAssignmentOperationOptions {
	return ReprocessLicenseAssignmentOperationOptions{}
}

func (o ReprocessLicenseAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReprocessLicenseAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ReprocessLicenseAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ReprocessLicenseAssignment - Invoke action reprocessLicenseAssignment. Reprocess all group-based license assignments
// for the user. To learn more about group-based licensing, see What is group-based licensing in Microsoft Entra ID.
// Also see Identify and resolve license assignment problems for a group in Microsoft Entra ID for more details.
func (c MeClient) ReprocessLicenseAssignment(ctx context.Context, options ReprocessLicenseAssignmentOperationOptions) (result ReprocessLicenseAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/reprocessLicenseAssignment",
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

	var model beta.User
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
