package go2php

// Permutation 排列组合

func Permutation(SS *[]string, Str string, index int) {
	S := []byte(Str)
	if index == len(S) {
		*SS = append(*SS, string(S))
		return
	}
	for i := index; i < len(S); i++ {
		S[i], S[index] = S[index], S[i]
		Permutation(SS, string(S), index+1)
		S[i], S[index] = S[index], S[i]
	}
}
