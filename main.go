package main

import (
	"fmt"
	"github.com/spf13/pflag"
	exoscale "github.com/vshn/provider-exoscale/apis/exoscale/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/printers"
	"log"
	"os"
)

var (
	kind string
)

func main() {
	pflag.StringVar(&kind, "kind", "", "K8s Kind")
	err := pflag.CommandLine.Parse(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("kind: %s\n", kind)

	em := exoscale.MySQL{
		TypeMeta:   v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{},
		Spec:       exoscale.MySQLSpec{},
		Status:     exoscale.MySQLStatus{},
	}

	yp := printers.YAMLPrinter{}

	yp.PrintObj(em, os.Stdout)

}
