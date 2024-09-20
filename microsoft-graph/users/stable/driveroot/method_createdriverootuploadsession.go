package driveroot

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

type CreateDriveRootUploadSessionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UploadSession
}

type CreateDriveRootUploadSessionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDriveRootUploadSessionOperationOptions() CreateDriveRootUploadSessionOperationOptions {
	return CreateDriveRootUploadSessionOperationOptions{}
}

func (o CreateDriveRootUploadSessionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDriveRootUploadSessionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDriveRootUploadSessionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDriveRootUploadSession - Invoke action createUploadSession
func (c DriveRootClient) CreateDriveRootUploadSession(ctx context.Context, id stable.UserIdDriveId, input CreateDriveRootUploadSessionRequest, options CreateDriveRootUploadSessionOperationOptions) (result CreateDriveRootUploadSessionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/root/createUploadSession", id.ID()),
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

	var model stable.UploadSession
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
