package sites

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsRestartOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type WebAppsRestartOperationOptions struct {
	SoftRestart *bool
	Synchronous *bool
}

func DefaultWebAppsRestartOperationOptions() WebAppsRestartOperationOptions {
	return WebAppsRestartOperationOptions{}
}

func (o WebAppsRestartOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WebAppsRestartOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o WebAppsRestartOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.SoftRestart != nil {
		out.Append("softRestart", fmt.Sprintf("%v", *o.SoftRestart))
	}
	if o.Synchronous != nil {
		out.Append("synchronous", fmt.Sprintf("%v", *o.Synchronous))
	}
	return &out
}

// WebAppsRestart ...
func (c SitesClient) WebAppsRestart(ctx context.Context, id commonids.AppServiceId, options WebAppsRestartOperationOptions) (result WebAppsRestartOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/restart", id.ID()),
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
