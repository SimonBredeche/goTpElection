package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Vote struct {
	candidates  map[string]int
	departement string
	nbVote      int
}

func main() {
	// open file
	readFile, err := os.Open("resultats-par-niveau-burvot-t1-france-entiere.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	totalOfVote := 0
	firstLine := true
	voteByCandidate := make(map[string]int)
	voteByCandidateByDepartement := make(map[string]int)
	rankingByDepartement := make(map[string]int)
	//Custom parsing for memory optimisations
	for fileScanner.Scan() {
		if firstLine {
			firstLine = false
			continue
		}
		vote := createEntryFromString(fileScanner.Text())
		for key, value := range vote.candidates {
			candidateDepartementKey := key + "_" + vote.departement
			voteByCandidateByDepartement[candidateDepartementKey] += value
			voteByCandidate[key] += value
		}
		rankingByDepartement[vote.departement] += vote.nbVote
		totalOfVote += vote.nbVote
	}

	fmt.Print("Nombre de vote : ", totalOfVote)

	printVoteByCandidate(voteByCandidate)

	printVoteByCandidateByDepartement(voteByCandidateByDepartement)

	printRankingByCirconscription(rankingByDepartement)

}

func printVoteByCandidate(voteByCandidate map[string]int) {
	for key, value := range voteByCandidate {
		fmt.Print("\n Candidate : ", key, " Number of vote : ", value)
	}
}

func printVoteByCandidateByDepartement(voteByCandidateByDepartement map[string]int) {
	for key, value := range voteByCandidateByDepartement {
		splitedKey := strings.Split(key, "_")
		candidat := splitedKey[0]
		departement := splitedKey[1]
		fmt.Print("\n Departement : ", departement, " Candidate : ", candidat, " Number of vote : ", value)
	}
}

func printRankingByCirconscription(rankingByDepartement map[string]int) {

	ranking := make([]string, 0, len(rankingByDepartement))

	for key := range rankingByDepartement {
		ranking = append(ranking, key)
	}

	sort.SliceStable(ranking, func(i, j int) bool {
		return rankingByDepartement[ranking[i]] > rankingByDepartement[ranking[j]]
	})

	for i := 0; i < len(ranking); i++ {
		fmt.Print("\n#", i+1, " : ", ranking[i])
	}
}

func createEntryFromString(data string) Vote {
	splitedData := strings.Split(data, ";")
	vote := &Vote{}
	vote.candidates = parseCandidates(splitedData)
	vote.departement = splitedData[1]
	intVar, err := strconv.Atoi(splitedData[10])
	if err != nil {
		fmt.Print(err.Error())
	}
	vote.nbVote = intVar
	return *vote
}

func parseCandidates(splitedData []string) map[string]int {
	candidat := make(map[string]int)
	voteByCandidatStartIndex := 23
	numberOfRowBetweenCandidates := 7
	for i := voteByCandidatStartIndex; i < len(splitedData); i += numberOfRowBetweenCandidates {
		if i > len(splitedData)-1 {
			break
		}
		candidatesName := splitedData[i]
		candidatesVoteStr := splitedData[i+2]
		candidatesVote, _ := strconv.Atoi(candidatesVoteStr)
		candidat[candidatesName] = candidatesVote
	}
	return candidat
}
