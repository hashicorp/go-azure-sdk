package identityb2xuserflowuserattributeassignmentuserattribute

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetIdentityB2xUserFlowByIdUserAttributeAssignmentByIdUserAttributeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *models.IdentityUserFlowAttribute
}

// GetIdentityB2xUserFlowByIdUserAttributeAssignmentByIdUserAttribute ...
func (c IdentityB2xUserFlowUserAttributeAssignmentUserAttributeClient) GetIdentityB2xUserFlowByIdUserAttributeAssignmentByIdUserAttribute(ctx context.Context, id IdentityB2xUserFlowUserAttributeAssignmentId) (result GetIdentityB2xUserFlowByIdUserAttributeAssignmentByIdUserAttributeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/userAttribute", id.ID()),
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
