package hashes

import (
	"encoding/hex"
	"hash"
	"io"
	"strings"
)

func HashWithAllFunctions(txt string) string {
	sb := strings.Builder{}
	sb.WriteString("Text: " + txt + "\n\nHash function\t\t\tHashSum\n\n")
	for hf, hash := range hashes {
		sb.WriteString(hf + "\t\t\t" + hashSum(txt, hash) + "\n")
	}
	return sb.String()
}

func hashSum(text string, hash hash.Hash) string {
	io.WriteString(hash, text)
	return hex.EncodeToString(hash.Sum(nil))
}
