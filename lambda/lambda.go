package main

import (
	"fmt"
	"io/ioutil"
	"strings"
//	"math"
	//"github.com/skilstak/go-input"
	"sort"
)
 
func main() {
	//file := input.Ask("File name for sequence: ")
	file := "superoxidedismutases"
	//hook := input.Ask("Starting bits: ")
	hook := "GC"	
	//fenzyme := input.Ask("Forward Enzyme sequence: ")
	fenzyme := "AAGCTT"
	//renzyme := input.Ask("Reverse Enzyme sequence: ")
	renzyme := "GGATCC"
	forwardp, reversep := getPrimers(file, hook, fenzyme, renzyme)
	forwardp = meltingPointRank(forwardp)
	reversep = meltingPointRank(reversep)
	forwardp = gcRank(forwardp)
	reversep = gcRank(reversep)
	finalPairs := getFinalPairs(forwardp, reversep)
	finalPairs = tempDelta(finalPairs)
	fmt.Println(rankFinalPairs(finalPairs))

}
func getPrimers(file string, hook string, fenzyme string, renzyme string) (map[string][]int, map[string][]int) {
	fbegin := hook + fenzyme
	rbegin := hook + reverse(complement(renzyme))
	sequence := getFile(file)
	rsequence := reverse(complement(sequence))
	forwardp := make(map[string][]int)
	reversep := make(map[string][]int)
	var empty []int
	for n := (16 - len(fbegin)); n < 12; n++ {
		fmt.Println(n)
		forwardp[fbegin + sequence[0:n]] = empty
	}
	for n := (16 - len(rbegin)); n < 12; n++ {
		reversep[rbegin + rsequence[0:n]] = empty
	}
	return forwardp, reversep
}

func getFinalPairs(forward map[string][]int, reverse map[string][]int) map[[2]string][]int {
	finalPairs := make(map[[2]string][]int)
	for seq, rank := range forward {
		for rseq, rrank := range reverse {
			finalPairs[[2]string{seq, rseq}] = append(rank, rrank...)
		}
	}
	return finalPairs
}
func rankFinalPairs(finalPairs map[[2]string][]int) [2]string {
	var bestPrimers [2]string
	currentBestRank := 100000.0
	for pair, ranks := range finalPairs {
		if mean(ranks) < currentBestRank {
			bestPrimers = pair
			currentBestRank = mean(ranks)
		}
	}
	return bestPrimers
}
func mean(list []int) float64 {
        var total int
        for _, value := range list {
                total += value
        }
        return float64(total / len(list))
}
func sortByVal(m map[string]int) []string{
	s := make([]string, 0)	
    type kv struct {
        Key   string
        Value int
    }

    var ss []kv
    for k, v := range m {
        ss = append(ss, kv{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value < ss[j].Value
    })
	for _, kv := range ss {
		s = append(s, kv.Key)
	}
	return s
}
func sortPairByVal(m map[[2]string]int) [][2]string{
	s := make([][2]string, 0)	
    type kv struct {
        Key   [2]string
        Value int
    }

    var ss []kv
    for k, v := range m {
        ss = append(ss, kv{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value < ss[j].Value
    })
	for _, kv := range ss {
		s = append(s, kv.Key)
	}
	return s
} 
func reverse(strand string) string {
    r := []rune(strand)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
func complement(strand string) string {
	var returnValue string
	for _, char := range(strand) {
		switch string(char) {
		case "A":
			returnValue += "T"
		case "T":
			returnValue += "A"
		case "C":
			returnValue += "G"
		case "G":
			returnValue += "C"
		}
	}
	return returnValue
}
func meltingPointRank(primers map[string][]int) map[string][]int {
	unsortedmelt := make(map[string]int)
	for seq, _ := range primers {
		unsortedmelt[seq] = meltingPoint(seq)
	}
	sortedMelt := sortByVal(unsortedmelt)	
	for i, v := range sortedMelt {
		primers[v] = append(primers[v], i)
	}
	return primers
}
func meltingPoint(strand string) int {
	var A int
	var T int
	var C int
	var G int 
	for _, char := range(strand) {
		switch string(char) {
		case "A":
			A++
		case "T":
			T++
		case "C":
			C++
		case "G":
			G++
		}
	}
	return (59 - (2*(A + T) + 4*(C + G)) - 5)
	
}
func gcRank(primers map[string][]int) map[string][]int {
	unsortedgc := make(map[string]int)
	for seq, _ := range primers {
		unsortedgc[seq] = gc(seq)
	}
	sortedMelt := sortByVal(unsortedgc)	
	for i, v := range sortedMelt {
		primers[v] = append(primers[v], i)
	}
	return primers
}
func gc(strand string) int {
	var A int
	var T int
	var C int
	var G int 
	for _, char := range(strand) {
		switch string(char) {
		case "A":
			A++
		case "T":
			T++
		case "C":
			C++
		case "G":
			G++
		}
	}
	return (G +C )/ (A + T)

}
func tempDelta(finalPairs map[[2]string][]int) map[[2]string][]int  {
	unsortedTempDelta := make(map[[2]string]int)	
	for k, _ := range finalPairs {
		unsortedTempDelta[k] = meltingPoint(k[0]) - meltingPoint(k[1]) 
	}
	sortedTempDelta := sortPairByVal(unsortedTempDelta)
	for i, v := range sortedTempDelta {
		finalPairs[v] = append(finalPairs[v], i)
	}
	return finalPairs

}
func getFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	var returnValue string	
	check(err)
    for _, char := range strings.ToUpper(string(dat)) {
		if string(char) == "A" || string(char) == "T" || string(char) == "C" || string(char) == "G" {
			returnValue += string(char)
		}
	}
	return returnValue

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}