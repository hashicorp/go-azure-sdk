package meinformationprotectionsensitivitylabelsublabel

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

type ListMeInformationProtectionSensitivityLabelByIdSublabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SensitivityLabelCollectionResponse
}

type ListMeInformationProtectionSensitivityLabelByIdSublabelsCompleteResult struct {
	Items []models.SensitivityLabelCollectionResponse
}

// ListMeInformationProtectionSensitivityLabelByIdSublabels ...
func (c MeInformationProtectionSensitivityLabelSublabelClient) ListMeInformationProtectionSensitivityLabelByIdSublabels(ctx context.Context, id MeInformationProtectionSensitivityLabelId) (result ListMeInformationProtectionSensitivityLabelByIdSublabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/sublabels", id.ID()),
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
		Values *[]models.SensitivityLabelCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeInformationProtectionSensitivityLabelByIdSublabelsComplete retrieves all the results into a single object
func (c MeInformationProtectionSensitivityLabelSublabelClient) ListMeInformationProtectionSensitivityLabelByIdSublabelsComplete(ctx context.Context, id MeInformationProtectionSensitivityLabelId) (ListMeInformationProtectionSensitivityLabelByIdSublabelsCompleteResult, error) {
	return c.ListMeInformationProtectionSensitivityLabelByIdSublabelsCompleteMatchingPredicate(ctx, id, models.SensitivityLabelCollectionResponseOperationPredicate{})
}

// ListMeInformationProtectionSensitivityLabelByIdSublabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeInformationProtectionSensitivityLabelSublabelClient) ListMeInformationProtectionSensitivityLabelByIdSublabelsCompleteMatchingPredicate(ctx context.Context, id MeInformationProtectionSensitivityLabelId, predicate models.SensitivityLabelCollectionResponseOperationPredicate) (result ListMeInformationProtectionSensitivityLabelByIdSublabelsCompleteResult, err error) {
	items := make([]models.SensitivityLabelCollectionResponse, 0)

	resp, err := c.ListMeInformationProtectionSensitivityLabelByIdSublabels(ctx, id)
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

	result = ListMeInformationProtectionSensitivityLabelByIdSublabelsCompleteResult{
		Items: items,
	}
	return
}
