// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package source

import (
	"context"
	platform "github.com/fastenhealth/fasten-sources/clients/internal/platform"
	models "github.com/fastenhealth/fasten-sources/clients/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
	logrus "github.com/sirupsen/logrus"
	"net/http"
)

type SourceClientNewYorkHotelTradesCouncilAndHotelAssociationOfNewYorkCityIncHealthBenefitsFund struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/e60dd76f-2355-47fe-85cf-f04cc40e0a16/metadata
func GetSourceClientNewYorkHotelTradesCouncilAndHotelAssociationOfNewYorkCityIncHealthBenefitsFund(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientNewYorkHotelTradesCouncilAndHotelAssociationOfNewYorkCityIncHealthBenefitsFund{baseClient}, updatedSourceCred, err
}