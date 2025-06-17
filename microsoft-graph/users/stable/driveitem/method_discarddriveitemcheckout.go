package driveitem

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

type DiscardDriveItemCheckoutOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DiscardDriveItemCheckoutOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDiscardDriveItemCheckoutOperationOptions() DiscardDriveItemCheckoutOperationOptions {
	return DiscardDriveItemCheckoutOperationOptions{}
}

func (o DiscardDriveItemCheckoutOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DiscardDriveItemCheckoutOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DiscardDriveItemCheckoutOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DiscardDriveItemCheckout - Invoke action discardCheckout. Discard the check out of a driveItem. This action releases
// a driveItem resource that was previously checked out. Any changes made to the item while it was checked out are
// discarded. The same user that performed the checkout must discard it. Another alternative is to use application
// permissions.
func (c DriveItemClient) DiscardDriveItemCheckout(ctx context.Context, id stable.UserIdDriveIdItemId, options DiscardDriveItemCheckoutOperationOptions) (result DiscardDriveItemCheckoutOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/discardCheckout", id.ID()),
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
