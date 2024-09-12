package settingstoragequota

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

type DeleteSettingStorageQuotaOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteSettingStorageQuotaOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteSettingStorageQuotaOperationOptions() DeleteSettingStorageQuotaOperationOptions {
	return DeleteSettingStorageQuotaOperationOptions{}
}

func (o DeleteSettingStorageQuotaOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteSettingStorageQuotaOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteSettingStorageQuotaOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteSettingStorageQuota - Delete navigation property quota for users
func (c SettingStorageQuotaClient) DeleteSettingStorageQuota(ctx context.Context, id stable.UserId, options DeleteSettingStorageQuotaOperationOptions) (result DeleteSettingStorageQuotaOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/settings/storage/quota", id.ID()),
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
