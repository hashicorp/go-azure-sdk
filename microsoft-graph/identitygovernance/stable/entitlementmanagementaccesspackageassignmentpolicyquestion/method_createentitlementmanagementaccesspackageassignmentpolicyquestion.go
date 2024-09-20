package entitlementmanagementaccesspackageassignmentpolicyquestion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.AccessPackageQuestion
}

type CreateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions() CreateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions {
	return CreateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageAssignmentPolicyQuestion - Create new navigation property to questions for
// identityGovernance
func (c EntitlementManagementAccessPackageAssignmentPolicyQuestionClient) CreateEntitlementManagementAccessPackageAssignmentPolicyQuestion(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId, input stable.AccessPackageQuestion, options CreateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions) (result CreateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/questions", id.ID()),
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
	model, err := stable.UnmarshalAccessPackageQuestionImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
