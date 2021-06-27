package uuid

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func GenerateTimeUUID() string {
	t := time.Now().UnixNano()
	rand.Seed(t)

	ts := strconv.FormatInt(t, 10)
	rs := uuid.New().String()

	return fmt.Sprintf("%s_%s", ts, rs)
}
