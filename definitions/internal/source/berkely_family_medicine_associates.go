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

// https://fhir-myrecord.cerner.com/r4/bf5eb65c-cdc6-4156-8349-5fcd9a13f4a2/metadata
func GetSourceBerkelyFamilyMedicineAssociates(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/bf5eb65c-cdc6-4156-8349-5fcd9a13f4a2/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/bf5eb65c-cdc6-4156-8349-5fcd9a13f4a2/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/bf5eb65c-cdc6-4156-8349-5fcd9a13f4a2"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/bf5eb65c-cdc6-4156-8349-5fcd9a13f4a2"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeBerkelyFamilyMedicineAssociates]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "Berkely Family Medicine Associates"
	sourceDef.SourceType = pkg.SourceTypeBerkelyFamilyMedicineAssociates
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}