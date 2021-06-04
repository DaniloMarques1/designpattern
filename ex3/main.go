package main

import "fmt"

func main() {
        shiftReport := ShiftReport{}
        dateReport := DateReport{}

        report := Report{report: shiftReport}
        report.GenerateReport()

        fmt.Println()

        report = Report{report: dateReport}
        report.GenerateReport()
}

func test(report iReport) {
        fmt.Println("Opa")
}
