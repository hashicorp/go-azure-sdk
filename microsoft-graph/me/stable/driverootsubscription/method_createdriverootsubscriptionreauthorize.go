package driverootsubscription

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

type CreateDriveRootSubscriptionReauthorizeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateDriveRootSubscriptionReauthorizeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDriveRootSubscriptionReauthorizeOperationOptions() CreateDriveRootSubscriptionReauthorizeOperationOptions {
	return CreateDriveRootSubscriptionReauthorizeOperationOptions{}
}

func (o CreateDriveRootSubscriptionReauthorizeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDriveRootSubscriptionReauthorizeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDriveRootSubscriptionReauthorizeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDriveRootSubscriptionReauthorize - Invoke action reauthorize. Reauthorize a subscription when you receive a
// reauthorizationRequired challenge.
func (c DriveRootSubscriptionClient) CreateDriveRootSubscriptionReauthorize(ctx context.Context, id stable.MeDriveIdRootSubscriptionId, options CreateDriveRootSubscriptionReauthorizeOperationOptions) (result CreateDriveRootSubscriptionReauthorizeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/reauthorize", id.ID()),
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
