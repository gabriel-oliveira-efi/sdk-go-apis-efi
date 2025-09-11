package main

import (
	"fmt"

	"github.com/efipay/sdk-go-apis-efi/efipay/pix"
	"github.com/efipay/sdk-go-apis-efi/examples/configs"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const inicio = "2025-01-01T03:01:35Z"
	const fim = "2025-12-31T22:01:35Z"

	res, err := efi.PixListLocationRecurrenceAutomatic(inicio, fim)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
