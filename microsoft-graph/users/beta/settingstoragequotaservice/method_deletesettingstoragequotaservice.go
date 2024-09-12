package settingstoragequotaservice

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

type DeleteSettingStorageQuotaServiceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteSettingStorageQuotaServiceOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteSettingStorageQuotaServiceOperationOptions() DeleteSettingStorageQuotaServiceOperationOptions {
	return DeleteSettingStorageQuotaServiceOperationOptions{}
}

func (o DeleteSettingStorageQuotaServiceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteSettingStorageQuotaServiceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteSettingStorageQuotaServiceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteSettingStorageQuotaService - Delete navigation property services for users
func (c SettingStorageQuotaServiceClient) DeleteSettingStorageQuotaService(ctx context.Context, id beta.UserIdSettingStorageQuotaServiceId, options DeleteSettingStorageQuotaServiceOperationOptions) (result DeleteSettingStorageQuotaServiceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
