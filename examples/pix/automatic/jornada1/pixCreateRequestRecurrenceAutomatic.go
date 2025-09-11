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
		"idRec": "RN0908935620250902c5bf68c68d4",
		"calendario": map[string]interface{}{

			"dataExpiracaoSolicitacao": "2025-10-02T12:17:11.926Z",
		},

		"destinatario": map[string]interface{}{
			"agencia":          "0001",
			"conta":            "0000000",
			"cpf":              "12345678910",
			"ispbParticipante": "09089356",
		},
	}

	res, err := efi.PixCreateRequestRecurrenceAutomatic(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
