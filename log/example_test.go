package log_test

import "go-common-utils/log"

func ExampleLogger()  {
	log.Config(log.Discard,log.Stdout,log.Stdout|log.EnableFile,log.Stderr|log.EnableFile,"error.log")
	log.Trace.Println("I have something standard to say")
	log.Info.Println("Special Information")
	log.Warning.Println("There is something you need to know about")
	log.Error.Println("Something has failed")
}