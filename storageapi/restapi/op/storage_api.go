// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/DiaElectronics/lea-central-wash/storageapi"
)

// NewStorageAPI creates a new Storage instance
func NewStorageAPI(spec *loads.Document) *StorageAPI {
	return &StorageAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		AddServiceAmountHandler: AddServiceAmountHandlerFunc(func(params AddServiceAmountParams, principal *storageapi.Profile) AddServiceAmountResponder {
			// return middleware.NotImplemented("operation AddServiceAmount has not yet been implemented")
			return AddServiceAmountNotImplemented()
		}),
		DelStationHandler: DelStationHandlerFunc(func(params DelStationParams) DelStationResponder {
			// return middleware.NotImplemented("operation DelStation has not yet been implemented")
			return DelStationNotImplemented()
		}),
		GetPingHandler: GetPingHandlerFunc(func(params GetPingParams) GetPingResponder {
			// return middleware.NotImplemented("operation GetPing has not yet been implemented")
			return GetPingNotImplemented()
		}),
		InfoHandler: InfoHandlerFunc(func(params InfoParams) InfoResponder {
			// return middleware.NotImplemented("operation Info has not yet been implemented")
			return InfoNotImplemented()
		}),
		KasseHandler: KasseHandlerFunc(func(params KasseParams) KasseResponder {
			// return middleware.NotImplemented("operation Kasse has not yet been implemented")
			return KasseNotImplemented()
		}),
		LoadHandler: LoadHandlerFunc(func(params LoadParams) LoadResponder {
			// return middleware.NotImplemented("operation Load has not yet been implemented")
			return LoadNotImplemented()
		}),
		LoadMoneyHandler: LoadMoneyHandlerFunc(func(params LoadMoneyParams) LoadMoneyResponder {
			// return middleware.NotImplemented("operation LoadMoney has not yet been implemented")
			return LoadMoneyNotImplemented()
		}),
		LoadRelayHandler: LoadRelayHandlerFunc(func(params LoadRelayParams) LoadRelayResponder {
			// return middleware.NotImplemented("operation LoadRelay has not yet been implemented")
			return LoadRelayNotImplemented()
		}),
		OpenStationHandler: OpenStationHandlerFunc(func(params OpenStationParams) OpenStationResponder {
			// return middleware.NotImplemented("operation OpenStation has not yet been implemented")
			return OpenStationNotImplemented()
		}),
		PingHandler: PingHandlerFunc(func(params PingParams) PingResponder {
			// return middleware.NotImplemented("operation Ping has not yet been implemented")
			return PingNotImplemented()
		}),
		ProgramRelaysHandler: ProgramRelaysHandlerFunc(func(params ProgramRelaysParams) ProgramRelaysResponder {
			// return middleware.NotImplemented("operation ProgramRelays has not yet been implemented")
			return ProgramRelaysNotImplemented()
		}),
		ProgramsHandler: ProgramsHandlerFunc(func(params ProgramsParams) ProgramsResponder {
			// return middleware.NotImplemented("operation Programs has not yet been implemented")
			return ProgramsNotImplemented()
		}),
		SaveHandler: SaveHandlerFunc(func(params SaveParams) SaveResponder {
			// return middleware.NotImplemented("operation Save has not yet been implemented")
			return SaveNotImplemented()
		}),
		SaveCollectionHandler: SaveCollectionHandlerFunc(func(params SaveCollectionParams) SaveCollectionResponder {
			// return middleware.NotImplemented("operation SaveCollection has not yet been implemented")
			return SaveCollectionNotImplemented()
		}),
		SaveIfNotExistsHandler: SaveIfNotExistsHandlerFunc(func(params SaveIfNotExistsParams) SaveIfNotExistsResponder {
			// return middleware.NotImplemented("operation SaveIfNotExists has not yet been implemented")
			return SaveIfNotExistsNotImplemented()
		}),
		SaveMoneyHandler: SaveMoneyHandlerFunc(func(params SaveMoneyParams) SaveMoneyResponder {
			// return middleware.NotImplemented("operation SaveMoney has not yet been implemented")
			return SaveMoneyNotImplemented()
		}),
		SaveRelayHandler: SaveRelayHandlerFunc(func(params SaveRelayParams) SaveRelayResponder {
			// return middleware.NotImplemented("operation SaveRelay has not yet been implemented")
			return SaveRelayNotImplemented()
		}),
		SetKasseHandler: SetKasseHandlerFunc(func(params SetKasseParams) SetKasseResponder {
			// return middleware.NotImplemented("operation SetKasse has not yet been implemented")
			return SetKasseNotImplemented()
		}),
		SetProgramNameHandler: SetProgramNameHandlerFunc(func(params SetProgramNameParams) SetProgramNameResponder {
			// return middleware.NotImplemented("operation SetProgramName has not yet been implemented")
			return SetProgramNameNotImplemented()
		}),
		SetProgramRelaysHandler: SetProgramRelaysHandlerFunc(func(params SetProgramRelaysParams) SetProgramRelaysResponder {
			// return middleware.NotImplemented("operation SetProgramRelays has not yet been implemented")
			return SetProgramRelaysNotImplemented()
		}),
		SetStationHandler: SetStationHandlerFunc(func(params SetStationParams) SetStationResponder {
			// return middleware.NotImplemented("operation SetStation has not yet been implemented")
			return SetStationNotImplemented()
		}),
		StationByHashHandler: StationByHashHandlerFunc(func(params StationByHashParams) StationByHashResponder {
			// return middleware.NotImplemented("operation StationByHash has not yet been implemented")
			return StationByHashNotImplemented()
		}),
		StationReportCurrentMoneyHandler: StationReportCurrentMoneyHandlerFunc(func(params StationReportCurrentMoneyParams) StationReportCurrentMoneyResponder {
			// return middleware.NotImplemented("operation StationReportCurrentMoney has not yet been implemented")
			return StationReportCurrentMoneyNotImplemented()
		}),
		StationReportDatesHandler: StationReportDatesHandlerFunc(func(params StationReportDatesParams) StationReportDatesResponder {
			// return middleware.NotImplemented("operation StationReportDates has not yet been implemented")
			return StationReportDatesNotImplemented()
		}),
		StationsVariablesHandler: StationsVariablesHandlerFunc(func(params StationsVariablesParams) StationsVariablesResponder {
			// return middleware.NotImplemented("operation StationsVariables has not yet been implemented")
			return StationsVariablesNotImplemented()
		}),
		StatusHandler: StatusHandlerFunc(func(params StatusParams) StatusResponder {
			// return middleware.NotImplemented("operation Status has not yet been implemented")
			return StatusNotImplemented()
		}),
		StatusCollectionHandler: StatusCollectionHandlerFunc(func(params StatusCollectionParams) StatusCollectionResponder {
			// return middleware.NotImplemented("operation StatusCollection has not yet been implemented")
			return StatusCollectionNotImplemented()
		}),

		// Applies when the "Pin" header is set
		PinCodeAuth: func(token string) (*storageapi.Profile, error) {
			return nil, errors.NotImplemented("api key auth (pinCode) Pin from header param [Pin] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*StorageAPI the storage API */
type StorageAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// PinCodeAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Pin provided in the header
	PinCodeAuth func(string) (*storageapi.Profile, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// AddServiceAmountHandler sets the operation handler for the add service amount operation
	AddServiceAmountHandler AddServiceAmountHandler
	// DelStationHandler sets the operation handler for the del station operation
	DelStationHandler DelStationHandler
	// GetPingHandler sets the operation handler for the get ping operation
	GetPingHandler GetPingHandler
	// InfoHandler sets the operation handler for the info operation
	InfoHandler InfoHandler
	// KasseHandler sets the operation handler for the kasse operation
	KasseHandler KasseHandler
	// LoadHandler sets the operation handler for the load operation
	LoadHandler LoadHandler
	// LoadMoneyHandler sets the operation handler for the load money operation
	LoadMoneyHandler LoadMoneyHandler
	// LoadRelayHandler sets the operation handler for the load relay operation
	LoadRelayHandler LoadRelayHandler
	// OpenStationHandler sets the operation handler for the open station operation
	OpenStationHandler OpenStationHandler
	// PingHandler sets the operation handler for the ping operation
	PingHandler PingHandler
	// ProgramRelaysHandler sets the operation handler for the program relays operation
	ProgramRelaysHandler ProgramRelaysHandler
	// ProgramsHandler sets the operation handler for the programs operation
	ProgramsHandler ProgramsHandler
	// SaveHandler sets the operation handler for the save operation
	SaveHandler SaveHandler
	// SaveCollectionHandler sets the operation handler for the save collection operation
	SaveCollectionHandler SaveCollectionHandler
	// SaveIfNotExistsHandler sets the operation handler for the save if not exists operation
	SaveIfNotExistsHandler SaveIfNotExistsHandler
	// SaveMoneyHandler sets the operation handler for the save money operation
	SaveMoneyHandler SaveMoneyHandler
	// SaveRelayHandler sets the operation handler for the save relay operation
	SaveRelayHandler SaveRelayHandler
	// SetKasseHandler sets the operation handler for the set kasse operation
	SetKasseHandler SetKasseHandler
	// SetProgramNameHandler sets the operation handler for the set program name operation
	SetProgramNameHandler SetProgramNameHandler
	// SetProgramRelaysHandler sets the operation handler for the set program relays operation
	SetProgramRelaysHandler SetProgramRelaysHandler
	// SetStationHandler sets the operation handler for the set station operation
	SetStationHandler SetStationHandler
	// StationByHashHandler sets the operation handler for the station by hash operation
	StationByHashHandler StationByHashHandler
	// StationReportCurrentMoneyHandler sets the operation handler for the station report current money operation
	StationReportCurrentMoneyHandler StationReportCurrentMoneyHandler
	// StationReportDatesHandler sets the operation handler for the station report dates operation
	StationReportDatesHandler StationReportDatesHandler
	// StationsVariablesHandler sets the operation handler for the stations variables operation
	StationsVariablesHandler StationsVariablesHandler
	// StatusHandler sets the operation handler for the status operation
	StatusHandler StatusHandler
	// StatusCollectionHandler sets the operation handler for the status collection operation
	StatusCollectionHandler StatusCollectionHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *StorageAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *StorageAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *StorageAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *StorageAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *StorageAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *StorageAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *StorageAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the StorageAPI
func (o *StorageAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.PinCodeAuth == nil {
		unregistered = append(unregistered, "PinAuth")
	}

	if o.AddServiceAmountHandler == nil {
		unregistered = append(unregistered, "AddServiceAmountHandler")
	}

	if o.DelStationHandler == nil {
		unregistered = append(unregistered, "DelStationHandler")
	}

	if o.GetPingHandler == nil {
		unregistered = append(unregistered, "GetPingHandler")
	}

	if o.InfoHandler == nil {
		unregistered = append(unregistered, "InfoHandler")
	}

	if o.KasseHandler == nil {
		unregistered = append(unregistered, "KasseHandler")
	}

	if o.LoadHandler == nil {
		unregistered = append(unregistered, "LoadHandler")
	}

	if o.LoadMoneyHandler == nil {
		unregistered = append(unregistered, "LoadMoneyHandler")
	}

	if o.LoadRelayHandler == nil {
		unregistered = append(unregistered, "LoadRelayHandler")
	}

	if o.OpenStationHandler == nil {
		unregistered = append(unregistered, "OpenStationHandler")
	}

	if o.PingHandler == nil {
		unregistered = append(unregistered, "PingHandler")
	}

	if o.ProgramRelaysHandler == nil {
		unregistered = append(unregistered, "ProgramRelaysHandler")
	}

	if o.ProgramsHandler == nil {
		unregistered = append(unregistered, "ProgramsHandler")
	}

	if o.SaveHandler == nil {
		unregistered = append(unregistered, "SaveHandler")
	}

	if o.SaveCollectionHandler == nil {
		unregistered = append(unregistered, "SaveCollectionHandler")
	}

	if o.SaveIfNotExistsHandler == nil {
		unregistered = append(unregistered, "SaveIfNotExistsHandler")
	}

	if o.SaveMoneyHandler == nil {
		unregistered = append(unregistered, "SaveMoneyHandler")
	}

	if o.SaveRelayHandler == nil {
		unregistered = append(unregistered, "SaveRelayHandler")
	}

	if o.SetKasseHandler == nil {
		unregistered = append(unregistered, "SetKasseHandler")
	}

	if o.SetProgramNameHandler == nil {
		unregistered = append(unregistered, "SetProgramNameHandler")
	}

	if o.SetProgramRelaysHandler == nil {
		unregistered = append(unregistered, "SetProgramRelaysHandler")
	}

	if o.SetStationHandler == nil {
		unregistered = append(unregistered, "SetStationHandler")
	}

	if o.StationByHashHandler == nil {
		unregistered = append(unregistered, "StationByHashHandler")
	}

	if o.StationReportCurrentMoneyHandler == nil {
		unregistered = append(unregistered, "StationReportCurrentMoneyHandler")
	}

	if o.StationReportDatesHandler == nil {
		unregistered = append(unregistered, "StationReportDatesHandler")
	}

	if o.StationsVariablesHandler == nil {
		unregistered = append(unregistered, "StationsVariablesHandler")
	}

	if o.StatusHandler == nil {
		unregistered = append(unregistered, "StatusHandler")
	}

	if o.StatusCollectionHandler == nil {
		unregistered = append(unregistered, "StatusCollectionHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *StorageAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *StorageAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "pinCode":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.PinCodeAuth(token)
			})

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *StorageAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *StorageAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *StorageAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *StorageAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the storage API
func (o *StorageAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *StorageAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/add-service-amount"] = NewAddServiceAmount(o.context, o.AddServiceAmountHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/del-station"] = NewDelStation(o.context, o.DelStationHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ping"] = NewGetPing(o.context, o.GetPingHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/info"] = NewInfo(o.context, o.InfoHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/kasse"] = NewKasse(o.context, o.KasseHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/load"] = NewLoad(o.context, o.LoadHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/load-money"] = NewLoadMoney(o.context, o.LoadMoneyHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/load-relay"] = NewLoadRelay(o.context, o.LoadRelayHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/open-station"] = NewOpenStation(o.context, o.OpenStationHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/ping"] = NewPing(o.context, o.PingHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/program-relays"] = NewProgramRelays(o.context, o.ProgramRelaysHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/programs"] = NewPrograms(o.context, o.ProgramsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/save"] = NewSave(o.context, o.SaveHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/save-collection"] = NewSaveCollection(o.context, o.SaveCollectionHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/save-if-not-exists"] = NewSaveIfNotExists(o.context, o.SaveIfNotExistsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/save-money"] = NewSaveMoney(o.context, o.SaveMoneyHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/save-relay"] = NewSaveRelay(o.context, o.SaveRelayHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/set-kasse"] = NewSetKasse(o.context, o.SetKasseHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/set-program-name"] = NewSetProgramName(o.context, o.SetProgramNameHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/set-program-relays"] = NewSetProgramRelays(o.context, o.SetProgramRelaysHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/set-station"] = NewSetStation(o.context, o.SetStationHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/station-by-hash"] = NewStationByHash(o.context, o.StationByHashHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/station-report-current-money"] = NewStationReportCurrentMoney(o.context, o.StationReportCurrentMoneyHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/station-report-dates"] = NewStationReportDates(o.context, o.StationReportDatesHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/stations-variables"] = NewStationsVariables(o.context, o.StationsVariablesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status"] = NewStatus(o.context, o.StatusHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status-collection"] = NewStatusCollection(o.context, o.StatusCollectionHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *StorageAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *StorageAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *StorageAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *StorageAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
