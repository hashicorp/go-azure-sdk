package siteinformationprotectionsensitivitylabelsublabel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ComputeRightsAndInheritanceResult
}

type CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceOperationOptions() CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceOperationOptions {
	return CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceOperationOptions{}
}

func (o CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritance - Invoke action
// computeRightsAndInheritance. Computes the rights and inheritance for sensitivity labels based on the input content
// and labels.
func (c SiteInformationProtectionSensitivityLabelSublabelClient) CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritance(ctx context.Context, id beta.GroupIdSiteIdInformationProtectionSensitivityLabelId, input CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceRequest, options CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceOperationOptions) (result CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/sublabels/computeRightsAndInheritance", id.ID()),
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

	var model beta.ComputeRightsAndInheritanceResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
