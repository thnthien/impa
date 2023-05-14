package gorm

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon/v2"
)

type CarbonTime struct {
	carbon.Carbon
}

func (c *CarbonTime) Scan(value any) error {
	switch v := value.(type) {
	case []byte:
		c.Carbon = carbon.Parse(string(v))
	case string:
		c.Carbon = carbon.Parse(v)
	case time.Time:
		c.Carbon = carbon.FromStdTime(v)
	case nil:
		c.Carbon = carbon.NewCarbon()
	default:
		return fmt.Errorf("cannot scan to CarbonTime from %#v", v)
	}
	return nil
}
