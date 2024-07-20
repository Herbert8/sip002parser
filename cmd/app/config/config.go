package config

import (
	"flag"
	"fmt"
	"os"
	"path"
)

type ApplicationConfig struct {
	inputType  string
	outputType string
	DataSource string
	Port       uint
	TFO        bool
}

const (
	kSIP002            = "sip002"
	kShadowsocks       = "ss"
	kSurge             = "surge"
	kGost              = "gost"
	kGostLoadBalancing = "gost-lb"
)

var AppConfig ApplicationConfig

func showHelp() {

	appName := path.Base(os.Args[0])

	const sTemplate = `
Usage:
  %[1]s -it <ss | sip002> -ot <gost | surge> [-p <port>] [-tfo] data_source
	  Generate configuration from data.

  %[1]s
	  Show usage information.

Options:
  -it <ss | sip002>                         Specify input type.
  -ot <gost | gost-lb | surge               Specify output type.
  -p <port>                                 Specify local proxy port.
  -tfo                                      Specifies to use tfo mode.
`
	fmt.Printf(sTemplate, appName)
}

func ParseConfig() {

	if len(os.Args) <= 1 {
		showHelp()
		os.Exit(1)
	}

	flag.StringVar(&AppConfig.inputType, "it", kSIP002, "Specify input type.")
	flag.StringVar(&AppConfig.outputType, "ot", kGost, "Specify output type.")
	flag.UintVar(&AppConfig.Port, "p", 5678, "Specify local proxy port.")
	flag.BoolVar(&AppConfig.TFO, "tfo", false, "Specifies to use tfo mode.")

	flag.Parse()

	AppConfig.DataSource = flag.Arg(0)

	//AppConfig.inputType = kSIP002
	//AppConfig.inputType = kShadowsocks
	//AppConfig.outputType = kGost
	//AppConfig.outputType = kSurge
	//
	//AppConfig.DataSource = "https://api.nxtlnodes.com/Subscription/ShadowRocketImportService?t=sip002&sid=19276&token=5ccRH3uhCwJ"
	//AppConfig.DataSource = "https://api.nxtlnodes.com/Subscription/ShadowSocks?sid=19276&token=5ccRH3uhCwJ"
}

func (receiver *ApplicationConfig) InputTypeIsSIP002() bool {
	return receiver.inputType == kSIP002
}

func (receiver *ApplicationConfig) InputTypeIsShadowsocks() bool {
	return receiver.inputType == kShadowsocks
}

func (receiver *ApplicationConfig) OutputTypeIsGost() bool {
	return receiver.outputType == kGost
}

func (receiver *ApplicationConfig) OutputTypeIsGostLoadBalancing() bool {
	return receiver.outputType == kGostLoadBalancing
}

func (receiver *ApplicationConfig) OutputTypeIsSurge() bool {
	return receiver.outputType == kSurge
}
