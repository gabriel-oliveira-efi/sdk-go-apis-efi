package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	body := map[string]interface{}{
		"idRec":         "RN09089356202509040648092afd6",
		"infoAdicional": "Streamming.",
		"calendario": map[string]interface{}{
			"dataDeVencimento": "2026-12-31",
		},
		"valor": map[string]interface{}{
			"original": "106.07",
		},
		"ajusteDiaUtil": true,
		"devedor": map[string]interface{}{
			"cep":        "12345678",
			"cidade":     "City",
			"email":      "client.mail@server.com",
			"logradouro": "Street name 123",
			"uf":         "MG",
		},
		"recebedor": map[string]interface{}{
			"agencia":   "9708",
			"conta":     "12682",
			"tipoConta": "CORRENTE",
		},
	}

	res, err := efi.PixCreateAutomaticCharge(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
