package mobilethreatdefenseconnector

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMobileThreatDefenseConnectorOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.MobileThreatDefenseConnector
}

type CreateMobileThreatDefenseConnectorOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateMobileThreatDefenseConnectorOperationOptions() CreateMobileThreatDefenseConnectorOperationOptions {
	return CreateMobileThreatDefenseConnectorOperationOptions{}
}

func (o CreateMobileThreatDefenseConnectorOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMobileThreatDefenseConnectorOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMobileThreatDefenseConnectorOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMobileThreatDefenseConnector - Create mobileThreatDefenseConnector. Create a new mobileThreatDefenseConnector
// object.
func (c MobileThreatDefenseConnectorClient) CreateMobileThreatDefenseConnector(ctx context.Context, input stable.MobileThreatDefenseConnector, options CreateMobileThreatDefenseConnectorOperationOptions) (result CreateMobileThreatDefenseConnectorOperationResponse, err error) {
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
		Path:          "/deviceManagement/mobileThreatDefenseConnectors",
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

	var model stable.MobileThreatDefenseConnector
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
