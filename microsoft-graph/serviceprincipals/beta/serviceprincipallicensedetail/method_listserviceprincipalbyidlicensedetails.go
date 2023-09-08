package serviceprincipallicensedetail

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

type ListServicePrincipalByIdLicenseDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.LicenseDetailsCollectionResponse
}

type ListServicePrincipalByIdLicenseDetailsCompleteResult struct {
	Items []models.LicenseDetailsCollectionResponse
}

// ListServicePrincipalByIdLicenseDetails ...
func (c ServicePrincipalLicenseDetailClient) ListServicePrincipalByIdLicenseDetails(ctx context.Context, id ServicePrincipalId) (result ListServicePrincipalByIdLicenseDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/licenseDetails", id.ID()),
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
		Values *[]models.LicenseDetailsCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServicePrincipalByIdLicenseDetailsComplete retrieves all the results into a single object
func (c ServicePrincipalLicenseDetailClient) ListServicePrincipalByIdLicenseDetailsComplete(ctx context.Context, id ServicePrincipalId) (ListServicePrincipalByIdLicenseDetailsCompleteResult, error) {
	return c.ListServicePrincipalByIdLicenseDetailsCompleteMatchingPredicate(ctx, id, models.LicenseDetailsCollectionResponseOperationPredicate{})
}

// ListServicePrincipalByIdLicenseDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicePrincipalLicenseDetailClient) ListServicePrincipalByIdLicenseDetailsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate models.LicenseDetailsCollectionResponseOperationPredicate) (result ListServicePrincipalByIdLicenseDetailsCompleteResult, err error) {
	items := make([]models.LicenseDetailsCollectionResponse, 0)

	resp, err := c.ListServicePrincipalByIdLicenseDetails(ctx, id)
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

	result = ListServicePrincipalByIdLicenseDetailsCompleteResult{
		Items: items,
	}
	return
}
