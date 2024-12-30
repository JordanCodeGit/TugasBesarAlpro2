package main

// sortQuestionsByCorrect mengurutkan soal berdasarkan jumlah jawaban benar secara descending.
func sortQuestionsByCorrect(questions *[MaxQuestions]Question, questionCount int) {
	for i := 0; i < questionCount-1; i++ {
		maxIdx := i
		for j := i + 1; j < questionCount; j++ {
			if questions[j].Correct > questions[maxIdx].Correct {
				maxIdx = j
			}
		}
		questions[i], questions[maxIdx] = questions[maxIdx], questions[i]
	}
}

// sortQuestionsByCorrectAscending mengurutkan soal berdasarkan jumlah jawaban benar secara ascending.
func sortQuestionsByCorrectAscending(questions *[MaxQuestions]Question, questionCount int) {
    for i := 0; i < questionCount-1; i++ {
        minIdx := i
        for j := i + 1; j < questionCount; j++ {
            if questions[j].Correct < questions[minIdx].Correct {
                minIdx = j
            }
        }
        questions[i], questions[minIdx] = questions[minIdx], questions[i]
    }
}

// sortQuestionsByWrong mengurutkan soal berdasarkan jumlah jawaban salah secara descending.
func sortQuestionsByWrong(questions *[MaxQuestions]Question, questionCount int) {
	for i := 0; i < questionCount-1; i++ {
		maxIdx := i
		for j := i + 1; j < questionCount; j++ {
			if questions[j].Wrong > questions[maxIdx].Wrong {
				maxIdx = j
			}
		}
		questions[i], questions[maxIdx] = questions[maxIdx], questions[i]
	}
}

// sortQuestionsByWrongAscending mengurutkan soal berdasarkan jumlah jawaban salah secara ascending.
func sortQuestionsByWrongAscending(questions *[MaxQuestions]Question, questionCount int) {
    for i := 0; i < questionCount-1; i++ {
        minIdx := i
        for j := i + 1; j < questionCount; j++ {
            if questions[j].Wrong < questions[minIdx].Wrong {
                minIdx = j
            }
        }
        questions[i], questions[minIdx] = questions[minIdx], questions[i]
    }
}

// sortParticipantsByID mengurutkan peserta berdasarkan ID secara ascending.
func sortParticipantsByID(participants *[MaxParticipants]Participant, participantCount int) {
	for i := 1; i < participantCount; i++ {
		key := participants[i]
		j := i - 1
		for j >= 0 && participants[j].ID > key.ID {
			participants[j+1] = participants[j]
			j = j - 1
		}
		participants[j+1] = key
	}
}

// sortParticipantsByScore menggunakan insertion sort untuk mengurutkan peserta berdasarkan skor secara descending.
func sortParticipantsByScore(participants *[MaxParticipants]Participant, participantCount int) {
	for i := 1; i < participantCount; i++ {
		key := participants[i]
		j := i - 1
		for j >= 0 && participants[j].Score < key.Score {
			participants[j+1] = participants[j]
			j = j - 1
		}
		participants[j+1] = key
	}
}

// sortParticipantsByScoreAscending menggunakan insertion sort untuk mengurutkan peserta berdasarkan skor secara ascending.
func sortParticipantsByScoreAscending(participants *[MaxParticipants]Participant, participantCount int) {
    for i := 1; i < participantCount; i++ {
        key := participants[i]
        j := i - 1
        for j >= 0 && participants[j].Score > key.Score {
            participants[j+1] = participants[j]
            j = j - 1
        }
        participants[j+1] = key
    }
}