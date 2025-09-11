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
		"vinculo": map[string]interface{}{
			"contrato": "63101869",
			"devedor": map[string]interface{}{
				"cpf":  "12345678910",
				"nome": "Gorbadoc Oldbuck",
			},
			"objeto": "Streamming1",
		},
		"calendario": map[string]interface{}{
			"dataInicial":   "2026-01-01",
			"dataFinal":     "2027-12-31",
			"periodicidade": "MENSAL", /*SEMANAL/MENSAL/TRIMESTRAL/SEMESTRAL/ANUAL*/
		},
		"valor": map[string]interface{}{
			"valorRec": "37.00",
			// "valorMinimoRecebedor": "30.00"
		},
		"politicaRetentativa": "NAO_PERMITE", // PERMITE_3R_7D
		"loc":                 1,
		"ativacao": map[string]interface{}{
			"dadosJornada": map[string]interface{}{
				"txid": "6d9020e1c6e6447f895d4b7f4f48c110",
			},
		},
	}

	res, err := efi.PixCreateRecurrenceAutomatic(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
