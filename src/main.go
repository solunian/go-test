// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {
	fmt.Println("Hello, 世界")
	err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "qr.png")

	if err != nil {
		fmt.Println("Error generating QR code:", err)
	} else {
		fmt.Println("QR code generated successfully!")
	}
}
