package getEventsJson

// type shifts struct {
// 	Ltas []struct {
// 		WorkDate      string      `json:"workDate"`
// 		StartTime     string      `json:"startTime"`
// 		EndTime       string      `json:"endTime"`
// 		WorkHours     float64     `json:"workHours"`
// 		ShiftSumID    int         `json:"shiftSumId"`
// 		TimeOffType   string      `json:"timeOffType"`
// 		TimeCodeColor interface{} `json:"timeCodeColor"`
// 		FullDay       bool        `json:"fullDay"`
// 	} `json:"ltas"`
// 	Shifts []struct {
// 		ShiftSource           string        `json:"shiftSource"`
// 		WorkDate              string        `json:"workDate"`
// 		StartTime             string        `json:"startTime"`
// 		EndTime               string        `json:"endTime"`
// 		ShiftSumID            int           `json:"shiftSumId"`
// 		ShiftLabel            string        `json:"shiftLabel"`
// 		ReqSkillIds           []interface{} `json:"reqSkillIds"`
// 		WorkHours             float64       `json:"workHours"`
// 		BillboardPosting      interface{}   `json:"billboardPosting"`
// 		PendingShiftTradeType string        `json:"pendingShiftTradeType"`
// 		PreValidationErrors   struct{}      `json:"preValidationErrors"`
// 		Modified              bool          `json:"modified"`
// 	} `json:"shifts"`
// }

type shifts struct {
	Ltas []struct {
		WorkDate      string      `json:"workDate"`
		StartTime     string      `json:"startTime"`
		EndTime       string      `json:"endTime"`
		WorkHours     float64     `json:"workHours"`
		ShiftSumID    int         `json:"shiftSumId"`
		TimeOffType   string      `json:"timeOffType"`
		TimeCodeColor interface{} `json:"timeCodeColor"`
		FullDay       bool        `json:"fullDay"`
	} `json:"ltas"`
	Shifts []struct {
		ShiftSource           string        `json:"shiftSource"`
		WorkDate              string        `json:"workDate"`
		StartTime             string        `json:"startTime"`
		EndTime               string        `json:"endTime"`
		ShiftSumID            int           `json:"shiftSumId"`
		ShiftLabel            string        `json:"shiftLabel"`
		ReqSkillIds           []interface{} `json:"reqSkillIds"`
		WorkHours             float64       `json:"workHours"`
		BillboardPosting      interface{}   `json:"billboardPosting"`
		PendingShiftTradeType string        `json:"pendingShiftTradeType"`
		PreValidationErrors   struct{}      `json:"preValidationErrors"`
		Modified              bool          `json:"modified"`
		Details               []struct {
			Department struct {
				ID      int    `json:"id"`
				LocName string `json:"locName"`
				LocDesc string `json:"locDesc"`
				Name    string `json:"name"`
			} `json:"department"`
		}
	} `json:"shifts"`
}
