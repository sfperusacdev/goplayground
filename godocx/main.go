package main

import (
	"log"

	"github.com/lukasjarosch/go-docx"
)

func main() {
	replaceMap := docx.PlaceholderMap{"nombre_empresa": "AGRÍCOLA ALAYA S.A.C", "dato_prueba": "este es un texto de prueba"}

	// read and parse the template docx
	doc, err := docx.Open("/home/user002/Desktop/template_contratos/my_word_template.docx")
	if err != nil {
		log.Fatalln(err)
	}

	// replace the keys with values from replaceMap
	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		log.Fatalln(err)
	}

	// write out a new file
	err = doc.WriteToFile("/home/user002/Desktop/docs_generateds/generated_doc_go.docx")

	if err != nil {
		log.Fatalln(err)
	}
}
