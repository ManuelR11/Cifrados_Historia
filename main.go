package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ManuelR11/Cifrados_Historia/affine"
	"github.com/ManuelR11/Cifrados_Historia/atbash"
	"github.com/ManuelR11/Cifrados_Historia/caesar"
	"github.com/ManuelR11/Cifrados_Historia/rot13"
	"github.com/ManuelR11/Cifrados_Historia/sustitucion"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("=== Menú de Cifrados ===")
		fmt.Println("1. Cifrado César")
		fmt.Println("2. Cifrado de Sustitución")
		fmt.Println("3. Cifrado Atbash")
		fmt.Println("4. Cifrado Afín")
		fmt.Println("5. Cifrado ROT13")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			for {
				fmt.Println("=== Submenú: Cifrado César ===")
				fmt.Println("1. Cifrar")
				fmt.Println("2. Decodificar")
				fmt.Println("3. Regresar")
				fmt.Print("Seleccione una opción: ")

				subOption, _ := reader.ReadString('\n')
				subOption = strings.TrimSpace(subOption)

				switch subOption {
				case "1":
					fmt.Print("Ingrese el texto a cifrar: ")
					text, _ := reader.ReadString('\n')
					text = strings.TrimSpace(text)
					fmt.Print("Ingrese el desplazamiento (clave): ")
					var shift int
					fmt.Scanf("%d", &shift)
					reader.ReadString('\n') // Limpia el buffer después de Scanf
					encrypted := caesar.Encrypt(text, shift)
					fmt.Println("Texto cifrado:", encrypted)
				case "2":
					fmt.Print("Ingrese el texto a decodificar: ")
					text, _ := reader.ReadString('\n')
					text = strings.TrimSpace(text)
					fmt.Print("Ingrese el desplazamiento (clave): ")
					var shift int
					fmt.Scanf("%d", &shift)
					reader.ReadString('\n') // Limpia el buffer después de Scanf
					decrypted := caesar.Decrypt(text, shift)
					fmt.Println("Texto decodificado:", decrypted)
				case "3":
					break
				default:
					fmt.Println("Opción no válida.")
				}
				if subOption == "3" {
					break
				}
			}

		case "2":
			for {
				fmt.Println("=== Submenú: Cifrado de Sustitución ===")
				fmt.Println("1. Cifrar")
				fmt.Println("2. Decodificar")
				fmt.Println("3. Regresar")
				fmt.Print("Seleccione una opción: ")

				subOption, _ := reader.ReadString('\n')
				subOption = strings.TrimSpace(subOption)

				switch subOption {
				case "1":
					fmt.Print("Ingrese el texto a cifrar: ")
					text, _ := reader.ReadString('\n')
					text = strings.TrimSpace(text)
					encrypted, err := sustitucion.Encrypt(text)
					if err != nil {
						fmt.Println("Error:", err)
					} else {
						fmt.Println("Texto cifrado:", encrypted)
					}
				case "2":
					fmt.Print("Ingrese el texto a decodificar: ")
					text, _ := reader.ReadString('\n')
					text = strings.TrimSpace(text)
					decrypted, err := sustitucion.Decrypt(text)
					if err != nil {
						fmt.Println("Error:", err)
					} else {
						fmt.Println("Texto decodificado:", decrypted)
					}
				case "3":
					break
				default:
					fmt.Println("Opción no válida.")
				}
				if subOption == "3" {
					break
				}
			}

		case "3":
			for {
				fmt.Println("=== Submenú: Cifrado Atbash ===")
				fmt.Println("1. Cifrar")
				fmt.Println("2. Decodificar")
				fmt.Println("3. Regresar")
				fmt.Print("Seleccione una opción: ")

				subOption, _ := reader.ReadString('\n')
				subOption = strings.TrimSpace(subOption)

				switch subOption {
				case "1":
					fmt.Print("Ingrese el texto a cifrar: ")
					text, _ := reader.ReadString('\n')
					text = strings.TrimSpace(text)
					encrypted := atbash.Encrypt(text)
					fmt.Println("Texto cifrado:", encrypted)
				case "2":
					fmt.Print("Ingrese el texto a decodificar: ")
					text, _ := reader.ReadString('\n')
					text = strings.TrimSpace(text)
					decrypted := atbash.Decrypt(text)
					fmt.Println("Texto decodificado:", decrypted)
				case "3":
					break
				default:
					fmt.Println("Opción no válida.")
				}
				if subOption == "3" {
					break
				}
			}

		case "4":
			for {
				fmt.Println("=== Submenú: Cifrado Afín ===")
				fmt.Println("1. Cifrar")
				fmt.Println("2. Decodificar")
				fmt.Println("3. Regresar")
				fmt.Print("Seleccione una opción: ")

				subOption, _ := reader.ReadString('\n')
				subOption = strings.TrimSpace(subOption)

				switch subOption {
				case "1":
					fmt.Print("Ingrese el texto a cifrar: ")
					text, _ := reader.ReadString('\n')
					text = strings.TrimSpace(text)
					fmt.Print("Ingrese el valor de a (1,2,4,5,7,8,10,11,13,14,16,17,19,20,22,23,25): ")
					var a int
					fmt.Scanf("%d", &a)
					reader.ReadString('\n') // Limpia el buffer
					fmt.Print("Ingrese el valor de b: ")
					var b int
					fmt.Scanf("%d", &b)
					reader.ReadString('\n') // Limpia el buffer
					encrypted, err := affine.Encrypt(text, a, b)
					if err != nil {
						fmt.Println("Error:", err)
					} else {
						fmt.Println("Texto cifrado:", encrypted)
					}
				case "2":
					fmt.Print("Ingrese el texto a decodificar: ")
					text, _ := reader.ReadString('\n')
					text = strings.TrimSpace(text)
					fmt.Print("Ingrese el valor de a: ")
					var a int
					fmt.Scanf("%d", &a)
					reader.ReadString('\n') // Limpia el buffer
					fmt.Print("Ingrese el valor de b: ")
					var b int
					fmt.Scanf("%d", &b)
					reader.ReadString('\n') // Limpia el buffer
					decrypted, err := affine.Decrypt(text, a, b)
					if err != nil {
						fmt.Println("Error:", err)
					} else {
						fmt.Println("Texto decodificado:", decrypted)
					}
				case "3":
					break
				default:
					fmt.Println("Opción no válida.")
				}
				if subOption == "3" {
					break
				}
			}
		case "5":
			for {
				fmt.Println("=== Submenú: Cifrado ROT13 ===")
				fmt.Println("1. Cifrar")
				fmt.Println("2. Decodificar")
				fmt.Println("3. Regresar")
				fmt.Print("Seleccione una opción: ")

				subOption, _ := reader.ReadString('\n')
				subOption = strings.TrimSpace(subOption)

				switch subOption {
				case "1":
					fmt.Print("Ingrese el texto a cifrar: ")
					text, _ := reader.ReadString('\n')
					text = strings.TrimSpace(text)
					encrypted := rot13.Encrypt(text)
					fmt.Println("Texto cifrado:", encrypted)
				case "2":
					fmt.Print("Ingrese el texto a decodificar: ")
					text, _ := reader.ReadString('\n')
					text = strings.TrimSpace(text)
					decrypted := rot13.Decrypt(text)
					fmt.Println("Texto decodificado:", decrypted)
				case "3":
					break
				default:
					fmt.Println("Opción no válida.")
				}
				if subOption == "3" {
					break
				}
			}

		case "6":
			fmt.Println("Saliendo del programa.")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}
