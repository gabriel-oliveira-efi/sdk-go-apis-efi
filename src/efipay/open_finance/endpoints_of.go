package open_finance

type endpoints struct {
	requester interface {
		request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}, headers map[string]string) (string, error)
	}
}

func (endpoints endpoints) OfListParticipants(nome string) (string, error) {
	params := map[string]string{"nome": (nome)}
	return endpoints.requester.request("/participantes", "GET", params, nil, nil)
}

func (endpoints endpoints) OfConfigUpdate(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/config", "PUT", nil, body, nil)
}

func (endpoints endpoints) OfConfigDetail() (string, error) {
	return endpoints.requester.request("/config", "GET", nil, nil, nil)
}

func (endpoints endpoints) OfDevolutionPix(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/devolucao/pagamento/pix", "POST", nil, body, nil)
}

func (endpoints endpoints) OfListPixPayment(inicio string, fim string) (string, error) {
	params := map[string]string{
		"inicio": (inicio),
		"fim":    (fim),
	}
	return endpoints.requester.request("/pagamentos/pix", "GET", params, nil, nil)
}

func (endpoints endpoints) OfStartPixPayment(body map[string]interface{}, headers map[string]string) (string, error) {
	return endpoints.requester.request("/pagamentos/pix", "POST", nil, body, headers)
}

func (endpoints endpoints) OfStartRecurrencyPixPayment(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix", "POST", nil, body, nil)
}

func (endpoints endpoints) OfListRecurrencyPixPayment(inicio string, fim string) (string, error) {
	params := map[string]string{
		"inicio": (inicio),
		"fim":    (fim),
	}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix", "GET", params, nil, nil)
}

func (endpoints endpoints) OfCancelRecurrencyPix(identificadorPagamento string) (string, error) {
	params := map[string]string{"identificadorPagamento": (identificadorPagamento)}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix/:identificadorPagamento/cancelar", "PATCH", params, nil, nil)
}

func (endpoints endpoints) OfDevolutionRecurrencyPix(identificadorPagamento string, body map[string]interface{}) (string, error) {
	params := map[string]string{
		"identificadorPagamento": (identificadorPagamento),
	}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix/:identificadorPagamento/devolver", "POST", nil, body, params)
}

func (endpoints endpoints) OfReplaceRecurrencyPixParcel(identificadorPagamento string, endToEndId string, body map[string]interface{}) (string, error) {
	params := map[string]string{
		"identificadorPagamento": (identificadorPagamento),
		"endToEndId":             (endToEndId),
	}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix/:identificadorPagamento/substituir/:endToEndId", "PATCH", nil, body, params)
}

// ofCancelSchedulePix
func (endpoints endpoints) OfCancelSchedulePix(identificadorPagamento string) (string, error) {
	params := map[string]string{"identificadorPagamento": (identificadorPagamento)}
	return endpoints.requester.request("/v1/pagamentos-agendados/pix/:identificadorPagamento/cancelar", "PATCH", params, nil, nil)
}

// ofListSchedulePixPayment
func (endpoints endpoints) OfListSchedulePixPayment(inicio string, fim string) (string, error) {
	params := map[string]string{
		"inicio": (inicio),
		"fim":    (fim),
	}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix", "GET", params, nil, nil)
}

// ofStartSchedulePixPayment
func (endpoints endpoints) OfStartSchedulePixPayment(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix", "POST", nil, body, nil)
}

// ofDevolutionSchedulePix
func (endpoints endpoints) OfDevolutionSchedulePix(identificadorPagamento string, body map[string]interface{}) (string, error) {
	params := map[string]string{
		"identificadorPagamento": (identificadorPagamento),
	}
	return endpoints.requester.request("/v1/pagamentos-recorrentes/pix/:identificadorPagamento/devolver", "POST", nil, body, params)
}
