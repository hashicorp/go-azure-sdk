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

type CreateCloudCertificationAuthorityUploadExternallySignedCertificateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudCertificationAuthority
}

type CreateCloudCertificationAuthorityUploadExternallySignedCertificateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateCloudCertificationAuthorityUploadExternallySignedCertificateOperationOptions() CreateCloudCertificationAuthorityUploadExternallySignedCertificateOperationOptions {
	return CreateCloudCertificationAuthorityUploadExternallySignedCertificateOperationOptions{}
}

func (o CreateCloudCertificationAuthorityUploadExternallySignedCertificateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCloudCertificationAuthorityUploadExternallySignedCertificateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCloudCertificationAuthorityUploadExternallySignedCertificateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCloudCertificationAuthorityUploadExternallySignedCertificate - Invoke action
// uploadExternallySignedCertificationAuthorityCertificate
func (c CloudCertificationAuthorityClient) CreateCloudCertificationAuthorityUploadExternallySignedCertificate(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, input CreateCloudCertificationAuthorityUploadExternallySignedCertificateRequest, options CreateCloudCertificationAuthorityUploadExternallySignedCertificateOperationOptions) (result CreateCloudCertificationAuthorityUploadExternallySignedCertificateOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/uploadExternallySignedCertificationAuthorityCertificate", id.ID()),
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
