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

// https://fhir.pinehurstsurgical.com/FHIR/metadata
func GetSourcePinehurstSurgicalClinic(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceAllscripts(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://fhir.pinehurstsurgical.com/authorization/connect/authorize"
	sourceDef.TokenEndpoint = "https://fhir.pinehurstsurgical.com/authorization/connect/token"

	sourceDef.Audience = "https://fhir.pinehurstsurgical.com/FHIR"

	sourceDef.ApiEndpointBaseUrl = "https://fhir.pinehurstsurgical.com/FHIR"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypePinehurstSurgicalClinic]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeAllscripts))

	sourceDef.Display = "Pinehurst Surgical Clinic"
	sourceDef.SourceType = pkg.SourceTypePinehurstSurgicalClinic
	sourceDef.Hidden = true
	sourceDef.SecretKeyPrefix = "allscripts"

	return sourceDef, err
}