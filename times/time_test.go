package times

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeStampToAlmatyZone(t *testing.T) {
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
	{
		res, err := TimeStampToAlmatyZone("11/16/2023 16:44:06.000000")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}
	{
		res, err := TimeStampToAlmatyZone("11/16/2023 16:44:06")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}
	{
		res, err := TimeStampToAlmatyZone("2023-11-16 16:44:06")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}
	{
		res, err := TimeStampToAlmatyZone("2023/11/16 16:44:06")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}

}
func TestDateToAlmatyTime(t *testing.T) {
	expected := "2023-11-16T00:00:00+06:00"
	{
		res, err := DateToAlmatyTime("16/11/23")
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}
	{
		res, err := DateToAlmatyTime("16/11/2023")
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}
	{
		res, err := DateToAlmatyTime("11/16/2023 00:00:00.000000")
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}
}
