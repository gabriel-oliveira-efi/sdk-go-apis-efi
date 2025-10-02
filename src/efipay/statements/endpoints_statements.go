package statements

type endpoints struct {
	requester interface {
		request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}) (string, error)
	}
}

// listStatementFiles	GET	/v1/extrato-cnab/arquivos
func (endpoints endpoints) ListStatementFiles(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/arquivos", "GEt", nil, body)
}

// getStatementFile	GET	/v1/extrato-cnab/download/:nome_arquivo
func (endpoints endpoints) GetStatementFile(params map[string]string) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/download/:nome_arquivo", "GEt", params, nil)
}

// listStatementRecurrences	GET	/v1/extrato-cnab/agendamentos
func (endpoints endpoints) ListStatementRecurrences(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/agendamentos", "GEt", nil, body)
}

// createStatementRecurrency	POST	/v1/extrato-cnab/agendar
func (endpoints endpoints) CreateStatementRecurrency(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/agendar", "POST", nil, body)
}

// updateStatementRecurrency	PATCH	/v1/extrato-cnab/agendar/:identificador
func (endpoints endpoints) UpdateStatementRecurrency(params map[string]string, body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/agendar/:identificador", "PATCH", params, body)
}

// createSftpKey	POST	/v1/extrato-cnab/gerar-chaves
func (endpoints endpoints) CreateSftpKey(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v1/extrato-cnab/gerar-chaves", "POST", nil, body)
}
