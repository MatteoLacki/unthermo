//The printspectrum tool prints out the spectrum (mz and intensity values) of a
//Thermo RAW File
//
//  Every line of the output is a peak registered by the mass spectrometer
//  characterized by an m/z value in Da and an intensity in the mass spectrometer's unit of abundance
package main

import (
	"bitbucket.org/proteinspector/ms/unthermo"
	"flag"
	"fmt"
	"log"
)

func main() {
	var filename string
	var scans int
	//Parse arguments
	flag.StringVar(&filename, "raw", "small.RAW", "name of the RAW file")
	flag.Parse()

	//open RAW file
	file, err := unthermo.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scans = file.NScans()
	defer file.Close()

	//Print the number of scans.
	fmt.Println(scans)
}
