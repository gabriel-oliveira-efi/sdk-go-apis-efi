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
			"contrato": "63100862",
			"devedor": map[string]interface{}{
				"cpf":  "11122233344",
				"nome": "Gorbadoc Oldbuck",
			},
			"objeto": "Streamming",
		},
		"calendario": map[string]interface{}{
			"dataInicial":   "2026-01-01",
			"dataFinal":     "2027-12-31",
			"periodicidade": "MENSAL", /*SEMANAL/MENSAL/TRIMESTRAL/SEMESTRAL/ANUAL*/
		},
		"valor": map[string]interface{}{
			"valorRec": "35.00",
			// "valorMinimoRecebedor": "30.00"
		},
		"politicaRetentativa": "NAO_PERMITE", // PERMITE_3R_7D
		"loc":                 1,
		"ativacao": map[string]interface{}{
			"dadosJornada": map[string]interface{}{
				"txid": "33beb661beda44a8928fef47dbeb2dc5",
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
