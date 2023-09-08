package meinformationprotectionsensitivitylabel

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

type ListMeInformationProtectionSensitivityLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SensitivityLabelCollectionResponse
}

type ListMeInformationProtectionSensitivityLabelsCompleteResult struct {
	Items []models.SensitivityLabelCollectionResponse
}

// ListMeInformationProtectionSensitivityLabels ...
func (c MeInformationProtectionSensitivityLabelClient) ListMeInformationProtectionSensitivityLabels(ctx context.Context) (result ListMeInformationProtectionSensitivityLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/informationProtection/sensitivityLabels",
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

// ListMeInformationProtectionSensitivityLabelsComplete retrieves all the results into a single object
func (c MeInformationProtectionSensitivityLabelClient) ListMeInformationProtectionSensitivityLabelsComplete(ctx context.Context) (ListMeInformationProtectionSensitivityLabelsCompleteResult, error) {
	return c.ListMeInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx, models.SensitivityLabelCollectionResponseOperationPredicate{})
}

// ListMeInformationProtectionSensitivityLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeInformationProtectionSensitivityLabelClient) ListMeInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx context.Context, predicate models.SensitivityLabelCollectionResponseOperationPredicate) (result ListMeInformationProtectionSensitivityLabelsCompleteResult, err error) {
	items := make([]models.SensitivityLabelCollectionResponse, 0)

	resp, err := c.ListMeInformationProtectionSensitivityLabels(ctx)
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

	result = ListMeInformationProtectionSensitivityLabelsCompleteResult{
		Items: items,
	}
	return
}
