package task5

func sanitize(in string) string {
	out := []rune{}
	for _, v := range in {
		if 65 <= v && v <= 90 {
			out = append(out, v)
		} else if 97 <= v && v <= 122 {
			out = append(out, v-32)
		}
	}

	return string(out)
}

func encodePair(a, b rune) rune {
	return (((a - 'A') + (b - 'A')) % 26) + 'A'
}

func decodePair(a, b rune) rune {
	return (((((a - 'A') - (b - 'A')) + 26) % 26) + 'A')
}

func Cryptor(msg, key string) string {
	smsg, skey := sanitize(msg), sanitize(key)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, encodePair(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}

func Decryptor(msg, key string) string {
	smsg, skey := sanitize(msg), sanitize(key)
	out := make([]rune, 0, len(msg))
	for i, v := range smsg {
		out = append(out, decodePair(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}
