package main

import (
	"fmt"
	"math/rand"
)

// Participant Functions

// registerParticipant mendaftarkan peserta baru ke aplikasi.
func registerParticipant(participants *[MaxParticipants]Participant, participantCount *int) {
	if *participantCount >= MaxParticipants {
		fmt.Println("Peserta sudah mencapai batas maksimum.")
		return
	}

	var p Participant
	fmt.Print("\nMasukkan nama Anda: ")
	fmt.Scanln(&p.Name)
	p.ID = generateID(getAllParticipantIDs(*participants, *participantCount))

	participants[*participantCount] = p
	(*participantCount)++
	fmt.Print("Pendaftaran berhasil! ID Anda: ", p.ID, "\n")
}

// getAllParticipantIDs mengembalikan daftar ID peserta yang terdaftar.
func getAllParticipantIDs(participants [MaxParticipants]Participant, participantCount int) []int {
	ids := make([]int, 0, participantCount)
	for i := 0; i < participantCount; i++ {
		ids = append(ids, participants[i].ID)
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
func displayLeaderboard(participants *[MaxParticipants]Participant, participantCount int) {
    hasScore := false
    for i := 0; i < participantCount && !hasScore; i++ {
        if participants[i].Score > 0 {
            hasScore = true
        }
    }

    if !hasScore {
        fmt.Println("Belum ada peserta yang mengerjakan kuis.")
        return
    }

    fmt.Println("\n--- Papan Skor ---")
    fmt.Println("1. Skor Tertinggi (Descending)")
    fmt.Println("2. Skor Terendah (Ascending)")
    fmt.Print("Pilih opsi: ")
    var choice int
    fmt.Scanln(&choice)

    switch choice {
    case 1:
        sortParticipantsByScore(participants, participantCount)
    case 2:
        sortParticipantsByScoreAscending(participants, participantCount)
    default:
        fmt.Println("Opsi tidak valid.")
        return
    }

    for i := 0; i < participantCount; i++ {
        if participants[i].Score > 0 {
            fmt.Printf("%d. %s - Skor: %d\n", i+1, participants[i].Name, participants[i].Score)
        }
    }
}