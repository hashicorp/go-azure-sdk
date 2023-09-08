package directoryadministrativeunit

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

type ListDirectoryAdministrativeUnitsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AdministrativeUnitCollectionResponse
}

type ListDirectoryAdministrativeUnitsCompleteResult struct {
	Items []models.AdministrativeUnitCollectionResponse
}

// ListDirectoryAdministrativeUnits ...
func (c DirectoryAdministrativeUnitClient) ListDirectoryAdministrativeUnits(ctx context.Context) (result ListDirectoryAdministrativeUnitsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/directory/administrativeUnits",
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
		Values *[]models.AdministrativeUnitCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryAdministrativeUnitsComplete retrieves all the results into a single object
func (c DirectoryAdministrativeUnitClient) ListDirectoryAdministrativeUnitsComplete(ctx context.Context) (ListDirectoryAdministrativeUnitsCompleteResult, error) {
	return c.ListDirectoryAdministrativeUnitsCompleteMatchingPredicate(ctx, models.AdministrativeUnitCollectionResponseOperationPredicate{})
}

// ListDirectoryAdministrativeUnitsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryAdministrativeUnitClient) ListDirectoryAdministrativeUnitsCompleteMatchingPredicate(ctx context.Context, predicate models.AdministrativeUnitCollectionResponseOperationPredicate) (result ListDirectoryAdministrativeUnitsCompleteResult, err error) {
	items := make([]models.AdministrativeUnitCollectionResponse, 0)

	resp, err := c.ListDirectoryAdministrativeUnits(ctx)
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

	result = ListDirectoryAdministrativeUnitsCompleteResult{
		Items: items,
	}
	return
}
