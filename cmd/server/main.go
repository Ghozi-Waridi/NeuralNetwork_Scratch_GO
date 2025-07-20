package main

import (
	"fmt"
	"neuralnetworks/pkg/dataset"
	"neuralnetworks/pkg/model"
	"neuralnetworks/pkg/encoder"
)

func main() {
	PATH_TEST := "/Users/ghoziwaridi/PEMOGRAMAN/go/NeuralNetworks/data/raw/test"
	PATH_TRAIN := "/Users/ghoziwaridi/PEMOGRAMAN/go/NeuralNetworks/data/raw/train"

	// Load Dataset
	train_loader, labels_train := loader.DataLoader(PATH_TRAIN, true)
	test_loader, labels_test := loader.DataLoader(PATH_TEST, true)

	fmt.Println("Train Dataset: ", len(train_loader))
	fmt.Println("Labels: ", labels_train)
	fmt.Println("Test Dataset: ", len(test_loader))
	fmt.Println("Labels: ", labels_test)

	model := model.InitHyperParameters(0.01, 764, 128, 64, 2)
	fmt.Println("Model Hyperparameters: ", model)

	encoder := encoder.NewEncoder(labels_train)
	encoder.Fit(labels_train)
	oneHotTargets := encoder.Transform(labels_train)

	hasil := model.Train(train_loader, oneHotTargets, 10)
	fmt.Println("Hasil LOSS Training: ", hasil)
}
