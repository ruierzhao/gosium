package gosium

import (
	_ "embed"
	"encoding/base64"
	"log"

	"fyne.io/fyne/v2/app"
)

/*
	 ***************************_***********************
		  __ _    ___    ___  (_)  _   _   _ __ ___
		 / _` |  / _ \  / __| | | | | | | | '_ ` _ \
		| (_| | | (_) | \__ \ | | | |_| | | | | | | |
		 \__, |  \___/  |___/ |_|  \__,_| |_| |_| |_|
		 |___/
	 ***************************************************
*/
func init() {
	var icon2 string = "KioqKioqKioqKioqKioqKioqKioqKioqKioqXyoqKioqKioqKioqKioqKioqKioqKgoJICBfXyBfICAgIF9fXyAgICBfX18gIChfKSAgXyAgIF8gICBfIF9fIF9fXwoJIC8gX2AgfCAgLyBfIFwgIC8gX198IHwgfCB8IHwgfCB8IHwgJ18gYCBfIFwKCXwgKF98IHwgfCAoXykgfCBcX18gXCB8IHwgfCB8X3wgfCB8IHwgfCB8IHwgfAoJIFxfXywgfCAgXF9fXy8gIHxfX18vIHxffCAgXF9fLF98IHxffCB8X3wgfF98CgkgfF9fXy8KKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKioqKg=="
	icon2byte, _ := base64.StdEncoding.DecodeString(icon2)
	println(string(icon2byte))
}

func Viewer() {
	app.New()
	log.Println("ruier")
}
