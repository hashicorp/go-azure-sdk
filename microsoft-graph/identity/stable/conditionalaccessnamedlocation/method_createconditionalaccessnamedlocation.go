package conditionalaccessnamedlocation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateConditionalAccessNamedLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.NamedLocation
}

type CreateConditionalAccessNamedLocationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateConditionalAccessNamedLocationOperationOptions() CreateConditionalAccessNamedLocationOperationOptions {
	return CreateConditionalAccessNamedLocationOperationOptions{}
}

func (o CreateConditionalAccessNamedLocationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateConditionalAccessNamedLocationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateConditionalAccessNamedLocationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateConditionalAccessNamedLocation - Create namedLocation. Create a new namedLocation object. Named locations can
// be either ipNamedLocation or countryNamedLocation objects.
func (c ConditionalAccessNamedLocationClient) CreateConditionalAccessNamedLocation(ctx context.Context, input stable.NamedLocation, options CreateConditionalAccessNamedLocationOperationOptions) (result CreateConditionalAccessNamedLocationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/identity/conditionalAccess/namedLocations",
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalNamedLocationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
