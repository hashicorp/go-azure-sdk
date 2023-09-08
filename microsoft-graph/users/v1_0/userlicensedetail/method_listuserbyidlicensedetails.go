package userlicensedetail

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

type ListUserByIdLicenseDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.LicenseDetailsCollectionResponse
}

type ListUserByIdLicenseDetailsCompleteResult struct {
	Items []models.LicenseDetailsCollectionResponse
}

// ListUserByIdLicenseDetails ...
func (c UserLicenseDetailClient) ListUserByIdLicenseDetails(ctx context.Context, id UserId) (result ListUserByIdLicenseDetailsOperationResponse, err error) {
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

// ListUserByIdLicenseDetailsComplete retrieves all the results into a single object
func (c UserLicenseDetailClient) ListUserByIdLicenseDetailsComplete(ctx context.Context, id UserId) (ListUserByIdLicenseDetailsCompleteResult, error) {
	return c.ListUserByIdLicenseDetailsCompleteMatchingPredicate(ctx, id, models.LicenseDetailsCollectionResponseOperationPredicate{})
}

// ListUserByIdLicenseDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserLicenseDetailClient) ListUserByIdLicenseDetailsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.LicenseDetailsCollectionResponseOperationPredicate) (result ListUserByIdLicenseDetailsCompleteResult, err error) {
	items := make([]models.LicenseDetailsCollectionResponse, 0)

	resp, err := c.ListUserByIdLicenseDetails(ctx, id)
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

	result = ListUserByIdLicenseDetailsCompleteResult{
		Items: items,
	}
	return
}
