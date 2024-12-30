package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Admin Functions

// Admin Authentication
func authenticateAdmin() bool {
    admins := map[string]string{
        "admin1": "password1",
        "admin2": "password2",
    }

    var username, password string
    fmt.Print("Masukkan username: ")
    fmt.Scanln(&username)
    username = strings.TrimSpace(username) // Remove newline character and any surrounding whitespace

    fmt.Print("Masukkan password: ")
    fmt.Scanln(&password)
    password = strings.TrimSpace(password) // Remove newline character and any surrounding whitespace

    if pass, ok := admins[username]; ok && pass == password {
        return true
    }
    fmt.Println("Username atau password salah.")
    return false
}

// addQuestion menambahkan soal baru ke bank soal.
// Parameter:
// - questions: pointer ke array soal
// - questionCount: pointer ke jumlah soal saat ini
func addQuestion(questions *[MaxQuestions]Question, questionCount *int) {
	if *questionCount >= MaxQuestions {
		fmt.Println("Bank soal penuh.")
		return
	}

	var q Question
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- Tambah Soal Baru ---")
	fmt.Print("Masukkan soal: ")
	q.Question, _ = reader.ReadString('\n')

	fmt.Println("Masukkan 4 pilihan jawaban:")
	for i := 0; i < 4; i++ {
		fmt.Printf("Pilihan %d: ", i+1)
		q.Options[i], _ = reader.ReadString('\n')
	}

	fmt.Print("Masukkan indeks jawaban yang benar (1-4): ")
	fmt.Scanln(&q.Answer)
	q.Answer--

	q.ID = generateID(getAllQuestionIDs(*questions, *questionCount))
	questions[*questionCount] = q
	(*questionCount)++
	fmt.Print("Soal berhasil ditambahkan.\n")
}

// editQuestion mengubah soal yang sudah ada berdasarkan ID.
// Parameter:
// - questions: pointer ke array soal
// - questionCount: jumlah soal saat ini
func editQuestion(questions *[MaxQuestions]Question, questionCount int) {
	fmt.Print("\nMasukkan ID soal yang ingin diubah: ")
	var id int
	fmt.Scanln(&id)
	index := sequentialSearchQuestions(*questions, questionCount, id)
	if index == -1 {
		fmt.Println("Soal tidak ditemukan.")
		return
	}

	fmt.Println("\n--- Edit Soal ---")
	fmt.Printf("Soal lama: %s", questions[index].Question)
	fmt.Print("Masukkan soal baru (kosongkan untuk mempertahankan soal lama): ")
	reader := bufio.NewReader(os.Stdin)
	newQuestion, _ := reader.ReadString('\n')
	if newQuestion != "\n" {
		questions[index].Question = newQuestion
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("Pilihan %d [%s]: ", i+1, questions[index].Options[i])
		newOption, _ := reader.ReadString('\n')
		newOption = strings.TrimSpace(newOption)
		if newOption != "" {
			questions[index].Options[i] = newOption
		}
	}

	fmt.Print("Masukkan indeks jawaban yang benar (1-4): ")
	fmt.Scanln(&questions[index].Answer)
	questions[index].Answer--
	fmt.Print("Soal berhasil diperbarui.\n")
}

// deleteQuestion menghapus soal dari bank soal berdasarkan ID.
// Parameter:
// - questions: pointer ke array soal
// - questionCount: pointer ke jumlah soal saat ini
func deleteQuestion(questions *[MaxQuestions]Question, questionCount *int) {
	fmt.Print("\nMasukkan ID soal yang ingin dihapus: ")
	var id int
	fmt.Scanln(&id)
	index := sequentialSearchQuestions(*questions, *questionCount, id)
	if index == -1 {
		fmt.Println("Soal tidak ditemukan.")
		return
	}

	for i := index; i < *questionCount-1; i++ {
		questions[i] = questions[i+1]
	}
	(*questionCount)--
	fmt.Print("Soal berhasil dihapus.\n")
}

// displayAllQuestions menampilkan semua soal yang tersedia.
// Parameter:
// - questions: array soal
// - questionCount: jumlah soal saat ini
func displayAllQuestions(questions [MaxQuestions]Question, questionCount int) {
	if questionCount == 0 {
		fmt.Println("Bank soal kosong.")
		return
	}

	fmt.Println("\n--- Daftar Soal ---")
	for i := 0; i < questionCount; i++ {
		fmt.Printf("ID: %d\n", questions[i].ID)
		fmt.Printf("Soal: %s", questions[i].Question)
		fmt.Println("Pilihan Jawaban:")
		for j, option := range questions[i].Options {
			fmt.Printf("  %d. %s", j+1, option)
		}
		fmt.Printf("Jawaban Benar: %d. %s", questions[i].Answer+1, questions[i].Options[questions[i].Answer])
		fmt.Println("\n-----------------------------------")
	}
}

// getAllQuestionIDs mengembalikan daftar ID soal yang tersedia.
func getAllQuestionIDs(questions [MaxQuestions]Question, questionCount int) []int {
	ids := make([]int, 0, questionCount)
	for i := 0; i < questionCount; i++ {
		ids = append(ids, questions[i].ID)
	}
	return ids
}

// displayMostAnsweredQuestions menampilkan soal yang paling banyak dijawab benar atau salah.
func displayMostAnsweredQuestions(questions *[MaxQuestions]Question, questionCount int) {
    if questionCount == 0 {
        fmt.Println("Bank soal kosong.")
        return
    }

    fmt.Println("\n--- Soal dengan Jawaban Terbanyak ---")
    fmt.Println("1. Jawaban Benar Terbanyak (Descending)")
    fmt.Println("2. Jawaban Benar Terbanyak (Ascending)")
    fmt.Println("3. Jawaban Salah Terbanyak (Descending)")
    fmt.Println("4. Jawaban Salah Terbanyak (Ascending)")
    fmt.Print("Pilih opsi: ")
    var choice int
    fmt.Scanln(&choice)

    switch choice {
    case 1:
        sortQuestionsByCorrect(questions, questionCount)
    case 2:
        sortQuestionsByCorrectAscending(questions, questionCount)
    case 3:
        sortQuestionsByWrong(questions, questionCount)
    case 4:
        sortQuestionsByWrongAscending(questions, questionCount)
    default:
        fmt.Println("Opsi tidak valid.")
        return
    }

    for i := 0; i < 5 && i < questionCount; i++ {
        fmt.Printf("ID: %d, Soal: %s, Benar: %d, Salah: %d\n", questions[i].ID, questions[i].Question, questions[i].Correct, questions[i].Wrong)
    }
}

// sequentialSearchQuestions mencari soal berdasarkan ID.
// Parameter:
// - questions: array soal
// - questionCount: jumlah soal saat ini
// - id: ID soal yang dicari
// Return: indeks soal jika ditemukan, -1 jika tidak ditemukan
func sequentialSearchQuestions(questions [MaxQuestions]Question, questionCount, id int) int {
	for i := 0; i < questionCount; i++ {
		if questions[i].ID == id {
			return i
		}
	}
	return -1
}