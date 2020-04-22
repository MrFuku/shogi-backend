package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Sample struct {
	UserID string `dynamo:"UserID,hash"`
	Name   string `dynamo:"Name"`
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String(os.Getenv("DYNAMO_ENDPOINT")),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}

	db := dynamo.New(sess)

	// err = db.CreateTable("Samples", Sample{}).Run()
	// if err != nil {
	// 	panic(err)
	// }

	table := db.Table("Samples")
	err = table.Put(&Sample{UserID: "1", Name: "Test1"}).Run()
	if err != nil {
		panic(err)
	}

	var sample Sample
	err = table.Get("UserID", "1").One(&sample)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", sample)

	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
