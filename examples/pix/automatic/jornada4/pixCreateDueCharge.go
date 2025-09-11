package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const txid = "d8265d2424114747b455ca884380d211"

	body := map[string]interface{}{
		"calendario": map[string]interface{}{
			"dataDeVencimento":       "2025-12-30",
			"validadeAposVencimento": 30,
		},
		"devedor": map[string]interface{}{
			"logradouro": "Alameda Souza, Numero 80, Bairro Braz",
			"cidade":     "Recife",
			"uf":         "PE",
			"cep":        "70011750",
			"cpf":        "12345678909",
			"nome":       "Francisco da Silva",
		},
		"valor": map[string]interface{}{
			"original": "123.45",
			"multa": map[string]interface{}{
				"modalidade": 2,
				"valorPerc":  "2.00",
			},
			"juros": map[string]interface{}{
				"modalidade": 2,
				"valorPerc":  "2.00",
			},
			"desconto": map[string]interface{}{
				"modalidade": 1,
				"descontoDataFixa": []map[string]interface{}{
					{
						"data":      "2025-10-30",
						"valorPerc": "30.00",
					},
				},
			},
		},
		"chave":              "88714dc7-dc1d-42ec-a4ed-44f2eaf1cd4c",
		"solicitacaoPagador": "Cobrança dos serviços prestados.",
	}

	res, err := efi.CreateDueCharge(txid, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
