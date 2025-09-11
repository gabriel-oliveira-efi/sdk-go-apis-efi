package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const idRec = "RN09089356202509033f1a3e38abd"

	body := map[string]interface{}{
		"loc": 1,
		"vinculo": map[string]interface{}{
			"devedor": map[string]interface{}{
				"nome": "Gorbadoc Oldbuck",
			},
		},
		"calendario": map[string]interface{}{
			"dataInicial": "2026-02-25",
		},
		"ativacao": map[string]interface{}{
			"dadosJornada": map[string]interface{}{
				"txid": "bd22a44ec40c4262a9f4e24150b5cb3f",
			},
		},
	}

	res, err := efi.PixUpdateRecurrenceAutomatic(idRec, body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
