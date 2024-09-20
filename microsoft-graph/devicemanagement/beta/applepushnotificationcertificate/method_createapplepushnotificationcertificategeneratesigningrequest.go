package applepushnotificationcertificate

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateApplePushNotificationCertificateGenerateSigningRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CreateApplePushNotificationCertificateGenerateSigningRequestResult
}

type CreateApplePushNotificationCertificateGenerateSigningRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateApplePushNotificationCertificateGenerateSigningRequestOperationOptions() CreateApplePushNotificationCertificateGenerateSigningRequestOperationOptions {
	return CreateApplePushNotificationCertificateGenerateSigningRequestOperationOptions{}
}

func (o CreateApplePushNotificationCertificateGenerateSigningRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateApplePushNotificationCertificateGenerateSigningRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateApplePushNotificationCertificateGenerateSigningRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateApplePushNotificationCertificateGenerateSigningRequest - Invoke action
// generateApplePushNotificationCertificateSigningRequest. Download Apple push notification certificate signing request
func (c ApplePushNotificationCertificateClient) CreateApplePushNotificationCertificateGenerateSigningRequest(ctx context.Context, options CreateApplePushNotificationCertificateGenerateSigningRequestOperationOptions) (result CreateApplePushNotificationCertificateGenerateSigningRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/applePushNotificationCertificate/generateApplePushNotificationCertificateSigningRequest",
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

	var model CreateApplePushNotificationCertificateGenerateSigningRequestResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
