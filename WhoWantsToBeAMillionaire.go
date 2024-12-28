package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
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
	admins           = map[string]string{
		"admin1": "password1",
		"admin2": "password2",
	}
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

// Admin Functions

// addQuestion menambahkan soal baru ke bank soal.
func addQuestion() {
	if QuestionCount >= MaxQuestions {
		fmt.Println("Bank soal penuh.")
		return
	}

	var q Question
	reader := bufio.NewReader(os.Stdin) // Membuat reader untuk membaca input dengan spasi

	fmt.Println("\n--- Tambah Soal Baru ---")
	fmt.Print("Masukkan soal: ")
	q.Question, _ = reader.ReadString('\n') // Membaca input hingga baris baru

	fmt.Println("Masukkan 4 pilihan jawaban:")
	for i := 0; i < 4; i++ {
		fmt.Printf("Pilihan %d: ", i+1)
		q.Options[i], _ = reader.ReadString('\n') // Membaca input hingga baris baru
	}

	fmt.Print("Masukkan indeks jawaban yang benar (1-4): ")
	fmt.Scanln(&q.Answer)
	q.Answer--                             // Indeks jawaban dimulai dari 0
	q.ID = generateID(getAllQuestionIDs()) // Generate ID unik untuk soal

	Questions[QuestionCount] = q
	QuestionCount++
	fmt.Print("Soal berhasil ditambahkan.\n")
}

// editQuestion mengubah soal yang sudah ada berdasarkan ID.
func editQuestion() {
	fmt.Print("\nMasukkan ID soal yang ingin diubah: ")
	var id int
	fmt.Scanln(&id)
	index := sequentialSearchQuestions(id)
	if index == -1 {
		fmt.Println("Soal tidak ditemukan.")
		return
	}

	fmt.Println("\n--- Edit Soal ---")
	fmt.Printf("Soal lama: %s", Questions[index].Question)
	fmt.Print("Masukkan soal baru (kosongkan untuk mempertahankan soal lama): ")
	reader := bufio.NewReader(os.Stdin)
	newQuestion, _ := reader.ReadString('\n')
	if newQuestion != "\n" {
		Questions[index].Question = newQuestion
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("Pilihan %d [%s]: ", i+1, Questions[index].Options[i])
		newOption, _ := reader.ReadString('\n')
		newOption = strings.TrimSpace(newOption) // Trim the newline character
		if newOption != "" {
			Questions[index].Options[i] = newOption
		}
	}

	fmt.Print("Masukkan indeks jawaban yang benar (1-4): ")
	fmt.Scanln(&Questions[index].Answer)
	Questions[index].Answer-- // Indeks jawaban dimulai dari 0
	fmt.Print("Soal berhasil diperbarui.\n")
}

// deleteQuestion menghapus soal dari bank soal berdasarkan ID.
func deleteQuestion() {
	fmt.Print("\nMasukkan ID soal yang ingin dihapus: ")
	var id int
	fmt.Scanln(&id)
	index := sequentialSearchQuestions(id)
	if index == -1 {
		fmt.Println("Soal tidak ditemukan.")
		return
	}

	// Menggeser elemen setelah soal yang dihapus untuk menghindari celah dalam array
	for i := index; i < QuestionCount-1; i++ {
		Questions[i] = Questions[i+1]
	}
	QuestionCount--
	fmt.Print("Soal berhasil dihapus.\n")
}

// displayAllQuestions menampilkan semua soal yang tersedia.
func displayAllQuestions() {
	if QuestionCount == 0 {
		fmt.Println("Bank soal kosong.")
		return
	}

	fmt.Println("\n--- Daftar Soal ---")
	for i := 0; i < QuestionCount; i++ {
		fmt.Printf("ID: %d\n", Questions[i].ID)       // Menampilkan ID soal
		fmt.Printf("Soal: %s", Questions[i].Question) // Menampilkan teks soal
		fmt.Println("Pilihan Jawaban:")
		for j, option := range Questions[i].Options { // Menampilkan semua pilihan jawaban
			fmt.Printf("  %d. %s", j+1, option)
		}
		// Menampilkan jawaban yang benar
		fmt.Printf("Jawaban Benar: %d. %s", Questions[i].Answer+1, Questions[i].Options[Questions[i].Answer])
		fmt.Println("\n-----------------------------------") // Separator antar soal
	}
}

// getAllQuestionIDs mengembalikan daftar ID soal yang tersedia.
func getAllQuestionIDs() []int {
	ids := make([]int, 0, QuestionCount)
	for i := 0; i < QuestionCount; i++ {
		ids = append(ids, Questions[i].ID)
	}
	return ids
}

