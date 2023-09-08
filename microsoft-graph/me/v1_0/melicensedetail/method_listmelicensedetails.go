package melicensedetail

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

type ListMeLicenseDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.LicenseDetailsCollectionResponse
}

type ListMeLicenseDetailsCompleteResult struct {
	Items []models.LicenseDetailsCollectionResponse
}

// ListMeLicenseDetails ...
func (c MeLicenseDetailClient) ListMeLicenseDetails(ctx context.Context) (result ListMeLicenseDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/licenseDetails",
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

// ListMeLicenseDetailsComplete retrieves all the results into a single object
func (c MeLicenseDetailClient) ListMeLicenseDetailsComplete(ctx context.Context) (ListMeLicenseDetailsCompleteResult, error) {
	return c.ListMeLicenseDetailsCompleteMatchingPredicate(ctx, models.LicenseDetailsCollectionResponseOperationPredicate{})
}

// ListMeLicenseDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeLicenseDetailClient) ListMeLicenseDetailsCompleteMatchingPredicate(ctx context.Context, predicate models.LicenseDetailsCollectionResponseOperationPredicate) (result ListMeLicenseDetailsCompleteResult, err error) {
	items := make([]models.LicenseDetailsCollectionResponse, 0)

	resp, err := c.ListMeLicenseDetails(ctx)
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

	result = ListMeLicenseDetailsCompleteResult{
		Items: items,
	}
	return
}
