package userinformationprotectionsensitivitylabel

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

type ListUserByIdInformationProtectionSensitivityLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SensitivityLabelCollectionResponse
}

type ListUserByIdInformationProtectionSensitivityLabelsCompleteResult struct {
	Items []models.SensitivityLabelCollectionResponse
}

// ListUserByIdInformationProtectionSensitivityLabels ...
func (c UserInformationProtectionSensitivityLabelClient) ListUserByIdInformationProtectionSensitivityLabels(ctx context.Context, id UserId) (result ListUserByIdInformationProtectionSensitivityLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/informationProtection/sensitivityLabels", id.ID()),
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

// ListUserByIdInformationProtectionSensitivityLabelsComplete retrieves all the results into a single object
func (c UserInformationProtectionSensitivityLabelClient) ListUserByIdInformationProtectionSensitivityLabelsComplete(ctx context.Context, id UserId) (ListUserByIdInformationProtectionSensitivityLabelsCompleteResult, error) {
	return c.ListUserByIdInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx, id, models.SensitivityLabelCollectionResponseOperationPredicate{})
}

// ListUserByIdInformationProtectionSensitivityLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInformationProtectionSensitivityLabelClient) ListUserByIdInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.SensitivityLabelCollectionResponseOperationPredicate) (result ListUserByIdInformationProtectionSensitivityLabelsCompleteResult, err error) {
	items := make([]models.SensitivityLabelCollectionResponse, 0)

	resp, err := c.ListUserByIdInformationProtectionSensitivityLabels(ctx, id)
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

	result = ListUserByIdInformationProtectionSensitivityLabelsCompleteResult{
		Items: items,
	}
	return
}
