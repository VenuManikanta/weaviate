package main

// Import fixtures.
// For now, what is importd, and which credentials are used is hardcoded.
// It'll work just fine with the tools/dev/run_dev_server.sh instance.

// Check out the init() function at the bottom of the file. It load the dataset and schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	spew "github.com/davecgh/go-spew/spew"

	apiclient "github.com/creativesoftwarefdn/weaviate/client"
	"github.com/creativesoftwarefdn/weaviate/client/things"
	"github.com/creativesoftwarefdn/weaviate/models"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const ROOTTOKEN string = "57ac8392-1ecc-4e17-9350-c9c866ac832b"
const ROOTKEY string = "657a48b9-e000-4d9a-b51d-69a0b621c1b9"

// Schema of a weaviate database.
type Schema struct {
	Actions *models.SemanticSchema
	Things  *models.SemanticSchema
}

var schema Schema

type DemoDataset struct {
	Things []DemoInstance `json:"Things"`
	//Actions []DemoInstance `json:"Actions"`
}

type DemoInstance map[string]interface{}

var client *apiclient.WeaviateDecentralisedKnowledgeGraph
var auth runtime.ClientAuthInfoWriterFunc

var demoDataset DemoDataset

func main() {
	createThings()
}

func createThings() {
	type fixupAddRef struct {
		fromId       string
		fromProperty string
		toClass      string
		toId         string
	}

	// need to fix up data later.
	fixups := []fixupAddRef{}

	// Map ID's to UUID's
	idMap := map[string]string{}

	for _, thing := range demoDataset.Things {
		className := thing["class"].(string)
		uuid := thing["uuid"].(string)

		properties := map[string]interface{}{}

		for key, value := range thing {
			if key == "class" || key == "uuid" {
				continue
			}

			ref, isRef := value.(map[string]interface{})
			if isRef {
				fixups = append(fixups, fixupAddRef{
					fromId:       uuid,
					fromProperty: key,
					toClass:      ref["class"].(string),
					toId:         ref["uuid"].(string),
				})
			} else {
				class := findClass(schema.Things, className)
				property := findProperty(class, key)
				if len(property.AtDataType) != 1 {
					panic(fmt.Sprintf("Only one datatype supported for import. Failed in thing %s.%s", className, property.Name))
				}
				dataType := property.AtDataType[0]

				switch dataType {
				case "string":
					properties[key] = value
				case "int":
					v, err := strconv.ParseInt(value.(string), 10, 64)
					if err != nil {
						panic(err)
					}
					properties[key] = v
				case "number":
					properties[key] = value.(float64)
				case "boolean":
					properties[key] = value.(bool)
				default:
					panic(fmt.Sprintf("No such datatype supported: %s", dataType))
				}
			}
		}

		t := models.ThingCreate{
			AtContext: "http://example.org",
			AtClass:   className,
			Schema:    properties,
		}

		thing := assertCreateThing(&t)
		idMap[uuid] = string(thing.ThingID) // Store mapping of ID's
		fmt.Printf("Created Thing %s\n", thing.ThingID)
	}

	// The things might not be persisted in the database yet.
	// So as a _very_ _ugly_ _hack_, we sleep for 1 second to give the weaviate instance enough time to persist everything to the backing database.
	time.Sleep(10 * time.Second)

	// Now fix up refs
	op := "add"
	for _, fixup := range fixups {
		path := fmt.Sprintf("/schema/%s", fixup.fromProperty)
		patch := &models.PatchDocument{
			Op:   &op,
			Path: &path,
			Value: map[string]interface{}{
				"$cref":       idMap[fixup.toId],
				"locationUrl": "http://localhost:8080",
				//"type":        fixup.toClass,
				"type": "Thing",
			},
		}

		assertPatchThing(idMap[fixup.fromId], patch)
		fmt.Printf("Patched %s\n", idMap[fixup.fromId])
	}
}

func findClass(schema *models.SemanticSchema, className string) *models.SemanticSchemaClass {
	for _, klass := range schema.Classes {
		if klass.Class == className {
			return klass
		}
	}

	panic(fmt.Sprintf("Class %s not found", className))
}

func findProperty(class *models.SemanticSchemaClass, propName string) *models.SemanticSchemaClassProperty {
	for _, prop := range class.Properties {
		if prop.Name == propName {
			return prop
		}
	}

	panic(fmt.Sprintf("property %s in class %s not found", propName, class.Class))
}

func assertCreateThing(t *models.ThingCreate) *models.ThingGetResponse {
	params := things.NewWeaviateThingsCreateParams().WithBody(t)

	resp, err := client.Things.WeaviateThingsCreate(params, auth)

	if err != nil {
		switch v := err.(type) {
		case *things.WeaviateThingsCreateUnprocessableEntity:
			panic(fmt.Sprintf("Can't create thing %#v, because %s", t, v.Payload.Error.Message))
		default:
			panic(fmt.Sprintf("Can't create thing %#v, because %#v", t, spew.Sdump(err)))
		}
	}

	return resp.Payload
}

func assertPatchThing(id string, p *models.PatchDocument) *models.ThingGetResponse {
	params := things.NewWeaviateThingsPatchParams().WithBody([]*models.PatchDocument{p}).WithThingID(strfmt.UUID(id))

	resp, err := client.Things.WeaviateThingsPatch(params, auth)

	if err != nil {
		switch v := err.(type) {
		case *things.WeaviateThingsPatchNotFound:
			panic(fmt.Sprintf("Can't patch thing with %s, because thing cannot be found", spew.Sdump(p)))
		case *things.WeaviateThingsPatchUnprocessableEntity:
			panic(fmt.Sprintf("Can't patch thing, because %s", v.Payload.Error.Message))
		default:
			_ = v
			panic(fmt.Sprintf("Can't patch thing with %#v, because %#v", p, spew.Sdump(err)))
		}
	}

	return resp.Payload
}

func loadSemantisSchema(path string) *models.SemanticSchema {
	var schema models.SemanticSchema

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Could not read things sche")
	}

	err = json.Unmarshal(dat, &schema)
	if err != nil {
		panic("Could not load demo dataset")
	}

	return &schema
}

func initSchema() {
	schema.Things = loadSemantisSchema("tools/dev/schema/things_schema.json")
	schema.Actions = loadSemantisSchema("tools/dev/schema/actions_schema.json")
}

func initClient() {
	transport := httptransport.New("localhost:8080", "/weaviate/v1", []string{"http"})
	client = apiclient.New(transport, strfmt.Default)

	auth = func(r runtime.ClientRequest, _ strfmt.Registry) error {
		err := r.SetHeaderParam("X-API-KEY", ROOTKEY)
		if err != nil {
			return err
		}

		return r.SetHeaderParam("X-API-TOKEN", ROOTTOKEN)
	}
}

func initDemoData() {
	dat, err := ioutil.ReadFile("tools/dev/schema/demo_data.json")
	if err != nil {
		panic("Could not read demo dataset")
	}

	err = json.Unmarshal(dat, &demoDataset)
	if err != nil {
		panic("Could not load demo dataset")
	}
}

func init() {
	initSchema()
	initClient()
	initDemoData()
}
