package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/golage/errors"
	pkg "github.com/pkg/errors"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	switch err, code := errors.Parse(something1()); code {
	case errors.CodeNil:
	case errors.CodeNotFound:
		log.Fatalf("can not something1: %v", err)
	default:
		log.Fatalf("can not something1: %v\n%v", err, err.StackTrace())
	}

	if err, _ := errors.Parse(something2()); err != nil {
		log.Fatalf("can not something2: %v", err)
	}

	log.Print("success!!")
}

func something1() error {
	switch n := rand.Intn(3); n {
	case 0:
		return errors.New(errors.CodeNotFound, "something1 not found with %v", n)
	case 1:
		return errors.Wrap(pkg.New("can not connect to db"), errors.CodeInternal, "something1 with %v", n)
	default:
		return nil
	}
}

func something2() error {
	switch n := rand.Intn(2); n {
	case 0:
		return errors.New(errors.CodeAlreadyExists, "something2 already exists")
	default:
		return nil
	}
}
