// Declaração do pacote principal GO
package main

// Import da função fmt
import "fmt"

// Função Principal
func main() {

	var temperaturaKelvin float64

	fmt.Print("Digite a temperatura em Kelvin: ") // Impressão de Mensagem na tela do Usuário
	fmt.Scanln(&temperaturaKelvin)                // Coleta os dados digitados pelo usuário

	temperaturaCelsius := kelvinToCelsius(temperaturaKelvin) // Atribução da função de conversão

	fmt.Printf("%.2f Kelvin é equivalente a %.2f Celsius\n", temperaturaKelvin, temperaturaCelsius) // Impressão do resultado Final

}

/* Funcões adicionais e melhor legibilidade e reaproveitamento de codigo futuro
==================================================================================*/

// Função para converter da escala Kelvin para Celsius
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}
