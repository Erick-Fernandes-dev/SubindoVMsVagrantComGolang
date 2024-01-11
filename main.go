package main

import (
	"fmt"
	"os"
	"os/exec"
)

func option(op string) string {
	switch op {
	case "up":
		return "up"
	case "destroy":
		return "destroy"
	case "reload":
		return "reload"
	case "suspend":
		return "suspend"
	default:
		return ""
	}
}

func main() {

	var value string

	fmt.Println("Enter many options: up, destroy, reload")
	_, erro := fmt.Scanln(&value)

	if erro != nil {
		fmt.Println("erroo ao ler a entrada:", erro)
		return
	}

	// fmt.Println("Opção escolhida:", &value)
	// fmt.Println("Opção escolhida:", option(value))

	// Substitua "/caminho/para/o/diretorio" pelo caminho do diretório que contém o Vagrantfile
	diretorio := "/home/wolf/Documents/MinhasFormacoes/FormacaoKubernetes/01_mod/kubeadmin_cluster-vagrant"

	// // Altera o diretório de trabalho para o caminho especificado
	err := os.Chdir(diretorio)
	if err != nil {
		fmt.Println("Erro ao alterar o diretório de trabalho:", err)
		return
	}

	if option(value) != "" {

		// // Comando Vagrant a ser executado
		vagrantCommand := "vagrant"
		vagrantArgs := []string{option(value)}

		// // Executa o comando Vagrant
		cmd := exec.Command(vagrantCommand, vagrantArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println("Erro ao executar 'vagrant up':", err)
			return
		}

	} else {
		fmt.Println("Opção inválida")
	}
}
