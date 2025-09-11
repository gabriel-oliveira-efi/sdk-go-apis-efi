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
		"calendario": map[string]interface{}{
			"expiracao": 3600, // Charge lifetime, specified in seconds from creation date
		},
		"devedor": map[string]interface{}{
			"cpf":  "12345678910",
			"nome": "Gorbadoc Oldbuck",
		},
		"valor": map[string]interface{}{
			"original": "1.80",
		},
		"chave":              "", // Pix key registered in the authenticated Efí account
		"solicitacaoPagador": "Enter the order number or identifier.",
	}

	res, err := efi.PixCreateImmediateCharge(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