// sequentialSearchQuestions mencari soal berdasarkan ID.
func sequentialSearchQuestions(id int) int {
	for i := 0; i < QuestionCount; i++ {
		if Questions[i].ID == id {
			return i
		}
	}
	return -1
}

// displayMostAnsweredQuestions menampilkan soal yang paling banyak dijawab benar atau salah.
func displayMostAnsweredQuestions() {
	if QuestionCount == 0 {
		fmt.Println("Bank soal kosong.")
		return
	}

	fmt.Println("\n--- Soal dengan Jawaban Terbanyak ---")
	fmt.Println("5 soal dengan jawaban benar terbanyak:")
	sortQuestionsByCorrect()
	for i := 0; i < 5 && i < QuestionCount; i++ {
		fmt.Printf("ID: %d, Soal: %s, Benar: %d\n", Questions[i].ID, Questions[i].Question, Questions[i].Correct)
	}

	fmt.Println("\n5 soal dengan jawaban salah terbanyak:")
	sortQuestionsByWrong()
	for i := 0; i < 5 && i < QuestionCount; i++ {
		fmt.Printf("ID: %d, Soal: %s, Salah: %d\n", Questions[i].ID, Questions[i].Question, Questions[i].Wrong)
	}
}

// sortQuestionsByCorrect mengurutkan soal berdasarkan jumlah jawaban benar secara descending.
func sortQuestionsByCorrect() {
	for i := 0; i < QuestionCount-1; i++ {
		maxIdx := i
		for j := i + 1; j < QuestionCount; j++ {
			if Questions[j].Correct > Questions[maxIdx].Correct {
				maxIdx = j
			}
		}
		Questions[i], Questions[maxIdx] = Questions[maxIdx], Questions[i]
	}
}

// sortQuestionsByWrong mengurutkan soal berdasarkan jumlah jawaban salah secara descending.
func sortQuestionsByWrong() {
	for i := 0; i < QuestionCount-1; i++ {
		maxIdx := i
		for j := i + 1; j < QuestionCount; j++ {
			if Questions[j].Wrong > Questions[maxIdx].Wrong {
				maxIdx = j
			}
		}
		Questions[i], Questions[maxIdx] = Questions[maxIdx], Questions[i]
	}
}

// Participant Functions

// registerParticipant mendaftarkan peserta baru ke aplikasi.
func registerParticipant() {
	if ParticipantCount >= MaxParticipants {
		fmt.Println("Peserta sudah mencapai batas maksimum.")
		return
	}

	var p Participant
	fmt.Print("\nMasukkan nama Anda: ")
	fmt.Scanln(&p.Name)
	p.ID = generateID(getAllParticipantIDs()) // Generate ID unik untuk peserta

	Participants[ParticipantCount] = p
	ParticipantCount++
	fmt.Print("Pendaftaran berhasil! ID Anda: ", p.ID, "\n")
}

// getAllParticipantIDs mengembalikan daftar ID peserta yang terdaftar.
func getAllParticipantIDs() []int {
	ids := make([]int, 0, ParticipantCount)
	for i := 0; i < ParticipantCount; i++ {
		ids = append(ids, Participants[i].ID)
	}
	return ids
}

// takeQuiz memungkinkan peserta mengikuti kuis dengan soal acak.
func takeQuiz(participantID int) {
	// Verifikasi apakah ID peserta valid
	index := -1
	for i := 0; i < ParticipantCount; i++ {
		if Participants[i].ID == participantID {
			index = i
			i = ParticipantCount
		}
	}

	if index == -1 {
		fmt.Println("ID tidak terdaftar.")
		return
	}

	// Memulai kuis jika ID valid
	fmt.Println("\n--- Memulai Kuis ---")
	score := 0
	usedQuestions := make(map[int]bool)
	totalQuestions := QuestionCount

	// Memastikan semua soal bisa dijawab
	for len(usedQuestions) < totalQuestions {
		qIndex := rand.Intn(totalQuestions)
		for usedQuestions[qIndex] {
			qIndex = rand.Intn(totalQuestions)
		}
		usedQuestions[qIndex] = true

		q := Questions[qIndex]
		fmt.Printf("Soal %d: %s", len(usedQuestions), q.Question)
		for j, option := range q.Options {
			fmt.Printf("%d. %s", j+1, option)
		}
		fmt.Print("Masukkan jawaban Anda (1-4): ")
		var ans int
		fmt.Scanln(&ans)
		ans--

		if ans == q.Answer {
			fmt.Println("Benar!")
			score++
			Questions[qIndex].Correct++
		} else {
			fmt.Println("Salah! Jawaban yang benar:", q.Options[q.Answer])
			Questions[qIndex].Wrong++
		}
	}

	// Menampilkan skor akhir setelah selesai kuis
	fmt.Printf("Kuis selesai! Skor Anda: %d dari %d\n", score, totalQuestions)

	// Memperbarui skor peserta
	updateParticipantScore(participantID, score)
}

