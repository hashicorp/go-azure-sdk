package useragreementacceptance

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

type ListUserByIdAgreementAcceptancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AgreementAcceptanceCollectionResponse
}

type ListUserByIdAgreementAcceptancesCompleteResult struct {
	Items []models.AgreementAcceptanceCollectionResponse
}

// ListUserByIdAgreementAcceptances ...
func (c UserAgreementAcceptanceClient) ListUserByIdAgreementAcceptances(ctx context.Context, id UserId) (result ListUserByIdAgreementAcceptancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/agreementAcceptances", id.ID()),
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
		Values *[]models.AgreementAcceptanceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdAgreementAcceptancesComplete retrieves all the results into a single object
func (c UserAgreementAcceptanceClient) ListUserByIdAgreementAcceptancesComplete(ctx context.Context, id UserId) (ListUserByIdAgreementAcceptancesCompleteResult, error) {
	return c.ListUserByIdAgreementAcceptancesCompleteMatchingPredicate(ctx, id, models.AgreementAcceptanceCollectionResponseOperationPredicate{})
}

// ListUserByIdAgreementAcceptancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAgreementAcceptanceClient) ListUserByIdAgreementAcceptancesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.AgreementAcceptanceCollectionResponseOperationPredicate) (result ListUserByIdAgreementAcceptancesCompleteResult, err error) {
	items := make([]models.AgreementAcceptanceCollectionResponse, 0)

	resp, err := c.ListUserByIdAgreementAcceptances(ctx, id)
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

	result = ListUserByIdAgreementAcceptancesCompleteResult{
		Items: items,
	}
	return
}
