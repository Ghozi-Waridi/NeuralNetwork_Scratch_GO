package main

import (
	"fmt"
	"neuralnetworks/pkg/dataset"
)

func main() {
	PATH_TEST := "/Users/ghoziwaridi/PEMOGRAMAN/go/NeuralNetworks/data/raw/test"
	PATH_TRAIN := "/Users/ghoziwaridi/PEMOGRAMAN/go/NeuralNetworks/data/raw/train"

	// Load Dataset
	train_loader, labels_train := loader.DataLoader(PATH_TRAIN, true)
	test_loader, labels_test := loader.DataLoader(PATH_TEST, true)

	fmt.Println("Train Dataset: ", train_loader)
	fmt.Println("Labels: ", labels_train)
	fmt.Println("Test Dataset: ", test_loader)
	fmt.Println("Labels: ", labels_test)

}
