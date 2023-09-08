package directoryoutboundshareduserprofiletenant

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDirectoryOutboundSharedUserProfileByIdTenantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TenantReferenceCollectionResponse
}

type ListDirectoryOutboundSharedUserProfileByIdTenantsCompleteResult struct {
	Items []models.TenantReferenceCollectionResponse
}

// ListDirectoryOutboundSharedUserProfileByIdTenants ...
func (c DirectoryOutboundSharedUserProfileTenantClient) ListDirectoryOutboundSharedUserProfileByIdTenants(ctx context.Context, id DirectoryOutboundSharedUserProfileId) (result ListDirectoryOutboundSharedUserProfileByIdTenantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/tenants", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]models.TenantReferenceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryOutboundSharedUserProfileByIdTenantsComplete retrieves all the results into a single object
func (c DirectoryOutboundSharedUserProfileTenantClient) ListDirectoryOutboundSharedUserProfileByIdTenantsComplete(ctx context.Context, id DirectoryOutboundSharedUserProfileId) (ListDirectoryOutboundSharedUserProfileByIdTenantsCompleteResult, error) {
	return c.ListDirectoryOutboundSharedUserProfileByIdTenantsCompleteMatchingPredicate(ctx, id, models.TenantReferenceCollectionResponseOperationPredicate{})
}

// ListDirectoryOutboundSharedUserProfileByIdTenantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryOutboundSharedUserProfileTenantClient) ListDirectoryOutboundSharedUserProfileByIdTenantsCompleteMatchingPredicate(ctx context.Context, id DirectoryOutboundSharedUserProfileId, predicate models.TenantReferenceCollectionResponseOperationPredicate) (result ListDirectoryOutboundSharedUserProfileByIdTenantsCompleteResult, err error) {
	items := make([]models.TenantReferenceCollectionResponse, 0)

	resp, err := c.ListDirectoryOutboundSharedUserProfileByIdTenants(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListDirectoryOutboundSharedUserProfileByIdTenantsCompleteResult{
		Items: items,
	}
	return
}
