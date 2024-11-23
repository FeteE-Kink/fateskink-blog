package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}

func GqlIdtoInt32(i graphql.ID) (int32, error) {
	r, err := strconv.ParseInt(string(i), 10, 32)

	if err != nil {
		return 0, errors.Wrap(err, "GqlIDToUint")
	}

	return int32(r), nil
}

func GetStringDefault(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func GetInt32OrDefault(num *int32) int32 {
	if num == nil {
		return 0
	}
	return *num
}

func GetIntOrDefault(num *int) int {
	if num == nil {
		return 0
	}
	return *num
}

func GqlIDPoiter[T SignedInteger](id T) *graphql.ID {
	if id == 0 {
		return nil
	}

	r := graphql.ID(fmt.Sprint(id))
	return &r
}

func Int32Pointer(num int32) *int32 {
	return &num
}

func GqlIDValue[T SignedInteger](id T) graphql.ID {
	return graphql.ID(fmt.Sprint(id))
}

func GetBoolOrDefault(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func NewUUID() string {
	newUUID := uuid.New().String()

	idSplitted := strings.Split(newUUID, "-")
	idJoined := strings.Join(idSplitted[:], "")

	return idJoined
}

func Int32Poiter(num int32) *int32 {
	return &num
}

func TimePointer(val time.Time) *time.Time {
	return &val
}

func IDpoiter(id graphql.ID) *graphql.ID {
	return &id
}

func GqlTimePointer(val *time.Time) *graphql.Time {
	if val != nil {
		time := graphql.Time{Time: *val}

		return &time
	}
	return nil
}