// updateParticipantScore memperbarui skor peserta setelah mengikuti kuis.
func updateParticipantScore(participantID, score int) {
	for i := 0; i < ParticipantCount; i++ {
		if Participants[i].ID == participantID {
			// Hanya memperbarui skor jika skor baru lebih tinggi
			if score > Participants[i].Score {
				Participants[i].Score = score
			}
			return
		}
	}
	// Jika ID tidak ditemukan (tidak seharusnya terjadi jika dicek sebelumnya)
	fmt.Println("Terjadi kesalahan saat memperbarui skor. ID tidak ditemukan.")
}

// Mencari user berdasarkan ID menggunakan binary search
func binarySearchParticipantID(participantID int) int {
	low := 0
	high := ParticipantCount - 1
	for low <= high {
		mid := low + (high-low)/2
		if Participants[mid].ID == participantID {
			return mid
		} else if Participants[mid].ID < participantID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// displayLeaderboard menampilkan papan skor peserta terurut berdasarkan skor tertinggi.
func displayLeaderboard() {
	hasScore := false
	for i := 0; i < ParticipantCount && !hasScore; i++ {
		if Participants[i].Score > 0 {
			hasScore = true
		}
	}

	if !hasScore {
		fmt.Println("Belum ada peserta yang mengerjakan kuis.")
		return
	}

	sortParticipantsByScore() // Mengurutkan peserta berdasarkan skor
	fmt.Println("\n--- Papan Skor ---")
	for i := 0; i < ParticipantCount; i++ {
		if Participants[i].Score > 0 {
			fmt.Printf("%d. %s - Skor: %d\n", i+1, Participants[i].Name, Participants[i].Score)
		}
	}
}

// sortParticipantsByScore menggunakan insertion sort untuk mengurutkan peserta berdasarkan skor secara descending.
func sortParticipantsByScore() {
    for i := 1; i < ParticipantCount; i++ {
        key := Participants[i]
        j := i - 1
        for j >= 0 && Participants[j].Score < key.Score {
            Participants[j+1] = Participants[j]
            j = j - 1
        }
        Participants[j+1] = key
    }
}

// Admin Authentication
func authenticateAdmin() bool {
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
							addQuestion()
						case 2:
							editQuestion()
						case 3:
							deleteQuestion()
						case 4:
							displayMostAnsweredQuestions()
						case 5:
							displayAllQuestions()
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
						registerParticipant()
					} else if participantChoice == 2 {
						fmt.Print("Masukkan ID peserta Anda: ")
						var participantID int
						fmt.Scanln(&participantID)
						takeQuiz(participantID)
					} else if participantChoice == 3 {
						displayLeaderboard()
					} else if participantChoice == 4 {
						fmt.Print("Masukkan ID peserta yang ingin dicari: ")
						var participantID int
						fmt.Scanln(&participantID)
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
		{ID: generateID(getAllQuestionIDs()), Question: "What is the capital of France?\n", Options: [4]string{"Berlin\n", "Madrid\n", "Paris\n", "Rome\n"}, Answer: 2},
		{ID: generateID(getAllQuestionIDs()), Question: "What is 2 + 2?\n", Options: [4]string{"3\n", "4\n", "5\n", "6\n"}, Answer: 1},
		{ID: generateID(getAllQuestionIDs()), Question: "What is the largest planet in our solar system?\n", Options: [4]string{"Earth\n", "Mars\n", "Jupiter\n", "Saturn\n"}, Answer: 2},
		{ID: generateID(getAllQuestionIDs()), Question: "What is the chemical symbol for water?\n", Options: [4]string{"H2O\n", "O2\n", "CO2\n", "NaCl\n"}, Answer: 0},
		{ID: generateID(getAllQuestionIDs()), Question: "Who wrote 'To Kill a Mockingbird'?\n", Options: [4]string{"Harper Lee\n", "Mark Twain\n", "Ernest Hemingway\n", "F. Scott Fitzgerald\n"}, Answer: 0},
	}

	for _, q := range staticQuestions {
		Questions[QuestionCount] = q
		QuestionCount++
	}
}