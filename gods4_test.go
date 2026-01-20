package gods4

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestListen(t *testing.T) {
	controllers := Find()
	if len(controllers) == 0 {
		t.Fatal("No connected DS4 controllers found")
	}

	// Select first controller from the list
	controller := controllers[0]

	// Connect to the controller
	err := controller.Connect()
	if err != nil {
		t.Fatal(err)
	}

	controller.On(EventCirclePress, func(_ any) error {
		_, err := fmt.Println(EventCirclePress)
		return err
	})

	go controller.Listen()
	time.Sleep(time.Second)
	err = controller.Disconnect()
	if err != nil {
		t.Fatal(err)
	}
}

func TestListenContext(t *testing.T) {
	controllers := Find()
	if len(controllers) == 0 {
		t.Fatal("No connected DS4 controllers found")
	}

	// Select first controller from the list
	controller := controllers[0]

	// Connect to the controller
	err := controller.Connect()
	if err != nil {
		t.Fatal(err)
	}

	controller.On(EventCirclePress, func(_ any) error {
		_, err := fmt.Println(EventCirclePress)
		return err
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = controller.ListenContext(ctx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
	err = controller.Disconnect()
	if err != nil {
		t.Fatal(err)
	}
}
