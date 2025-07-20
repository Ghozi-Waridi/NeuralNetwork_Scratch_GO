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

	fmt.Printf("Tipe Train Loader: %T\n", train_loader)
	fmt.Printf("Tipe Train Labels: %T\n\n", labels_train)
	
	// --- Menampilkan Info Dataset Testing ---
	fmt.Printf("Tipe Test Loader: %T\n", test_loader)
	fmt.Printf("Tipe Test Labels: %T\n", labels_test)

	fmt.Println(len(train_loader), len(labels_train[0]))

	model := model.InitHyperParameters(0.01, 14400, 128, 64, 609)
	// fmt.Println("Model Hyperparameters: ", model)

	encoder := encoder.NewEncoder(labels_train)
	encoder.Fit(labels_train)
	oneHotTargets := encoder.Transform(labels_train)
	
	fmt.Println("Masuk Bagian Training")
	hasil := model.Train(train_loader, oneHotTargets, 5)
	fmt.Println("Hasil LOSS Training: ", hasil)
}
