package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	inicio := time.Now()
	generateFile("alunos.csv", generateStudents(12000000))

	// Ler arquivo
	file, err := os.OpenFile("alunos.csv", os.O_RDWR|os.O_APPEND, os.ModeAppend) // Open file in edition mode
	if err != nil {
		log.Fatalf("Read file: %v", err)
	}

	// Pegar o conteúdo
	r := csv.NewReader(file)
	var resultados [][]string
	for {
		linha, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		nota1, _ := strconv.ParseFloat(linha[1], 64)
		nota2, _ := strconv.ParseFloat(linha[2], 64)
		nota3, _ := strconv.ParseFloat(linha[3], 64)
		nota4, _ := strconv.ParseFloat(linha[4], 64)
		media := calcMedia(nota1, nota2, nota3, nota4)
		resultados = append(resultados, []string{linha[0], fmt.Sprintf("%.2f", media), getResult(media)})
	}
	generateFile("resultado.csv", resultados)
	fim := time.Now()
	fmt.Println("Tempo de execução: ", fim.Sub(inicio))
}

func calcMedia(nota1 float64, nota2 float64, nota3 float64, nota4 float64) float64 {
	return (nota1 + nota2 + nota3 + nota4) / 4
}

func getResult(media float64) string {
	if media >= 7.0 {
		return "Aprovado"
	}
	return "Reprovado"
}

func generateFile(filename string, content [][]string) {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("Create file: %v", err)
	}

	w := csv.NewWriter(file)
	w.WriteAll(content)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	defer file.Close()

	fmt.Println(filename, "created with", len(content), "rows")
}

func generateStudents(qtd int) [][]string {
	var alunos [][]string

	for i := 0; i < qtd; i++ {
		nomeAluno := fmt.Sprintf("Aluno %v", strconv.Itoa(i+1))
		nota1 := fmt.Sprintf("%.2f", rand.Float64()*10)
		nota2 := fmt.Sprintf("%.2f", rand.Float64()*10)
		nota3 := fmt.Sprintf("%.2f", rand.Float64()*10)
		nota4 := fmt.Sprintf("%.2f", rand.Float64()*10)

		alunos = append(alunos, []string{nomeAluno, nota1, nota2, nota3, nota4})
	}

	return alunos
}
