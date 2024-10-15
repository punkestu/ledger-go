package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetsService struct {
	service       *sheets.Service
	spreadSheetId string
}

func NewSheetsService(ctx context.Context, client *http.Client) *SheetsService {
	sheetService, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		panic(err)
	}
	return &SheetsService{service: sheetService}
}

func (s *SheetsService) SetSpreadSheetId(spreadSheetId string) {
	s.spreadSheetId = spreadSheetId
}

func (s *SheetsService) GetCurrentBalance(wallet string) int64 {
	sheet, err := s.service.Spreadsheets.Values.Get(s.spreadSheetId, wallet).Do()
	if err != nil {
		panic(err)
	}
	if len(sheet.Values) <= 1 {
		return 0
	}
	currentBalance, err := strconv.ParseInt(sheet.Values[len(sheet.Values)-1][5].(string), 10, 64)
	if err != nil {
		panic(err)
	}
	return currentBalance
}

func (s *SheetsService) PushMutation(wallet string, kredit float64, debit float64, description string) {
	sheet, err := s.service.Spreadsheets.Values.Get(s.spreadSheetId, wallet).Do()
	if err != nil {
		panic(err)
	}
	var total float64
	if len(sheet.Values) > 1 {
		total, err = strconv.ParseFloat(sheet.Values[len(sheet.Values)-1][5].(string), 64)
		if err != nil {
			panic(err)
		}
	}
	id := uuid.New().String()
	total += kredit - debit
	newRow := []interface{}{id, time.Now().Format("2006-01-02 15:04:05"), kredit, debit, description, total}
	_, err = s.service.Spreadsheets.Values.Append(s.spreadSheetId, wallet, &sheets.ValueRange{Values: [][]interface{}{newRow}}).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		panic(err)
	}
}

func (s *SheetsService) GetTotalBalance() int64 {
	sheets, err := s.service.Spreadsheets.Get(s.spreadSheetId).Do()
	if err != nil {
		panic(err)
	}
	var total int64
	for _, sheet := range sheets.Sheets {
		totalInSheet := s.GetCurrentBalance(sheet.Properties.Title)
		total += totalInSheet
	}
	return total
}
