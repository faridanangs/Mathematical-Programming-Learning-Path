package main

import (
	"fmt"
	"testing"
)

type Set map[string]bool

func NewSet(items ...string) Set {
	set := make(Set)
	for _, i := range items {
		set[i] = true
	}
	return set
}

func Intersection(set1, set2 Set) Set {
	inter := make(Set)
	for item := range set1 {
		if set2[item] {
			inter[item] = true
		}
	}
	return inter
}

// Fungsi untuk menghitung frekuensi kemunculan setiap item di seluruh pengguna lain
func CountItemFrequency(allUser map[string]Set) map[string]int {
	frequency := make(map[string]int)
	for _, userSet := range allUser {
		for item := range userSet {
			frequency[item]++
		}
	}
	return frequency
}

// Fungsi rekomendasi berdasarkan frekuensi item yang tertinggi
func Recomendation(userRef Set, allUser map[string]Set) {
	frequency := CountItemFrequency(allUser)
	var maxFrequency int
	recommendations := make([]string, 0)

	// Temukan frekuensi tertinggi
	for _, count := range frequency {
		if count > maxFrequency {
			maxFrequency = count
		}
	}

	// Tambahkan item dengan frekuensi tertinggi ke dalam rekomendasi
	for item, count := range frequency {
		if !userRef[item] && count == maxFrequency {
			recommendations = append(recommendations, item)
		}
	}

	fmt.Println("Rekomendasi HP yang paling banyak dimiliki oleh pengguna lain:")
	for _, item := range recommendations {
		fmt.Println(item)
	}
}

func TestMain(t *testing.T) {
	users := map[string]Set{
		"usman": NewSet("apel", "advan"),
		"joko":  NewSet("samsung", "advan", "somi"),
		"raka":  NewSet("lg", "samsung", "huwawe"),
		"iras":  NewSet("lg", "huwawe", "apel"),
		"suka":  NewSet("lg", "samsung", "huwawe"),
		"loki":  NewSet("advan", "apel", "huwawe"),
		"loli":  NewSet("advan", "somi", "huwawe"),
	}

	setUser := users["usman"]

	Recomendation(setUser, users)
}
