package mesecurityinformationprotectionsensitivitylabel

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

type ListMeSecurityInformationProtectionSensitivityLabelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.SecuritySensitivityLabelCollectionResponse
}

type ListMeSecurityInformationProtectionSensitivityLabelsCompleteResult struct {
	Items []models.SecuritySensitivityLabelCollectionResponse
}

// ListMeSecurityInformationProtectionSensitivityLabels ...
func (c MeSecurityInformationProtectionSensitivityLabelClient) ListMeSecurityInformationProtectionSensitivityLabels(ctx context.Context) (result ListMeSecurityInformationProtectionSensitivityLabelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/security/informationProtection/sensitivityLabels",
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

// ListMeSecurityInformationProtectionSensitivityLabelsComplete retrieves all the results into a single object
func (c MeSecurityInformationProtectionSensitivityLabelClient) ListMeSecurityInformationProtectionSensitivityLabelsComplete(ctx context.Context) (ListMeSecurityInformationProtectionSensitivityLabelsCompleteResult, error) {
	return c.ListMeSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx, models.SecuritySensitivityLabelCollectionResponseOperationPredicate{})
}

// ListMeSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeSecurityInformationProtectionSensitivityLabelClient) ListMeSecurityInformationProtectionSensitivityLabelsCompleteMatchingPredicate(ctx context.Context, predicate models.SecuritySensitivityLabelCollectionResponseOperationPredicate) (result ListMeSecurityInformationProtectionSensitivityLabelsCompleteResult, err error) {
	items := make([]models.SecuritySensitivityLabelCollectionResponse, 0)

	resp, err := c.ListMeSecurityInformationProtectionSensitivityLabels(ctx)
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

	result = ListMeSecurityInformationProtectionSensitivityLabelsCompleteResult{
		Items: items,
	}
	return
}
