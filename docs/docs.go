// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mukmin Life API Support",
            "email": "hello@mukmin.life"
        },
        "license": {
            "name": "Proprietary"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/prayer_time/{date}": {
            "get": {
                "description": "get prayer time by date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List prayer times of the day",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Date",
                        "name": "date",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/prayer_time.PrayerTimeOutput"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "prayer_time.PrayerTimeOutput": {
            "type": "object",
            "properties": {
                "asr": {
                    "type": "string"
                },
                "dhuhr": {
                    "type": "string"
                },
                "fajr": {
                    "type": "string"
                },
                "imsak": {
                    "type": "string"
                },
                "isha": {
                    "type": "string"
                },
                "maghrib": {
                    "type": "string"
                },
                "syuruk": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "Mukmin Life API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
