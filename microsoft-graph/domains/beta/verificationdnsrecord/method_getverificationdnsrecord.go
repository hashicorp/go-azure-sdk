package verificationdnsrecord

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVerificationDnsRecordOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DomainDnsRecord
}

type GetVerificationDnsRecordOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetVerificationDnsRecordOperationOptions() GetVerificationDnsRecordOperationOptions {
	return GetVerificationDnsRecordOperationOptions{}
}

func (o GetVerificationDnsRecordOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetVerificationDnsRecordOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetVerificationDnsRecordOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetVerificationDnsRecord - Get verificationDnsRecords from domains. DNS records that the customer adds to the DNS
// zone file of the domain before the customer can complete domain ownership verification with Microsoft Entra ID.
// Read-only, Nullable. Does not support $expand.
func (c VerificationDnsRecordClient) GetVerificationDnsRecord(ctx context.Context, id beta.DomainIdVerificationDnsRecordId, options GetVerificationDnsRecordOperationOptions) (result GetVerificationDnsRecordOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalDomainDnsRecordImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
