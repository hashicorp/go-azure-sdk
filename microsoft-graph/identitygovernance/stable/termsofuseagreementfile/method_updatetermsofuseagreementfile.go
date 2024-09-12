package termsofuseagreementfile

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTermsOfUseAgreementFileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateTermsOfUseAgreementFile - Update the navigation property files in identityGovernance
func (c TermsOfUseAgreementFileClient) UpdateTermsOfUseAgreementFile(ctx context.Context, id stable.IdentityGovernanceTermsOfUseAgreementIdFileId, input stable.AgreementFileLocalization) (result UpdateTermsOfUseAgreementFileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPatch,
		Path:       id.ID(),
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
