package androidmanagedstoreaccountenterprisesetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAndroidManagedStoreAccountEnterpriseSettingUnbindOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAndroidManagedStoreAccountEnterpriseSettingUnbindOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateAndroidManagedStoreAccountEnterpriseSettingUnbindOperationOptions() CreateAndroidManagedStoreAccountEnterpriseSettingUnbindOperationOptions {
	return CreateAndroidManagedStoreAccountEnterpriseSettingUnbindOperationOptions{}
}

func (o CreateAndroidManagedStoreAccountEnterpriseSettingUnbindOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAndroidManagedStoreAccountEnterpriseSettingUnbindOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAndroidManagedStoreAccountEnterpriseSettingUnbindOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAndroidManagedStoreAccountEnterpriseSettingUnbind - Invoke action unbind
func (c AndroidManagedStoreAccountEnterpriseSettingClient) CreateAndroidManagedStoreAccountEnterpriseSettingUnbind(ctx context.Context, options CreateAndroidManagedStoreAccountEnterpriseSettingUnbindOperationOptions) (result CreateAndroidManagedStoreAccountEnterpriseSettingUnbindOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/androidManagedStoreAccountEnterpriseSettings/unbind",
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

	return
}
