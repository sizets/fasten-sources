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

type SourceClientMadeliaCommunityHospitalAndClinicDBAMadeliaHealth struct {
	models.SourceClient
}

// https://fhir-myrecord.cerner.com/r4/f2aa0e6f-59f8-4af5-916b-5f2f9445c1cb/metadata
func GetSourceClientMadeliaCommunityHospitalAndClinicDBAMadeliaHealth(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := platform.GetSourceClientCerner(env, ctx, globalLogger, sourceCreds, testHttpClient...)

	return SourceClientMadeliaCommunityHospitalAndClinicDBAMadeliaHealth{baseClient}, updatedSourceCred, err
}