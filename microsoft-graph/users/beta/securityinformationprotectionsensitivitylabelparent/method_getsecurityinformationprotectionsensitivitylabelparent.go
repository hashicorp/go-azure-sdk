package securityinformationprotectionsensitivitylabelparent

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

type GetSecurityInformationProtectionSensitivityLabelParentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SecuritySensitivityLabel
}

type GetSecurityInformationProtectionSensitivityLabelParentOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetSecurityInformationProtectionSensitivityLabelParentOperationOptions() GetSecurityInformationProtectionSensitivityLabelParentOperationOptions {
	return GetSecurityInformationProtectionSensitivityLabelParentOperationOptions{}
}

func (o GetSecurityInformationProtectionSensitivityLabelParentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSecurityInformationProtectionSensitivityLabelParentOperationOptions) ToOData() *odata.Query {
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

func (o GetSecurityInformationProtectionSensitivityLabelParentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSecurityInformationProtectionSensitivityLabelParent - Get parent from users. The parent label associated with a
// child label. Null if the label has no parent.
func (c SecurityInformationProtectionSensitivityLabelParentClient) GetSecurityInformationProtectionSensitivityLabelParent(ctx context.Context, id beta.UserIdSecurityInformationProtectionSensitivityLabelId, options GetSecurityInformationProtectionSensitivityLabelParentOperationOptions) (result GetSecurityInformationProtectionSensitivityLabelParentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/parent", id.ID()),
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

	var model beta.SecuritySensitivityLabel
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
