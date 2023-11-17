package times

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDocumentType(t *testing.T) {
	expected := "2023-11-16T16:44:06+06:00"
	{
		res, err := TimeStampToAlmatyZone("2023-11-16T16:44:06+06:00")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}
	{
		res, err := TimeStampToAlmatyZone("2023/11/16 16:44:06.000000")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}
}
