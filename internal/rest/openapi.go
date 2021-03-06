package rest

import "github.com/getkin/kin-openapi/openapi3"

func newOpenAPI() *openapi3.T {
	swagger := &openapi3.T{
		OpenAPI:    "3.0.0",
		Components: openapi3.Components{},
		Info: &openapi3.Info{
			ExtensionProps: openapi3.ExtensionProps{},
			Title:          "Theater API",
			Version:        "0.1",
		},
		Paths:    map[string]*openapi3.PathItem{},
		Security: []openapi3.SecurityRequirement{},
		Servers: []*openapi3.Server{
			{
				URL:         "http://localhost:8000/",
				Description: "Local dev",
			},
		},
	}

	swagger.Components.Schemas = openapi3.Schemas{
		"Tag": openapi3.NewSchemaRef(
			"", openapi3.NewObjectSchema().WithProperties(
				map[string]*openapi3.Schema{
					"id":   openapi3.NewIntegerSchema(),
					"name": openapi3.NewStringSchema(),
				},
			),
		),
		"Cloth": openapi3.NewSchemaRef(
			"", openapi3.NewObjectSchema().WithProperties(
				map[string]*openapi3.Schema{
					"id":   openapi3.NewIntegerSchema(),
					"name": openapi3.NewStringSchema(),
					"type": openapi3.NewStringSchema(),
				},
			),
		),
		"Costume": openapi3.NewSchemaRef(
			"", openapi3.NewObjectSchema().WithProperties(
				map[string]*openapi3.Schema{
					"id":          openapi3.NewIntegerSchema(),
					"name":        openapi3.NewStringSchema(),
					"description": openapi3.NewStringSchema().WithNullable(),
					"clothes": openapi3.NewArraySchema().WithItems(openapi3.NewObjectSchema().WithProperties(
						map[string]*openapi3.Schema{
							"id":   openapi3.NewIntegerSchema(),
							"name": openapi3.NewStringSchema(),
							"type": openapi3.NewStringSchema(),
						},
					)),
					"tags": openapi3.NewObjectSchema().WithProperties(
						map[string]*openapi3.Schema{
							"colors":    openapi3.NewArraySchema().WithItems(openapi3.NewStringSchema()),
							"materials": openapi3.NewArraySchema().WithItems(openapi3.NewStringSchema()),
						},
					),
					"images": openapi3.NewObjectSchema().WithProperties(
						map[string]*openapi3.Schema{
							"front":   openapi3.NewStringSchema().WithNullable(),
							"back":    openapi3.NewStringSchema().WithNullable(),
							"sideway": openapi3.NewStringSchema().WithNullable(),
							"details": openapi3.NewStringSchema().WithNullable(),
						},
					).WithNullable(),
					"location":   openapi3.NewStringSchema(),
					"designer":   openapi3.NewStringSchema(),
					"size":       openapi3.NewIntegerSchema(),
					"isArchived": openapi3.NewBoolSchema().WithNullable(),
					"isDecor":    openapi3.NewBoolSchema(),
					"condition":  openapi3.NewStringSchema().WithEnum("????????????", "????????????????????", "??????????????"),
				},
			),
		),
		"Performance": openapi3.NewSchemaRef(
			"", openapi3.NewObjectSchema().WithProperties(
				map[string]*openapi3.Schema{
					"id":       openapi3.NewIntegerSchema(),
					"name":     openapi3.NewStringSchema(),
					"location": openapi3.NewStringSchema(),
					"startAt":  openapi3.NewDateTimeSchema(),
					"duration": openapi3.NewIntegerSchema(),
					"costumes": openapi3.NewArraySchema().WithItems(openapi3.NewObjectSchema().
						WithPropertyRef("", &openapi3.SchemaRef{Ref: "#/components/schemas/Costume"})),
				},
			),
		),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"CreateUpdateClothRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("???????? ?????????????? ?????? ???????????????? ?? ???????????????????? ???????????????? ??????????????").
				WithRequired(true).
				WithJSONSchema(openapi3.NewObjectSchema().WithProperties(
					map[string]*openapi3.Schema{
						"name":   openapi3.NewStringSchema(),
						"typeId": openapi3.NewIntegerSchema(),
					},
				)),
		},
		"CreateUpdateCostumeRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("???????? ?????????????? ?????? ???????????????? ?? ???????????????????? ??????????????").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().WithProperties(
					map[string]*openapi3.Schema{
						"name":        openapi3.NewStringSchema(),
						"location":    openapi3.NewStringSchema().WithNullable(),
						"description": openapi3.NewStringSchema().WithNullable(),
						"clothes":     openapi3.NewArraySchema().WithItems(openapi3.NewIntegerSchema()),
						"images": openapi3.NewObjectSchema().WithProperties(
							map[string]*openapi3.Schema{
								"front":   openapi3.NewStringSchema().WithNullable(),
								"back":    openapi3.NewStringSchema().WithNullable(),
								"sideway": openapi3.NewStringSchema().WithNullable(),
								"details": openapi3.NewStringSchema().WithNullable(),
							},
						),
						"designer":   openapi3.NewStringSchema().WithNullable(),
						"condition":  openapi3.NewStringSchema(),
						"size":       openapi3.NewIntegerSchema(),
						"isArchived": openapi3.NewBoolSchema().WithNullable(),
						"isDecor":    openapi3.NewBoolSchema(),
					},
				))),
		},
		"MakeWriteOffsRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("???????? ?????????????? ?????? ???????????????? ????????????????").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().WithProperty(
					"id", openapi3.NewArraySchema().WithItems(openapi3.NewIntegerSchema()))),
				),
		},
		"CreateTagsRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("???????? ?????????????? ???? ???????????????? ????????????/????????????????????").
				WithContent(
					openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().WithProperty(
						"names", openapi3.NewArraySchema().WithItems(openapi3.NewStringSchema())),
					),
				),
		},
		"UpdateTagRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("???????? ?????????????? ???? ???????????????????? ???????????????? ????????").
				WithContent(openapi3.NewContentWithJSONSchema(
					openapi3.NewObjectSchema().
						WithProperty("name", openapi3.NewStringSchema()),
				)),
		},
	}

	swagger.Components.Responses = openapi3.Responses{
		"ErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("?????????? ?????? ?????????? ??????????????").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().WithProperties(
					map[string]*openapi3.Schema{
						"error":  openapi3.NewStringSchema(),
						"result": openapi3.NewAnyOfSchema().WithNullable(),
					},
				))),
		},
		"CreateClothResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("?????????? ?????? ???????????????? ???????????????? ???????????????? ??????????????").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().
					WithProperty("error", openapi3.NewStringSchema().WithNullable()).
					WithPropertyRef("result", &openapi3.SchemaRef{Ref: "#/components/schemas/Cloth"}))),
		},
		"CreateCostumeResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("?????????? ?????? ???????????????? ???????????????? ??????????????").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().
					WithProperty("error", openapi3.NewStringSchema().WithNullable()).
					WithPropertyRef("result", &openapi3.SchemaRef{Ref: "#/components/schemas/Costume"}))),
		},
		"GetClothesByPageResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("?????????? ?????? ?????????????? ???? ?????????????????? ?????????????????? ???????????????? ???? ????????????????").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().
					WithProperty("error", openapi3.NewAnyOfSchema().WithNullable()).
					WithProperty("result", openapi3.NewArraySchema().WithItems(swagger.Components.Schemas["Cloth"].Value)))),
		},
		"GetCostumesByPageResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("?????????? ?????? ?????????????? ???? ?????????????????? ???????????????? ???? ????????????????").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().
					WithProperty("error", openapi3.NewAnyOfSchema().WithNullable()).
					WithProperty("result", openapi3.NewArraySchema().WithItems(swagger.Components.Schemas["Costume"].Value)))),
		},
		"DeleteResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("?????????? ?????? ???????????????? ????????-????????").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().WithProperties(
					map[string]*openapi3.Schema{
						"error":  openapi3.NewAnyOfSchema().WithNullable(),
						"result": openapi3.NewIntegerSchema(),
					},
				))),
		},
		"MakeWriteOffsResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("?????????? ?????? ???????????????? ????????????????").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().WithProperties(
					map[string]*openapi3.Schema{
						"error":  openapi3.NewAnyOfSchema().WithNullable(),
						"result": openapi3.NewStringSchema(),
					},
				))),
		},
		"TagsResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("?????????? ?????? ???????????????? ???????????????? ??????????").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().
					WithProperty("error", openapi3.NewAnyOfSchema().WithNullable()).
					WithProperty("result", openapi3.NewArraySchema().WithItems(swagger.Components.Schemas["Tag"].Value)))),
		},
		"TagUpdateResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("?????????? ?????? ???????????????????? ???????????????? ????????").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewObjectSchema().
					WithProperty("error", openapi3.NewAnyOfSchema().WithNullable()).
					WithProperty("result", openapi3.NewArraySchema().WithItems(swagger.Components.Schemas["Tag"].Value)))),
		},
	}

	swagger.Paths = openapi3.Paths{
		// ???????????????? ????????????????
		"/api/v1/clothes": &openapi3.PathItem{
			Post: &openapi3.Operation{
				OperationID: "CreateCloth",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CreateUpdateClothRequest",
				},
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/CreateClothResponse"},
				},
			},
		},
		"/api/v1/clothes/pages/{page}": &openapi3.PathItem{
			Parameters: []*openapi3.ParameterRef{
				{Value: openapi3.NewPathParameter("page").WithSchema(openapi3.NewIntegerSchema())},
			},
			Get: &openapi3.Operation{
				OperationID: "GetClothesByPage",
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/GetClothesByPageResponse"},
				},
			},
		},
		"/api/v1/clothes/{id}": &openapi3.PathItem{
			Parameters: []*openapi3.ParameterRef{
				{Value: openapi3.NewPathParameter("id").WithSchema(openapi3.NewIntegerSchema())},
			},
			Put: &openapi3.Operation{
				OperationID: "UpdateCloth",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CreateUpdateClothRequest",
				},
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/CreateClothResponse"},
				},
			},
			Delete: &openapi3.Operation{
				OperationID: "DeleteCloth",
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/DeleteResponse"},
				},
			},
		},

		// ??????????????
		"/api/v1/costumes": &openapi3.PathItem{
			Post: &openapi3.Operation{
				OperationID: "CreateCostume",
				RequestBody: &openapi3.RequestBodyRef{Ref: "#/components/requestBodies/CreateUpdateCostumeRequest"},
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/CreateCostumeResponse"},
				},
			},
		},
		"/api/v1/costumes/pages/{page}": &openapi3.PathItem{
			Parameters: []*openapi3.ParameterRef{
				{Value: openapi3.NewPathParameter("page").WithSchema(openapi3.NewIntegerSchema())},
			},
			Get: &openapi3.Operation{
				OperationID: "GetCostumesByPage",
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/GetCostumesByPageResponse"},
				},
			},
		},
		"/api/v1/costumes/{id}": &openapi3.PathItem{
			Parameters: []*openapi3.ParameterRef{
				{Value: openapi3.NewPathParameter("id").WithSchema(openapi3.NewIntegerSchema())},
			},
			Put: &openapi3.Operation{
				OperationID: "UpdateCostume",
				RequestBody: &openapi3.RequestBodyRef{Ref: "#/components/requestBodies/CreateUpdateCostumeRequest"},
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/CreateCostumeResponse"},
				},
			},
			Delete: &openapi3.Operation{
				OperationID: "DeleteCostume",
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/DeleteResponse"},
				},
			},
		},

		// ?????????? ?????????????????? ??????????????
		"/api/v1/tags/colors": &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "GetColors",
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/TagsResponse"},
				},
			},
			Post: &openapi3.Operation{
				OperationID: "CreateColors",
				RequestBody: &openapi3.RequestBodyRef{Ref: "#/components/requestBodies/CreateTagsRequest"},
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/TagsResponse"},
				},
			},
		},
		"/api/v1/tags/colors/{id}": &openapi3.PathItem{
			Parameters: []*openapi3.ParameterRef{
				{
					Value: openapi3.NewPathParameter("id").WithSchema(openapi3.NewIntegerSchema()),
				},
			},
			Put: &openapi3.Operation{
				OperationID: "UpdateColorName",
				RequestBody: &openapi3.RequestBodyRef{Ref: "#/components/requestBodies/UpdateTagRequest"},
				Responses: openapi3.Responses{
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/TagUpdateResponse"},
				},
			},
			Delete: &openapi3.Operation{
				OperationID: "DeleteColor",
				Responses: openapi3.Responses{
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/DeleteResponse"},
				},
			},
		},

		// ?????????????????? ?????????????????? ??????????????
		"/api/v1/tags/materials": &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "GetMaterials",
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/TagsResponse"},
				},
			},
			Post: &openapi3.Operation{
				OperationID: "CreateMaterials",
				RequestBody: &openapi3.RequestBodyRef{Ref: "#/components/requestBodies/CreateTagsRequest"},
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/TagsResponse"},
				},
			},
		},
		"/api/v1/tags/materials/{id}": &openapi3.PathItem{
			Parameters: []*openapi3.ParameterRef{
				{Value: openapi3.NewPathParameter("id").WithSchema(openapi3.NewIntegerSchema())},
			},
			Put: &openapi3.Operation{
				OperationID: "UpdateMaterialName",
				RequestBody: &openapi3.RequestBodyRef{Ref: "#/components/requestBodies/UpdateTagRequest"},
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/TagUpdateResponse"},
				},
			},
			Delete: &openapi3.Operation{
				OperationID: "DeleteMaterial",
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/DeleteResponse"},
				},
			},
		},

		// ???????? ?????????????????? ??????????????
		"/api/v1/types": &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "GetClothesTypes",
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/TagsResponse"},
				},
			},
			Post: &openapi3.Operation{
				OperationID: "CreateTypes",
				RequestBody: &openapi3.RequestBodyRef{Ref: "#/components/requestBodies/CreateTagsRequest"},
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/TagsResponse"},
				},
			},
		},
		"/api/v1/types/{id}": &openapi3.PathItem{
			Parameters: []*openapi3.ParameterRef{
				{Value: openapi3.NewPathParameter("id").WithSchema(openapi3.NewIntegerSchema())},
			},
			Put: &openapi3.Operation{
				OperationID: "UpdateClothesTypeName",
				RequestBody: &openapi3.RequestBodyRef{Ref: "#/components/requestBodies/UpdateTagRequest"},
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/TagUpdateResponse"},
				},
			},
			Delete: &openapi3.Operation{
				OperationID: "DeleteClothType",
				Responses: openapi3.Responses{
					"422": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"500": &openapi3.ResponseRef{Ref: "#/components/responses/ErrorResponse"},
					"200": &openapi3.ResponseRef{Ref: "#/components/responses/DeleteResponse"},
				},
			},
		},
	}

	return swagger
}
