package main

import "fmt"

func soalKedua() {
	// Data 1
	markWeight1 := 78.0
	markHeight1 := 1.69

	johnWeight1 := 92.0
	johnHeight1 := 1.95

	// Mengitung BMI untuk Mark dan John pada Data 1
	markBMI1 := markWeight1 / (markHeight1 * markHeight1)
	johnBMI1 := johnWeight1 / (johnHeight1 * johnHeight1)

	// Menentukan apakah Mark memiliki BMI yang lebih tinggi dari John pada Data 1
	markHigherBMI1 := markBMI1 > johnBMI1

	// Menampilkan hasil Data 1
	fmt.Println("Data 1:")
	fmt.Printf("BMI Mark: %.2f\n", markBMI1)
	fmt.Printf("BMI John: %.2f\n", johnBMI1)
	fmt.Printf("Apakah Mark memiliki BMI lebih tinggi dari John pada Data 1? %t\n\n", markHigherBMI1)

	// Data 2
	markWeight2 := 95.0
	markHeight2 := 1.88

	johnWeight2 := 85.0
	johnHeight2 := 1.76

	// Menghitung BMI untuk Mark dan John pada Data 2
	markBMI2 := markWeight2 / (markHeight2 * markHeight2)
	johnBMI2 := johnWeight2 / (johnHeight2 * johnHeight2)

	// Menentukan apakah Mark memiliki BMI yang lebih tinggi dari John pada Data 2
	markHigherBMI2 := markBMI2 > johnBMI2

	// Menampilkan hasil Data 2
	fmt.Println("Data 2:")
	fmt.Printf("BMI Mark: %.2f\n", markBMI2)
	fmt.Printf("BMI John: %.2f\n", johnBMI2)
	fmt.Printf("Apakah Mark memiliki BMI lebih tinggi dari John pada Data 2? %t\n", markHigherBMI2)
}
