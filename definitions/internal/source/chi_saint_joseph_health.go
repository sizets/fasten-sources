// Copyright (C) Fasten Health, Inc. - All Rights Reserved.
//
// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-sources-gen
// PLEASE DO NOT EDIT BY HAND

package source

import (
	platform "github.com/fastenhealth/fasten-sources/definitions/internal/platform"
	models "github.com/fastenhealth/fasten-sources/definitions/models"
	pkg "github.com/fastenhealth/fasten-sources/pkg"
)

// https://rpsouth.catholichealth.net/fhir/FHIRKY/api/FHIR/R4/metadata
func GetSourceChiSaintJosephHealth(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceEpic(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://rpsouth.catholichealth.net/fhir/FHIRKY/oauth2/authorize"
	sourceDef.TokenEndpoint = "https://rpsouth.catholichealth.net/fhir/FHIRKY/oauth2/token"

	sourceDef.Audience = "https://rpsouth.catholichealth.net/fhir/FHIRKY/api/FHIR/R4"

	sourceDef.ApiEndpointBaseUrl = "https://rpsouth.catholichealth.net/fhir/FHIRKY/api/FHIR/R4"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeChiSaintJosephHealth]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeEpic))

	sourceDef.Display = "CHI Saint Joseph Health"
	sourceDef.SourceType = pkg.SourceTypeChiSaintJosephHealth
	sourceDef.SecretKeyPrefix = "epic"

	return sourceDef, err
}