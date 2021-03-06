package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"

	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	c := new(bluemix.Config)

	var cluster string
	flag.StringVar(&cluster, "cluster", "", "Clusetr Name")

	var InstanceID string
	flag.StringVar(&InstanceID, "InstanceID", "", " monitoring InstanceID")

	var ingestionKey string
	flag.StringVar(&ingestionKey, "ingestionKey", "", "ingestion Key")

	var endPoint bool
	flag.BoolVar(&endPoint, "endPoint", false, "private EndPoint (true/false)")

	flag.Parse()

	trace.Logger = trace.NewLogger("true")
	if cluster == "" || InstanceID == "" {
		flag.Usage()
		os.Exit(1)
	}

	var loggingInfo = v2.LoggingCreateRequest{
		Cluster:         cluster,
		LoggingInstance: InstanceID,
		IngestionKey:    ingestionKey,
		PrivateEndpoint: endPoint,
	}

	sess, err := session.New(c)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	target := v2.LoggingTargetHeader{}

	loggingClient, err := v2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	loggingAPI := loggingClient.Logging()

	out, err := loggingAPI.CreateLoggingConfig(loggingInfo, target)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("out=", out)
}
