package calendar

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllEvents(c *gin.Context) {
	var ShortEventsList map[time.Time][]ShortEventExpanded
	var LongEventsList map[time.Time][]LongEventExpanded

	// Short categories
	for date, _ := range AllShortEvents {
		var emptyPointEvents []ShortEventExpanded
		ShortEventsList[date] = emptyPointEvents

		for _, targetShortEvent := range AllShortEvents[date] {
			ShortEventsList[date] = append(ShortEventsList[date], GetShortCategory(targetShortEvent))
		}
	}

	// Long categories
	for date, _ := range LongEventsList {
		var emptyPointEvents []LongEventExpanded
		LongEventsList[date] = emptyPointEvents

		for _, targetLongEvent := range AllLongEvents[date] {
			LongEventsList[date] = append(LongEventsList[date], GetLongCategory(targetLongEvent))
		}
	}

	fullList := map[string]interface{}{
		"short": ShortEventsList,
		"long":  LongEventsList,
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    fullList,
	})
}

func GetShortEvents(c *gin.Context) {
	var ShortEventsList map[time.Time][]ShortEventExpanded

	for date, _ := range AllShortEvents {
		var emptyPointEvents []ShortEventExpanded
		ShortEventsList[date] = emptyPointEvents

		for _, targetShortEvent := range AllShortEvents[date] {
			ShortEventsList[date] = append(ShortEventsList[date], GetShortCategory(targetShortEvent))
		}
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    ShortEventsList,
	})
}

func GetLongEvents(c *gin.Context) {
	var LongEventsList map[time.Time][]LongEventExpanded

	for date, _ := range LongEventsList {
		var emptyPointEvents []LongEventExpanded
		LongEventsList[date] = emptyPointEvents

		for _, targetLongEvent := range AllLongEvents[date] {
			LongEventsList[date] = append(LongEventsList[date], GetLongCategory(targetLongEvent))
		}
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    LongEventsList,
	})
}

func CreateShortEvent(c *gin.Context) {
	var request NewShortEventReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	categoryExist := CheckCategoryExist(request.Category)

	if !categoryExist {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Category is not exist",
			"body":    nil,
		})
		return
	}

	newShortEvent := ShortEvent{
		ID:          CurrentShortEventID,
		Category:    request.Category,
		Title:       request.Title,
		Description: request.Description,
		Time:        request.Time,
	}

	_, ok := AllShortEvents[request.Time]
	if ok {
		AllShortEvents[request.Time] = append(AllShortEvents[request.Time], newShortEvent)
	} else {
		AllShortEvents[request.Time] = []ShortEvent{newShortEvent}
	}

	CurrentShortEventID++

	expandedNewEvent := GetShortCategory(newShortEvent)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    expandedNewEvent,
	})
}

func EditShortEvent(c *gin.Context) {
	var request EditShortEventReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	categoryExist := CheckCategoryExist(request.Category)

	founded, index, date := FindShortEvent(request.ID)

	if !categoryExist {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Category is not exist",
			"body":    nil,
		})
		return
	}

	if !founded {
		c.JSON(400, gin.H{
			"status":  3,
			"message": "Event is not exist",
			"body":    nil,
		})
		return
	}

	AllShortEvents[date][index].Category = request.Category
	AllShortEvents[date][index].Description = request.Description
	AllShortEvents[date][index].Time = request.Time
	AllShortEvents[date][index].Title = request.Title

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    AllShortEvents[date][index],
	})
}

func EditLongEvent(c *gin.Context) {
	var request EditLongEventReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	categoryExist := CheckCategoryExist(request.Category)

	founded, index, date := FindLongEvent(request.ID)

	if !categoryExist {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Category is not exist",
			"body":    nil,
		})
		return
	}

	if !founded {
		c.JSON(400, gin.H{
			"status":  3,
			"message": "Event is not exist",
			"body":    nil,
		})
		return
	}

	AllLongEvents[date][index].Category = request.Category
	AllLongEvents[date][index].Description = request.Description
	AllLongEvents[date][index].StartTime = request.StartTime
	AllLongEvents[date][index].EndTime = request.EndTime
	AllLongEvents[date][index].Title = request.Title

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    AllLongEvents[date][index],
	})
}

func CreateLongEvent(c *gin.Context) {
	var request NewLongEventReq
	err := c.Bind(&request)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  1,
			"message": "Incorrect data",
			"body":    nil,
		})
		return
	}

	categoryExist := CheckCategoryExist(request.Category)

	if !categoryExist {
		c.JSON(400, gin.H{
			"status":  2,
			"message": "Category is not exist",
			"body":    nil,
		})
		return
	}

	newLongEvent := LongEvent{
		ID:          CurrentShortEventID,
		Category:    request.Category,
		Title:       request.Title,
		Description: request.Description,
		StartTime:   request.StartTime,
		EndTime:     request.EndTime,
	}

	_, ok := AllLongEvents[request.StartTime]
	if ok {
		AllLongEvents[request.StartTime] = append(AllLongEvents[request.StartTime], newLongEvent)
	} else {
		AllLongEvents[request.StartTime] = []LongEvent{newLongEvent}
	}

	CurrentShortEventID++

	expandedNewEvent := GetLongCategory(newLongEvent)

	c.JSON(200, gin.H{
		"status":  0,
		"message": "Success",
		"body":    expandedNewEvent,
	})
}
