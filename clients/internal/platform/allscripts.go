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

type SourceClientAllscripts struct {
	models.SourceClient
}

/*
https://allscripts.vanillacommunities.com/search?query=sandbox&scope=site&source=community
https://open.allscripts.com/fhirendpoints

Allscripts is not actually a confidential client (no client_secret present), however the token endpoint does not support CORS,
so we need to swap the code for the access_token on the server
*/
// https://fhir.fhirpoint.open.allscripts.com/fhirroute/open/CustProProdSand201SMART/metadata
// https://developer.veradigm.com/Fhir/FHIR_Sandboxes#pehr
func GetSourceClientAllscripts(env pkg.FastenLighthouseEnvType, ctx context.Context, globalLogger logrus.FieldLogger, sourceCreds models.SourceCredential, testHttpClient ...*http.Client) (models.SourceClient, *models.SourceCredential, error) {
	baseClient, updatedSourceCred, err := base.GetSourceClientFHIR401(env, ctx, globalLogger, sourceCreds, testHttpClient...)
	// API requires the following headers for every request
	baseClient.Headers["Accept"] = "application/json+fhir"

	return SourceClientAllscripts{baseClient}, updatedSourceCred, err
}

// Operation-PatientEverything is not supported - https://build.fhir.org/operation-patient-everything.html
// Manually processing individual resources
func (c SourceClientAllscripts) SyncAll(db models.DatabaseRepository) (models.UpsertSummary, error) {
	supportedResources := append(c.GetUsCoreResources(), []string{"Account", "Appointment", "Consent", "FamilyMemberHistory", "InsurancePlan", "MedicationRequest", "NutritionOrder", "Person", "Provenance", "Questionnaire", "QuestionnaireResponse", "RelatedPerson", "Schedule", "ServiceRequest", "Slot"}...)
	return c.SyncAllByResourceName(db, supportedResources)
}