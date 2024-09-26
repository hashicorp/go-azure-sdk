package microsofttunnelsite

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

type RequestMicrosoftTunnelSiteUpgradeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RequestMicrosoftTunnelSiteUpgradeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRequestMicrosoftTunnelSiteUpgradeOperationOptions() RequestMicrosoftTunnelSiteUpgradeOperationOptions {
	return RequestMicrosoftTunnelSiteUpgradeOperationOptions{}
}

func (o RequestMicrosoftTunnelSiteUpgradeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RequestMicrosoftTunnelSiteUpgradeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RequestMicrosoftTunnelSiteUpgradeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RequestMicrosoftTunnelSiteUpgrade - Invoke action requestUpgrade
func (c MicrosoftTunnelSiteClient) RequestMicrosoftTunnelSiteUpgrade(ctx context.Context, id beta.DeviceManagementMicrosoftTunnelSiteId, options RequestMicrosoftTunnelSiteUpgradeOperationOptions) (result RequestMicrosoftTunnelSiteUpgradeOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/requestUpgrade", id.ID()),
		RetryFunc:     options.RetryFunc,
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
