package main

import (
	"verle_go/pkg/sheets"
	"verle_go/pkg/weeek"
)

func main() {
	weeek.InitClient()
	sheets.InitClient()
}
