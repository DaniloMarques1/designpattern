package main

import "fmt"

type iReport interface {
        BuildHeader()
        BuildBody()
        BuildFooter()
}

type Report struct {
        report iReport
}

func (r Report) GenerateReport() {
        r.report.BuildHeader()
        r.report.BuildBody()
        r.report.BuildFooter()
}

type ShiftReport struct {
}

func (sr ShiftReport) BuildHeader() {
        fmt.Println("Construindo header do shift report")
}

func (sr ShiftReport) BuildBody() {
        fmt.Println("Construindo o body do shift report")
}

func (sr ShiftReport) BuildFooter() {
        fmt.Println("Construindo o footer do shift report")
}

type DateReport struct {
}

func (dr DateReport) BuildHeader() {
        fmt.Println("Construindo header do date report")
}

func (dr DateReport) BuildBody() {
        fmt.Println("Construindo o body do date report")
}

func (dr DateReport) BuildFooter() {
}
