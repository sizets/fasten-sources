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

type SourceClientTheStevenACohenMilitaryFamilyClinicAtTheUpCenter struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/c4f2491a-95ed-4fe6-a71d-a2eb4e47373a/metadata
func GetSourceClientTheStevenACohenMilitaryFamilyClinicAtTheUpCenter(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientTheStevenACohenMilitaryFamilyClinicAtTheUpCenter{baseClient}, updatedSourceCred, err
}