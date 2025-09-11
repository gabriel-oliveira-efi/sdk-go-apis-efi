package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const txid = "00000000000000000000000001"

	body := map[string]interface{}{
		"idRec":         "RR000000000000000000000000001",
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

	res, err := efi.PixCreateAutomaticChargeTxid(txid, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
