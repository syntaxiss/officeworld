package ventas

import (
	"context"
	"net/http"

	"github.com/syntaxiss/officeworld-sdk-go/pkg/configuracion"
	"github.com/syntaxiss/officeworld-sdk-go/pkg/internal/httpclient"
)

const (
	urlBase                                  = "https://api.worldoffice.cloud/api/v1/"
	urlActualizarDocumentoDeVenta            = urlBase + "documentos/editarDocumentoVenta"
	urlListarDocumentosDeVenta               = urlBase + "documentos/listarDocumentosVenta"
	urlConsultarProductosDelDocumentoDeVenta = urlBase + "documentos/getRenglonesByDocumentoEncabezado/" + "{id}"
	urlFacturaElectronicaPorId               = urlBase + "documentos/facturaElectronica/" + "{id}"
	urlCrearDocumentoDeVenta                 = urlBase + "documentos/crearDocumentoVenta"
	urlContabilizarDocumentoDeVenta          = urlBase + "documentos/contabilizarDocumento/" + "{id}"
	urlAnularDocumentoDeVenta                = urlBase + "documentos/anularDocumento/" + "{id}"
	urlListarMotivosDeGeneracion             = urlBase + "documentoMotivosGeneracion/listarGeneracion"
	urlCrearPdfDocumentoDeVenta              = urlBase + "documentos/visualizarDocumento/" + "{id}"
	urlCrearPdfTicketDocumentoDeVenta        = urlBase + "documentos/ticketDocumento/" + "{id}"
	urlObtenerDocumentoDeVenta               = urlBase + "documentos/getDocumentoId/" + "{id}"
	urlEnviarDocumentoDeVentaPorCorreo       = urlBase + "documentos/enviaDocumentoMail/" + "{id}" + "/{correo}" + "/{nombre}"
	urlConsultarClasificaciones              = urlBase + "documentos/clasificaciones/" + "{documentotipo}"
	urlEliminarDocumentosDeVenta             = urlBase + "documentos/eliminarLista/" + "{ids}"
	urlEliminarDocumentoDeVenta              = urlBase + "documentos/eliminar/" + "{id}"
)

type Client interface {
	ActualizarDocumentoDeVenta(ctx context.Context, request DocumentVentaRequest) (*VentaResponse, error) // Este pide ID porque lo actualiza

	ListarDocumentosDeVenta(ctx context.Context, request SearchRequest) (*SearchResponseLista, error)

	ConsultarProductosDelDocumentoDeVenta(ctx context.Context, id string) (*SearchResponseDetalles, error)

	FacturaElectronicaPorId(ctx context.Context, id string) (*BillElectronicResponse, error)

	CrearDocumentoDeVenta(ctx context.Context, request DocumentVentaRequest) (*VentaResponse, error) // Este no pide ID porque lo crea

	ContabilizarDocumentoDeVenta(ctx context.Context, id string) (*SearchResponse, error)

	AnularDocumentoDeVenta(ctx context.Context, id string) (*SearchResponse, error)

	ListarMotivosDeGeneracion(ctx context.Context, request SearchRequest) (*SearchResponseConsulta, error)

	CrearPdfDocumentoDeVenta(ctx context.Context, id string) ([]byte, error)

	ObtenerDocumentoDeVenta(ctx context.Context, id string) (*VentaResponse, error)

	CrearPdfTicketDocumentoDeVenta(ctx context.Context, id string) ([]byte, error)

	EnviarDocumentoDeVentaPorCorreo(ctx context.Context, id string, correo string, nombre string) (*SearchResponse, error)

	ConsultarClasificaciones(ctx context.Context, documentotipo string) (*SearchResponse, error)

	EliminarDocumentosDeVenta(ctx context.Context, ids string) (*SearchResponse, error)

	EliminarDocumentoDeVenta(ctx context.Context, id string) (*SearchResponse, error)
}

type client struct {
	configuracion *configuracion.Configuracion
}

func NewClient(c *configuracion.Configuracion) Client {
	return &client{
		configuracion: c,
	}
}

func (c *client) ActualizarDocumentoDeVenta(ctx context.Context, request DocumentVentaRequest) (*VentaResponse, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPut,
		URL:    urlActualizarDocumentoDeVenta,
	}
	resource, err := httpclient.DoRequest[*VentaResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ListarDocumentosDeVenta(ctx context.Context, request SearchRequest) (*SearchResponseLista, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlListarDocumentosDeVenta,
	}
	resource, err := httpclient.DoRequest[*SearchResponseLista](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ConsultarProductosDelDocumentoDeVenta(ctx context.Context, id string) (*SearchResponseDetalles, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodPost,
		URL:        urlConsultarProductosDelDocumentoDeVenta,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*SearchResponseDetalles](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) FacturaElectronicaPorId(ctx context.Context, id string) (*BillElectronicResponse, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodPost,
		URL:        urlFacturaElectronicaPorId,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*BillElectronicResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) CrearDocumentoDeVenta(ctx context.Context, request DocumentVentaRequest) (*VentaResponse, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlCrearDocumentoDeVenta,
	}
	resource, err := httpclient.DoRequest[*VentaResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ContabilizarDocumentoDeVenta(ctx context.Context, id string) (*SearchResponse, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodPost,
		URL:        urlContabilizarDocumentoDeVenta,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) AnularDocumentoDeVenta(ctx context.Context, id string) (*SearchResponse, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodPost,
		URL:        urlAnularDocumentoDeVenta,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ListarMotivosDeGeneracion(ctx context.Context, request SearchRequest) (*SearchResponseConsulta, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    urlListarMotivosDeGeneracion,
	}
	resource, err := httpclient.DoRequest[*SearchResponseConsulta](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) CrearPdfDocumentoDeVenta(ctx context.Context, id string) ([]byte, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlCrearPdfDocumentoDeVenta,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[[]byte](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) CrearPdfTicketDocumentoDeVenta(ctx context.Context, id string) ([]byte, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlCrearPdfTicketDocumentoDeVenta,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[[]byte](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ObtenerDocumentoDeVenta(ctx context.Context, id string) (*VentaResponse, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlObtenerDocumentoDeVenta,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*VentaResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) EnviarDocumentoDeVentaPorCorreo(ctx context.Context, id string, correo string, nombre string) (*SearchResponse, error) {
	pathParams := map[string]string{
		"id":     id,
		"correo": correo,
		"nombre": nombre,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlEnviarDocumentoDeVentaPorCorreo,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ConsultarClasificaciones(ctx context.Context, documentotipo string) (*SearchResponse, error) {
	pathParams := map[string]string{
		"documentotipo": documentotipo,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlConsultarClasificaciones,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) EliminarDocumentosDeVenta(ctx context.Context, ids string) (*SearchResponse, error) {
	pathParams := map[string]string{
		"ids": ids,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodDelete,
		URL:        urlEliminarDocumentosDeVenta,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) EliminarDocumentoDeVenta(ctx context.Context, id string) (*SearchResponse, error) {
	pathParams := map[string]string{
		"id": id,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodDelete,
		URL:        urlEliminarDocumentoDeVenta,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.configuracion, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
