package cloudcertificationauthority

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

type ChangeCloudCertificationAuthorityStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudCertificationAuthority
}

type ChangeCloudCertificationAuthorityStatusOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultChangeCloudCertificationAuthorityStatusOperationOptions() ChangeCloudCertificationAuthorityStatusOperationOptions {
	return ChangeCloudCertificationAuthorityStatusOperationOptions{}
}

func (o ChangeCloudCertificationAuthorityStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ChangeCloudCertificationAuthorityStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ChangeCloudCertificationAuthorityStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ChangeCloudCertificationAuthorityStatus - Invoke action changeCloudCertificationAuthorityStatus
func (c CloudCertificationAuthorityClient) ChangeCloudCertificationAuthorityStatus(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, input ChangeCloudCertificationAuthorityStatusRequest, options ChangeCloudCertificationAuthorityStatusOperationOptions) (result ChangeCloudCertificationAuthorityStatusOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/changeCloudCertificationAuthorityStatus", id.ID()),
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

	var model beta.CloudCertificationAuthority
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
