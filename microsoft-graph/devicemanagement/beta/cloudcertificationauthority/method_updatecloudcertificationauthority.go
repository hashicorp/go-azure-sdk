package cloudcertificationauthority

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCloudCertificationAuthorityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCloudCertificationAuthorityOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateCloudCertificationAuthorityOperationOptions() UpdateCloudCertificationAuthorityOperationOptions {
	return UpdateCloudCertificationAuthorityOperationOptions{}
}

func (o UpdateCloudCertificationAuthorityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCloudCertificationAuthorityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCloudCertificationAuthorityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCloudCertificationAuthority - Update the navigation property cloudCertificationAuthority in deviceManagement
func (c CloudCertificationAuthorityClient) UpdateCloudCertificationAuthority(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, input beta.CloudCertificationAuthority, options UpdateCloudCertificationAuthorityOperationOptions) (result UpdateCloudCertificationAuthorityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
