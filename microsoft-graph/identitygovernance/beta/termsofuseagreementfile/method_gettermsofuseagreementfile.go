package termsofuseagreementfile

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTermsOfUseAgreementFileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AgreementFileLocalization
}

type GetTermsOfUseAgreementFileOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetTermsOfUseAgreementFileOperationOptions() GetTermsOfUseAgreementFileOperationOptions {
	return GetTermsOfUseAgreementFileOperationOptions{}
}

func (o GetTermsOfUseAgreementFileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTermsOfUseAgreementFileOperationOptions) ToOData() *odata.Query {
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

func (o GetTermsOfUseAgreementFileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTermsOfUseAgreementFile - Get files from identityGovernance. PDFs linked to this agreement. Note: This property is
// in the process of being deprecated. Use the file property instead.
func (c TermsOfUseAgreementFileClient) GetTermsOfUseAgreementFile(ctx context.Context, id beta.IdentityGovernanceTermsOfUseAgreementIdFileId, options GetTermsOfUseAgreementFileOperationOptions) (result GetTermsOfUseAgreementFileOperationResponse, err error) {
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

	var model beta.AgreementFileLocalization
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
