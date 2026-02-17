package main

func isIsomorphic(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	s1Map := make(map[byte]byte)
	s2Map := make(map[byte]byte)

	for i := 0; i < len(s1); i++ {
		s1Char := s1[i]
		s2Char := s2[i]
		char, found := s1Map[s1Char]
		if found {
			if char != s2Char {
				return false
			}
		} else {
			s1Map[s1Char] = s2Char
		}

		ch, ok := s2Map[s2Char]
		if ok {
			if ch != s1Char {
				return false
			}
		} else {
			s2Map[s2Char] = s1Char
		}
	}
	return true
}
