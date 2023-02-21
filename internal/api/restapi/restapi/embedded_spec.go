// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "microservice for komiac",
    "title": "komiac",
    "version": "1.3.0"
  },
  "basePath": "/v1",
  "paths": {
    "/application": {
      "get": {
        "description": "Get applications by filters",
        "tags": [
          "Application"
        ],
        "operationId": "getList",
        "parameters": [
          {
            "type": "string",
            "name": "divisionOID",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "year",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "MNN",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/applications"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "Add new and update existing applications",
        "tags": [
          "Application"
        ],
        "operationId": "addList",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/applications"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/application/{uuid}": {
      "delete": {
        "description": "Delete application",
        "tags": [
          "Application"
        ],
        "operationId": "delete",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "uuid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "application": {
      "type": "object",
      "required": [
        "UUID",
        "medicalOrganizationOID",
        "divisionOID",
        "year",
        "SMNN",
        "MNN",
        "form",
        "dosage",
        "consumerUnit",
        "itemUnit",
        "privilegeProgramCode",
        "privilegeProgram"
      ],
      "properties": {
        "MNN": {
          "type": "string"
        },
        "SMNN": {
          "type": "string"
        },
        "UUID": {
          "type": "string",
          "format": "uuid"
        },
        "consumerUnit": {
          "type": "string"
        },
        "divisionOID": {
          "type": "string"
        },
        "dosage": {
          "type": "string"
        },
        "form": {
          "type": "string"
        },
        "itemUnit": {
          "type": "string"
        },
        "medicalOrganizationOID": {
          "type": "string"
        },
        "privilegeProgram": {
          "type": "string"
        },
        "privilegeProgramCode": {
          "type": "string"
        },
        "year": {
          "type": "integer"
        }
      }
    },
    "applications": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/application"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "microservice for komiac",
    "title": "komiac",
    "version": "1.3.0"
  },
  "basePath": "/v1",
  "paths": {
    "/application": {
      "get": {
        "description": "Get applications by filters",
        "tags": [
          "Application"
        ],
        "operationId": "getList",
        "parameters": [
          {
            "type": "string",
            "name": "divisionOID",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "name": "year",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "MNN",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/applications"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "Add new and update existing applications",
        "tags": [
          "Application"
        ],
        "operationId": "addList",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/applications"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/application/{uuid}": {
      "delete": {
        "description": "Delete application",
        "tags": [
          "Application"
        ],
        "operationId": "delete",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "uuid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "application": {
      "type": "object",
      "required": [
        "UUID",
        "medicalOrganizationOID",
        "divisionOID",
        "year",
        "SMNN",
        "MNN",
        "form",
        "dosage",
        "consumerUnit",
        "itemUnit",
        "privilegeProgramCode",
        "privilegeProgram"
      ],
      "properties": {
        "MNN": {
          "type": "string"
        },
        "SMNN": {
          "type": "string"
        },
        "UUID": {
          "type": "string",
          "format": "uuid"
        },
        "consumerUnit": {
          "type": "string"
        },
        "divisionOID": {
          "type": "string"
        },
        "dosage": {
          "type": "string"
        },
        "form": {
          "type": "string"
        },
        "itemUnit": {
          "type": "string"
        },
        "medicalOrganizationOID": {
          "type": "string"
        },
        "privilegeProgram": {
          "type": "string"
        },
        "privilegeProgramCode": {
          "type": "string"
        },
        "year": {
          "type": "integer"
        }
      }
    },
    "applications": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/application"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
}
