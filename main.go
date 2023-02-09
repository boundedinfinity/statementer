package main

import (
	"fmt"
	"os"

	"github.com/boundedinfinity/docsorter/runtime"
	"github.com/boundedinfinity/docsorter/util"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	rt := runtime.New(logger)

	if err := rt.LoadUserConfig(os.Args[1]); err != nil {
		handleError(err)
		return
	}

	ocrs, err := rt.LoadFiles()

	if err != nil {
		handleError(err)
		return
	}

	for _, ocr := range ocrs {
		logger.Infof(util.PrintSep())
		logger.Infof(util.PrintLabeled("Source", ocr.Stage1.Source))

		if err := rt.OcrSingle(&ocr.Stage1); err != nil {
			handleError(err)
			return
		}

		if err := rt.Process(&ocr); err != nil {
			handleError(err)
			return
		}

		if err := rt.Rename(ocr, &ocr.Stage2, ocr.Stage1); err != nil {
			handleError(err)
			return
		}

		if err := rt.DumpCvs(&ocr); err != nil {
			handleError(err)
			return
		}

		if err := rt.DumpYaml(&ocr); err != nil {
			handleError(err)
			return
		}

		if err := rt.Output(&ocr.Dest, &ocr.Stage2); err != nil {
			handleError(err)
			return
		}
	}
}

func handleError(err error) {
	fmt.Printf("error %v\n", err)
}
