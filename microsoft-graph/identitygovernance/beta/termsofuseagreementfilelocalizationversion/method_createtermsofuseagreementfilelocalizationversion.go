package termsofuseagreementfilelocalizationversion

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

type CreateTermsOfUseAgreementFileLocalizationVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AgreementFileVersion
}

type CreateTermsOfUseAgreementFileLocalizationVersionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTermsOfUseAgreementFileLocalizationVersionOperationOptions() CreateTermsOfUseAgreementFileLocalizationVersionOperationOptions {
	return CreateTermsOfUseAgreementFileLocalizationVersionOperationOptions{}
}

func (o CreateTermsOfUseAgreementFileLocalizationVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTermsOfUseAgreementFileLocalizationVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTermsOfUseAgreementFileLocalizationVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTermsOfUseAgreementFileLocalizationVersion - Create new navigation property to versions for identityGovernance
func (c TermsOfUseAgreementFileLocalizationVersionClient) CreateTermsOfUseAgreementFileLocalizationVersion(ctx context.Context, id beta.IdentityGovernanceTermsOfUseAgreementIdFileLocalizationId, input beta.AgreementFileVersion, options CreateTermsOfUseAgreementFileLocalizationVersionOperationOptions) (result CreateTermsOfUseAgreementFileLocalizationVersionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/versions", id.ID()),
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

	var model beta.AgreementFileVersion
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
