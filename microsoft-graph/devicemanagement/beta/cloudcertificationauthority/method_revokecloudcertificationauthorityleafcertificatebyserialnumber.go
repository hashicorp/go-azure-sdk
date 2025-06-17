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

type RevokeCloudCertificationAuthorityLeafCertificateBySerialNumberOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CloudCertificationAuthorityLeafCertificate
}

type RevokeCloudCertificationAuthorityLeafCertificateBySerialNumberOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRevokeCloudCertificationAuthorityLeafCertificateBySerialNumberOperationOptions() RevokeCloudCertificationAuthorityLeafCertificateBySerialNumberOperationOptions {
	return RevokeCloudCertificationAuthorityLeafCertificateBySerialNumberOperationOptions{}
}

func (o RevokeCloudCertificationAuthorityLeafCertificateBySerialNumberOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RevokeCloudCertificationAuthorityLeafCertificateBySerialNumberOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RevokeCloudCertificationAuthorityLeafCertificateBySerialNumberOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RevokeCloudCertificationAuthorityLeafCertificateBySerialNumber - Invoke action revokeLeafCertificateBySerialNumber
func (c CloudCertificationAuthorityClient) RevokeCloudCertificationAuthorityLeafCertificateBySerialNumber(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, input RevokeCloudCertificationAuthorityLeafCertificateBySerialNumberRequest, options RevokeCloudCertificationAuthorityLeafCertificateBySerialNumberOperationOptions) (result RevokeCloudCertificationAuthorityLeafCertificateBySerialNumberOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/revokeLeafCertificateBySerialNumber", id.ID()),
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

	var model beta.CloudCertificationAuthorityLeafCertificate
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
