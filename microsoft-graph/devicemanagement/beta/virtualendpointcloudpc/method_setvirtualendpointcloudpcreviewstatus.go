package virtualendpointcloudpc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetVirtualEndpointCloudPCReviewStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetVirtualEndpointCloudPCReviewStatusOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSetVirtualEndpointCloudPCReviewStatusOperationOptions() SetVirtualEndpointCloudPCReviewStatusOperationOptions {
	return SetVirtualEndpointCloudPCReviewStatusOperationOptions{}
}

func (o SetVirtualEndpointCloudPCReviewStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetVirtualEndpointCloudPCReviewStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetVirtualEndpointCloudPCReviewStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetVirtualEndpointCloudPCReviewStatus - Invoke action setReviewStatus. Set the review status of a specific Cloud PC
// device using the Cloud PC ID. Use this API to set the review status of a Cloud PC to in review if you consider a
// Cloud PC suspicious. After the review is completed, use this API again to set the Cloud PC back to a normal state.
func (c VirtualEndpointCloudPCClient) SetVirtualEndpointCloudPCReviewStatus(ctx context.Context, id beta.DeviceManagementVirtualEndpointCloudPCId, input SetVirtualEndpointCloudPCReviewStatusRequest, options SetVirtualEndpointCloudPCReviewStatusOperationOptions) (result SetVirtualEndpointCloudPCReviewStatusOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/setReviewStatus", id.ID()),
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
