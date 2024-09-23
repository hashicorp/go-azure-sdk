package androidmanagedstoreaccountenterprisesetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateOperationOptions() SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateOperationOptions {
	return SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateOperationOptions{}
}

func (o SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentState - Invoke action
// setAndroidDeviceOwnerFullyManagedEnrollmentState. Sets the AndroidManagedStoreAccountEnterpriseSettings
// AndroidDeviceOwnerFullyManagedEnrollmentEnabled to the given value.
func (c AndroidManagedStoreAccountEnterpriseSettingClient) SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentState(ctx context.Context, input SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateRequest, options SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateOperationOptions) (result SetAndroidManagedStoreAccountEnterpriseSettingsAndroidDeviceOwnerFullyManagedEnrollmentStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/androidManagedStoreAccountEnterpriseSettings/setAndroidDeviceOwnerFullyManagedEnrollmentState",
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
