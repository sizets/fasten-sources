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

// https://fhir-myrecord.cerner.com/r4/429a252c-d8ed-49c2-a7fe-d64be676fdca/metadata
func GetSourceTheStevenACohenMilitaryFamilyClinicAtTheUniversityOfPennsylvania(env pkg.FastenLighthouseEnvType, clientIdLookup map[pkg.SourceType]string) (models.LighthouseSourceDefinition, error) {
	sourceDef, err := platform.GetSourceCerner(env, clientIdLookup)
	sourceDef.AuthorizationEndpoint = "https://authorization.cerner.com/tenants/429a252c-d8ed-49c2-a7fe-d64be676fdca/protocols/oauth2/profiles/smart-v1/personas/patient/authorize"
	sourceDef.TokenEndpoint = "https://authorization.cerner.com/tenants/429a252c-d8ed-49c2-a7fe-d64be676fdca/protocols/oauth2/profiles/smart-v1/token"
	sourceDef.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"

	sourceDef.Audience = "https://fhir-myrecord.cerner.com/r4/429a252c-d8ed-49c2-a7fe-d64be676fdca"

	sourceDef.ApiEndpointBaseUrl = "https://fhir-myrecord.cerner.com/r4/429a252c-d8ed-49c2-a7fe-d64be676fdca"
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[pkg.SourceTypeTheStevenACohenMilitaryFamilyClinicAtTheUniversityOfPennsylvania]; clientIdOk {
		sourceDef.ClientId = clientId
	}
	sourceDef.RedirectUri = pkg.GetCallbackEndpoint(string(pkg.SourceTypeCerner))

	sourceDef.Display = "The Steven A. Cohen Military Family Clinic At The University Of Pennsylvania"
	sourceDef.SourceType = pkg.SourceTypeTheStevenACohenMilitaryFamilyClinicAtTheUniversityOfPennsylvania
	sourceDef.SecretKeyPrefix = "cerner"

	return sourceDef, err
}