package times

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeStampToAlmatyZone(t *testing.T) {
	expected := "2023-11-16T16:44:06+06:00"
	InitTimeZone(6)
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

	{
		expected := "2020-02-10T00:00:00+06:00"
		res, err := TimeStampToAlmatyZone("02/10/2020 00:00:00")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}
	{
		expected := "2020-01-10T00:00:00+06:00"
		res, err := TimeStampToAlmatyZone("2020/01/10 00:00:00")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}
	{
		expected := "2006-01-02T15:04:05+06:00"
		res, err := TimeStampToAlmatyZone("01/02/2006 15:04:05.000000")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	}

}
func TestNegativeTimeStampToAlmatyZone(t *testing.T) {
	expected := ""
	{
		res, err := TimeStampToAlmatyZone("2023-11-16T16:44:06+06:00")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Equal(t, err, NotHasAlmatyZoneOffset)
		assert.Equal(t, expected, res)
	}
	{
		res, err := TimeStampToAlmatyZone("2023/11/16 16:44:06.000000")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Equal(t, err, NotHasAlmatyZoneOffset)
		assert.Equal(t, expected, res)
	}
	{
		res, err := TimeStampToAlmatyZone("11/16/2023 16:44:06.000000")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Equal(t, err, NotHasAlmatyZoneOffset)
		assert.Equal(t, expected, res)
	}
	{
		res, err := TimeStampToAlmatyZone("11/16/2023 16:44:06")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Equal(t, err, NotHasAlmatyZoneOffset)
		assert.Equal(t, expected, res)
	}
	{
		res, err := TimeStampToAlmatyZone("2023-11-16 16:44:06")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Equal(t, err, NotHasAlmatyZoneOffset)
		assert.Equal(t, expected, res)
	}
	{
		res, err := TimeStampToAlmatyZone("2023/11/16 16:44:06")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Equal(t, err, NotHasAlmatyZoneOffset)
		assert.Equal(t, expected, res)
	}

	{
		res, err := TimeStampToAlmatyZone("02/10/2020 00:00:00")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Equal(t, err, NotHasAlmatyZoneOffset)
		assert.Equal(t, expected, res)
	}
	{
		res, err := TimeStampToAlmatyZone("2020/01/10 00:00:00")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Equal(t, err, NotHasAlmatyZoneOffset)
		assert.Equal(t, expected, res)
	}
	{
		res, err := TimeStampToAlmatyZone("01/02/2006 15:04:05.000000")
		log.Printf("\n \n parsed time %+v \v\v err %+v", res, err)
		assert.Equal(t, err, NotHasAlmatyZoneOffset)
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
