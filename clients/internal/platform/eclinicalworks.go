// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package platform

import (
	"context"
	base "github.com/fastenhealth/fasten-sources/clients/internal/base"
	models "github.com/fastenhealth/fasten-sources/clients/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
	logrus "github.com/sirupsen/logrus"
	"net/http"
)

type sourceClientEclinicalworks struct {
	models.SourceClient
}

/*
https://fhir.eclinicalworks.com/ecwopendev/
*/
// https://fhir4.eclinicalworks.com/fhir/r4/JAFJCD/.well-known/smart-configuration
// https://fhir4.eclinicalworks.com/fhir/r4/JAFJCD/metadata
// https://fhir.eclinicalworks.com/ecwopendev/documentation
func GetSourceClientEclinicalworks(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, error) {
	baseClient, err := base.GetSourceClientFHIR401(env, ctx, globalLogger, sourceCreds, testHttpClient...)
	// API requires the following headers for every request
	baseClient.Headers["Accept"] = "application/json+fhir"

	return sourceClientEclinicalworks{baseClient}, err
}

// Operation-PatientEverything is not supported - https://build.fhir.org/operation-patient-everything.html
// Manually processing individual resources
func (c sourceClientEclinicalworks) SyncAll(db models.DatabaseRepository) (models.UpsertSummary, error) {
	supportedResources := append(c.GetUsCoreResources(), []string{}...)
	return c.SyncAllByResourceName(db, supportedResources)
}