package main

import (
	"fmt"
	"neuralnetworks/pkg/dataset"
	"neuralnetworks/pkg/model"
)

func main() {
	PATH_TEST := "/Users/ghoziwaridi/PEMOGRAMAN/go/NeuralNetworks/data/raw/test"
	PATH_TRAIN := "/Users/ghoziwaridi/PEMOGRAMAN/go/NeuralNetworks/data/raw/train"

	// Load Dataset
	train_loader, labels_train := loader.DataLoader(PATH_TRAIN, true)
	test_loader, labels_test := loader.DataLoader(PATH_TEST, true)

	// --- Menampilkan Info Dataset Training ---
	// Menggunakan Printf dengan format verb %T untuk tipe data
	fmt.Printf("Tipe Train Loader: %T\n", train_loader)
	fmt.Printf("Tipe Train Labels: %T\n\n", labels_train)

	// Anda tetap bisa melihat nilainya jika perlu
	// fmt.Println("Train Dataset: ", train_loader)
	// fmt.Println("Labels: ", labels_train)

	// --- Menampilkan Info Dataset Testing ---
	fmt.Printf("Tipe Test Loader: %T\n", test_loader)
	fmt.Printf("Tipe Test Labels: %T\n", labels_test)

	model := model.InitHyperParameters(0.01, 764, 128, 64, 2)
	fmt.Println("Model Hyperparameters: ", model)

	hasil := model.Train(train_loader, labels_train, 10)

}
