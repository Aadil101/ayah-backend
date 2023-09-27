package internal

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

var verseIDCumSum = []int{
	7, 293, 493, 669, 789, 954, 1160, 1235, 1364, 1473,
	1596, 1707, 1750, 1802, 1901, 2029, 2140, 2250, 2348, 2483,
	2595, 2673, 2791, 2855, 2932, 3159, 3252, 3340, 3409, 3469,
	3503, 3533, 3606, 3660, 3705, 3788, 3970, 4058, 4133, 4218,
	4272, 4325, 4414, 4473, 4510, 4545, 4583, 4612, 4630, 4675,
	4735, 4784, 4846, 4901, 4979, 5075, 5104, 5126, 5150, 5163,
	5177, 5188, 5199, 5217, 5229, 5241, 5271, 5323, 5375, 5419,
	5447, 5475, 5495, 5551, 5591, 5622, 5672, 5712, 5758, 5800,
	5829, 5848, 5884, 5909, 5931, 5948, 5967, 5993, 6023, 6043,
	6058, 6079, 6090, 6098, 6106, 6125, 6130, 6138, 6146, 6157,
	6168, 6176, 6179, 6188, 6193, 6197, 6204, 6207, 6213, 6216,
	6221, 6225, 6230, 6236,
}

func getVerseIDCumSum() []int {
	f, err := os.Open("/home/code_is_life/ayah-backend/pkg/internal/quran_stats.csv")
	if err != nil {
		log.Fatal("Error while reading .CSV file: ", err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = '\t'
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error while reading records: ", err)
	}

	result := make([]int, nChapters)

	for i, r := range records[1:] {
		nVerses, err := strconv.Atoi(r[2])
		if err != nil {
			log.Fatal("Error while converting string to int", err)
		}
		if i == 0 {
			result[i] = nVerses
		} else {
			result[i] = result[i-1] + nVerses
		}
	}

	return result
}
