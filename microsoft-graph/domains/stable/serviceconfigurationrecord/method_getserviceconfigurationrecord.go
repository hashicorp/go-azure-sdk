package serviceconfigurationrecord

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

type GetServiceConfigurationRecordOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.DomainDnsRecord
}

type GetServiceConfigurationRecordOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetServiceConfigurationRecordOperationOptions() GetServiceConfigurationRecordOperationOptions {
	return GetServiceConfigurationRecordOperationOptions{}
}

func (o GetServiceConfigurationRecordOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetServiceConfigurationRecordOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetServiceConfigurationRecordOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetServiceConfigurationRecord - Get serviceConfigurationRecords from domains. DNS records the customer adds to the
// DNS zone file of the domain before the domain can be used by Microsoft Online services. Read-only, Nullable. Supports
// $expand.
func (c ServiceConfigurationRecordClient) GetServiceConfigurationRecord(ctx context.Context, id stable.DomainIdServiceConfigurationRecordId, options GetServiceConfigurationRecordOperationOptions) (result GetServiceConfigurationRecordOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalDomainDnsRecordImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
