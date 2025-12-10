package dashboard

import (
	"time"

	"github.com/go-resty/resty/v2"
)

var Resty = resty.New().
	SetTimeout(5 * time.Second)
