package usersecurityinformationprotectionsensitivitylabel

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

type ListUserByIdSecurityInformationProtectionSensitivityLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SecuritySensitivityLabelCollectionResponse
}

type ListUserByIdSecurityInformationProtectionSensitivityLabelsCompleteResult struct {
	Items []models.SecuritySensitivityLabelCollectionResponse
}

// ListUserByIdSecurityInformationProtectionSensitivityLabels ...
func (c UserSecurityInformationProtectionSensitivityLabelClient) ListUserByIdSecurityInformationProtectionSensitivityLabels(ctx context.Context, id UserId) (result ListUserByIdSecurityInformationProtectionSensitivityLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/security/informationProtection/sensitivityLabels", id.ID()),
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
		Values *[]models.SecuritySensitivityLabelCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdSecurityInformationProtectionSensitivityLabelsComplete retrieves all the results into a single object
func (c UserSecurityInformationProtectionSensitivityLabelClient) ListUserByIdSecurityInformationProtectionSensitivityLabelsComplete(ctx context.Context, id UserId) (ListUserByIdSecurityInformationProtectionSensitivityLabelsCompleteResult, error) {
	return c.ListUserByIdSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx, id, models.SecuritySensitivityLabelCollectionResponseOperationPredicate{})
}

// ListUserByIdSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserSecurityInformationProtectionSensitivityLabelClient) ListUserByIdSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.SecuritySensitivityLabelCollectionResponseOperationPredicate) (result ListUserByIdSecurityInformationProtectionSensitivityLabelsCompleteResult, err error) {
	items := make([]models.SecuritySensitivityLabelCollectionResponse, 0)

	resp, err := c.ListUserByIdSecurityInformationProtectionSensitivityLabels(ctx, id)
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

	result = ListUserByIdSecurityInformationProtectionSensitivityLabelsCompleteResult{
		Items: items,
	}
	return
}
