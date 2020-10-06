package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kingpin"
)

var (
	app = kingpin.New("font2css", "embeds a font-face as css")

	cmdFile          = app.Command("file", "make css from font file")
	flagFamily       = cmdFile.Flag("family", "font-family").Short('f').Required().String()
	flagStyle        = cmdFile.Flag("style", "font-style [normal, italic]").Short('s').Required().Enum("normal", "italic")
	flagWeight       = cmdFile.Flag("weight", "font-weight (numeric)").Short('w').Required().Enum("100", "200", "300", "400", "500", "600", "700", "800", "900")
	flagDisplay      = cmdFile.Flag("display", "font-display [auto, block, swap, fallback, optional]").Short('d').Default("swap").Enum("auto", "block", "swap", "fallback", "optional")
	flagLocals       = cmdFile.Flag("locals", "local install names, can specify multiple times. ex: -l \"Nunito Sans Regular\" -l NunitoSans-Regular (optional)").Short('l').Strings()
	flagUnicodeRange = cmdFile.Flag("unicode", "unicode-range. ex -u \"U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD\" (optional)").Short('u').String()
	flagInput        = cmdFile.Flag("input", "input filename").Short('i').Required().String()
	flagOutput       = cmdFile.Flag("output", "output file name (optional)").Short('o').String()

	cmdURL          = app.Command("url", "make css from google fonts url")
	flagURL         = cmdURL.Flag("url", "Google Fonts URL ex: https://fonts.googleapis.com/css2?family=Nunito+Sans:ital,wght@0,200;0,300;0,400;0,600;0,700;0,800;0,900;1,200;1,300;1,400;1,600;1,700;1,800;1,900&display=swap").Short('u').Required().String()
	flagURLCharsets = cmdURL.Flag("charsets", "only do certain charsets, can specify multiple times. ex: -c latin -c latin-ext (optional)").Short('c').Strings()
	flagURLStyles   = cmdURL.Flag("styles", "only do certain styles, can specify multiple times. ex: -s normal -s italic (optional)").Short('s').Strings()
	flagURLWeights  = cmdURL.Flag("weights", "weights to include, can specify multiple times. ex: -w 200 -w 400 -w 600 -w 900 (optional)").Short('w').Strings()
	flagURLOutput   = cmdURL.Flag("output", "output file name (optional)").Short('o').String()
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"--help-long"}
	}
	cmd, err := app.Parse(args)
	if err != nil {
		log.Fatal(err)
	}

	switch cmd {
	case cmdFile.FullCommand():
		runFile()
	case cmdURL.FullCommand():
		runURL()
	}
}

type fontFace struct {
	charset string
	style   string
	weight  int
	raw     string
}

func runURL() {
	fontsCSS, err := wget(*flagURL)
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(fontsCSS, []byte{'\n'})

	// TODO

	_ = lines
	_ = flagURLCharsets
	_ = flagURLStyles
	_ = flagURLWeights
}

func wget(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func runFile() {
	format := strings.TrimPrefix(filepath.Ext(*flagInput), ".")

	switch format {
	case "ttf":
	case "otf":
	case "woff":
	case "woff2":
	default:
		log.Fatalf("unhandled font format '%s'", format)
	}

	// it's actually this easy. SUPPOSEDLY. https://www.iana.org/assignments/media-types/media-types.xhtml#font
	mimeType := fmt.Sprintf("font/%s", format)

	b, err := ioutil.ReadFile(*flagInput)
	if err != nil {
		log.Fatal(err)
	}

	b64 := base64.StdEncoding.EncodeToString(b)

	var locals string
	for _, val := range *flagLocals {
		locals += fmt.Sprintf("local('%s'), ", val)
	}

	var fontFace strings.Builder
	fontFace.WriteString("@font-face {\n")
	fontFace.WriteString(fmt.Sprintf("  font-family: '%s';\n", *flagFamily))
	fontFace.WriteString(fmt.Sprintf("  font-style: %s;\n", *flagStyle))
	fontFace.WriteString(fmt.Sprintf("  font-weight: %s;\n", *flagWeight))
	fontFace.WriteString(fmt.Sprintf("  font-display: %s;\n", *flagDisplay))
	fontFace.WriteString(fmt.Sprintf("  src: %surl(data:%s;charset=utf-8;base64,%s) format('%s');\n", locals, mimeType, b64, format))
	if *flagUnicodeRange != "" {
		fontFace.WriteString(fmt.Sprintf("  unicode-range: %s;\n", *flagUnicodeRange))
	}
	fontFace.WriteString("}\n")

	if *flagOutput != "" {
		if err = ioutil.WriteFile(*flagOutput, []byte(fontFace.String()), 0644); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("%s", fontFace)
	}
}
