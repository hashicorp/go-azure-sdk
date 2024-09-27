package devicemanagement

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendCustomNotificationToCompanyPortalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendCustomNotificationToCompanyPortalOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSendCustomNotificationToCompanyPortalOperationOptions() SendCustomNotificationToCompanyPortalOperationOptions {
	return SendCustomNotificationToCompanyPortalOperationOptions{}
}

func (o SendCustomNotificationToCompanyPortalOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendCustomNotificationToCompanyPortalOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendCustomNotificationToCompanyPortalOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendCustomNotificationToCompanyPortal - Invoke action sendCustomNotificationToCompanyPortal
func (c DeviceManagementClient) SendCustomNotificationToCompanyPortal(ctx context.Context, input SendCustomNotificationToCompanyPortalRequest, options SendCustomNotificationToCompanyPortalOperationOptions) (result SendCustomNotificationToCompanyPortalOperationResponse, err error) {
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
		Path:          "/deviceManagement/sendCustomNotificationToCompanyPortal",
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
