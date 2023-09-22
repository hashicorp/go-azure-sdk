package backups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountBackupsListByNetAppAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *BackupsList
}

type AccountBackupsListByNetAppAccountOperationOptions struct {
	IncludeOnlyBackupsFromDeletedVolumes *string
}

func DefaultAccountBackupsListByNetAppAccountOperationOptions() AccountBackupsListByNetAppAccountOperationOptions {
	return AccountBackupsListByNetAppAccountOperationOptions{}
}

func (o AccountBackupsListByNetAppAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AccountBackupsListByNetAppAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o AccountBackupsListByNetAppAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IncludeOnlyBackupsFromDeletedVolumes != nil {
		out.Append("includeOnlyBackupsFromDeletedVolumes", fmt.Sprintf("%v", *o.IncludeOnlyBackupsFromDeletedVolumes))
	}
	return &out
}

// AccountBackupsListByNetAppAccount ...
func (c BackupsClient) AccountBackupsListByNetAppAccount(ctx context.Context, id NetAppAccountId, options AccountBackupsListByNetAppAccountOperationOptions) (result AccountBackupsListByNetAppAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/accountBackups", id.ID()),
		OptionsObject: options,
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
