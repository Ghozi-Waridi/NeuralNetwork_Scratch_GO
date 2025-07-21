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

	model := model.InitHyperParameters(0.001, 14400, 128, 64, 609)
	// fmt.Println("Model Hyperparameters: ", model)

	encoder_train := encoder.NewEncoder(labels_train)
	encoder_train.Fit(labels_train)
	oneHotTargets := encoder_train.Transform(labels_train)
	
	fmt.Println("Masuk Bagian Training")
	hasil := model.Train(train_loader, oneHotTargets, 2)
	fmt.Println("Hasil LOSS Training: ", hasil)

	// encoder_test := encoder.NewEncoder(labels_test)
	// encoder_test.Fit(labels_test)
	// labels_test_one_hot := encoder_test.Transform(labels_test)

	accuracy, loss, err := model.Test(test_loader, oneHotTargets)
	if err != nil {
		fmt.Println("Error during testing:", err)
	}
	fmt.Printf("Akurasi: %.2f%%, Loss: %.4f\n", accuracy*100, loss)
}
