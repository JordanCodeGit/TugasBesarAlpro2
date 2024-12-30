package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Constants
const (
	MaxQuestions    = 100
	MaxParticipants = 100
)

// Data structures
type Question struct {
	ID       int
	Question string
	Options  [4]string
	Answer   int // Index of the correct option (0-3)
	Correct  int // Count of correct answers
	Wrong    int // Count of wrong answers
}

type Participant struct {
	ID    int
	Name  string
	Score int
}

// Global arrays and counters
var (
	Questions        [MaxQuestions]Question
	Participants     [MaxParticipants]Participant
	QuestionCount    int
	ParticipantCount int
)

// Utility functions
func generateID(existingIDs []int) int {
	id := rand.Intn(10000)
	for contains(existingIDs, id) {
		id = rand.Intn(10000)
	}
	return id
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// Main Function
// Fungsi utama menjalankan menu aplikasi dan pengelolaan fitur admin serta peserta.
func main() {
	fmt.Println("\n-=-=-=-=-=- WELCOME TO OUR PROGRAM 'WHO WANTS TO BE A MILLIONAIRE' -=-=-=-=-=-")
	fmt.Println("Program ini dibuat oleh Kelompok 4 Kelas IF-11-01")
	fmt.Println("1. Anita Nurazizah Agussalim 		(2311102017)")
	fmt.Println("2. Jordan Angkawijaya 			(2311102139)")
	fmt.Println("3. Danendra Arden Shaduq 		(2311102146)")
	fmt.Println("4. Dheva Dewa Septiantoni 		(2311102324)")
	rand.Seed(time.Now().UnixNano())

	var exit bool = false
	for !exit {
		fmt.Println("\n--- Menu Mode ---")
		fmt.Println("1. Mode Admin")
		fmt.Println("2. Mode Peserta")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih opsi: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if authenticateAdmin() {
				adminMode := true
				for adminMode {
					fmt.Println("\n--- Mode Admin ---")
					fmt.Println("1. Tambah Soal")
					fmt.Println("2. Edit Soal")
					fmt.Println("3. Hapus Soal")
					fmt.Println("4. Lihat Soal Terbanyak Benar/Salah")
					fmt.Println("5. Lihat Semua Soal")
					fmt.Println("6. Exit Admin Mode")
					fmt.Print("Pilih opsi: ")
					var adminChoice int
					fmt.Scanln(&adminChoice)
					switch adminChoice {
					case 1:
						addQuestion(&Questions, &QuestionCount)
					case 2:
						editQuestion(&Questions, QuestionCount)
					case 3:
						deleteQuestion(&Questions, &QuestionCount)
					case 4:
						displayMostAnsweredQuestions(&Questions, QuestionCount)
					case 5:
						displayAllQuestions(Questions, QuestionCount)
					case 6:
						adminMode = false
					default:
						fmt.Println("Opsi tidak valid, coba lagi.")
					}
				}
			}
		case 2:
			participantMode := true
			for participantMode {
				fmt.Println("\n--- Mode Peserta ---")
				fmt.Println("1. Daftar")
				fmt.Println("2. Ikuti Kuis")
				fmt.Println("3. Papan Skor")
				fmt.Println("4. Cari Peserta Berdasarkan ID")
				fmt.Println("5. Exit Participant Mode")
				fmt.Print("Pilih opsi: ")
				var participantChoice int
				fmt.Scanln(&participantChoice)
				if participantChoice == 1 {
					registerParticipant(&Participants, &ParticipantCount)
				} else if participantChoice == 2 {
					fmt.Print("Masukkan ID peserta Anda: ")
					var participantID int
					fmt.Scanln(&participantID)
					takeQuiz(participantID)
				} else if participantChoice == 3 {
					displayLeaderboard(&Participants, ParticipantCount)
				} else if participantChoice == 4 {
					fmt.Print("Masukkan ID peserta yang ingin dicari: ")
					var participantID int
					fmt.Scanln(&participantID)
					sortParticipantsByID(&Participants, ParticipantCount)
					index := binarySearchParticipantID(participantID)
					if index != -1 {
						fmt.Printf("Peserta ditemukan: %s dengan skor %d\n", Participants[index].Name, Participants[index].Score)
					} else {
						fmt.Println("Peserta tidak ditemukan.")
					}
				} else if participantChoice == 5 {
					participantMode = false
				} else {
					fmt.Println("Opsi tidak valid, coba lagi.")
				}
			}
		case 3:
			fmt.Println("Keluar... Sampai jumpa!")
			exit = true
		default:
			fmt.Println("Opsi tidak valid, coba lagi.")
		}
	}
}

// Initialize static questions
func init() {
	staticQuestions := []Question{
		{ID: generateID(getAllQuestionIDs(Questions, QuestionCount)), Question: "What is the capital of France?\n", Options: [4]string{"Berlin\n", "Madrid\n", "Paris\n", "Rome\n"}, Answer: 2},
		{ID: generateID(getAllQuestionIDs(Questions, QuestionCount)), Question: "What is 2 + 2?\n", Options: [4]string{"3\n", "4\n", "5\n", "6\n"}, Answer: 1},
		{ID: generateID(getAllQuestionIDs(Questions, QuestionCount)), Question: "What is the largest planet in our solar system?\n", Options: [4]string{"Earth\n", "Mars\n", "Jupiter\n", "Saturn\n"}, Answer: 2},
		{ID: generateID(getAllQuestionIDs(Questions, QuestionCount)), Question: "What is the chemical symbol for water?\n", Options: [4]string{"H2O\n", "O2\n", "CO2\n", "NaCl\n"}, Answer: 0},
		{ID: generateID(getAllQuestionIDs(Questions, QuestionCount)), Question: "Who wrote 'To Kill a Mockingbird'?\n", Options: [4]string{"Harper Lee\n", "Mark Twain\n", "Ernest Hemingway\n", "F. Scott Fitzgerald\n"}, Answer: 0},
	}

	for _, q := range staticQuestions {
		Questions[QuestionCount] = q
		QuestionCount++
	}
}
