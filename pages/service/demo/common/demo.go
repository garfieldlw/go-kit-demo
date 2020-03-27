package common

import (
	"context"
)

func GetValue(ctx context.Context, req string) (string, error) {
	return req + "!!!", nil
}
