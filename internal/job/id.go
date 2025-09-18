package job

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type ID string

func MakeID(title, company string, date time.Time, url url.URL) ID {
	seed := title +
		"|" + company +
		"|" + date.Format("02-01-2006") +
		"|" + url.String()
	base := strings.ToLower(seed)
	sum := sha1.Sum([]byte(base))
	id := hex.EncodeToString(sum[:])[:10]
	return ID(id)
}

func MakeIDFromString(s string) (ID, error) {
	if len(s) != 10 {
		return "", fmt.Errorf("given string of length: %d", len(s))
	}
	return ID(s), nil
}

func (id ID) String() string {
	return string(id)
}
