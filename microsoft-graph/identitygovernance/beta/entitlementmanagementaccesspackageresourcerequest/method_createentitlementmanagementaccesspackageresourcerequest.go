package entitlementmanagementaccesspackageresourcerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementAccessPackageResourceRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResourceRequest
}

// CreateEntitlementManagementAccessPackageResourceRequest - Create accessPackageResourceRequest. Create a new
// accessPackageResourceRequest object to request the addition of a resource to an access package catalog, update of a
// resource, or the removal of a resource from a catalog. A resource must be included in an access package catalog
// before a role of that resource can be added to an access package.
func (c EntitlementManagementAccessPackageResourceRequestClient) CreateEntitlementManagementAccessPackageResourceRequest(ctx context.Context, input beta.AccessPackageResourceRequest) (result CreateEntitlementManagementAccessPackageResourceRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       "/identityGovernance/entitlementManagement/accessPackageResourceRequests",
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

	var model beta.AccessPackageResourceRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
