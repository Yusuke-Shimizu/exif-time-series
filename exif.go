package main

import (
	// "fmt"
	"github.com/rwcarlsen/goexif/exif"
	//"github.com/rwcarlsen/goexif/mknote"
	"log"
	"os"
	"reflect"
)

func get_focal_length(x *exif.Exif) float64 {
	focal_length, _ := x.Get(exif.FocalLength)
	log.Println(focal_length)
	log.Println(*focal_length)
	// list_struct(focal_length)
	// log.Println("focal_length.ratVals", (*focal_length).ratVals)
	log.Println(reflect.TypeOf(focal_length))
	log.Printf("%+v\n", *focal_length)
	// exif.Rat2(focal_length)

	// return focal_length
	return 1.1
}

func print_info(x *exif.Exif) {
	log.Println("exif ", x)

	log.Println("--- Cam info ---")
	make, _ := x.Get(exif.Make)
	software, _ := x.Get(exif.Software)
	model, _ := x.Get(exif.Model)
	lens_model, _ := x.Get(exif.LensModel)
	log.Println(make)
	log.Println(software)
	log.Println(model)
	log.Println(lens_model)

	log.Println("--- Date ---")
	tm, _ := x.DateTime()
	log.Println(tm)

	log.Println("--- Sensor info ---")
	iso_speed_ratings, _ := x.Get(exif.ISOSpeedRatings)
	f_number, _ := x.Get(exif.FNumber)
	exposure_time, _ := x.Get(exif.ExposureTime)
	exposure_bias_value, _ := x.Get(exif.ExposureBiasValue)
	focal_length, _ := x.Get(exif.FocalLength)
	focal_length_in_35mm_film, _ := x.Get(exif.FocalLengthIn35mmFilm)
	log.Println(iso_speed_ratings)
	log.Println(f_number)
	log.Println(exposure_time)
	log.Println(exposure_bias_value)
	log.Println(focal_length)
	log.Println(focal_length_in_35mm_film)

	// camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
	// log.Println(camModel.StringVal())
	// log.Println(camModel)

	// focal, _ := x.Get(exif.FocalLength)
	// numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
	// fmt.Printf("%v/%v\n", numer, denom)
	// log.Println("focal ", focal)

	// // Two convenience functions exist for date/time taken and GPS coords:
	// tm, _ := x.DateTime()
	// log.Println("Taken: ", tm)

	// lat, long, _ := x.LatLong()
	// log.Println("lat, long: ", lat, ", ", long)
}

// main
func main() {
	// fname := "DSC01189.ARW"
	fname := "./exif-samples/jpg/Canon_40D.jpg"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	// exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	print_info(x)
	get_focal_length(x)
	// log.Println(get_focal_length(x))
}
