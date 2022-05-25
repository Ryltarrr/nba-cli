package displayer

import (
	"github.com/charmbracelet/lipgloss"
)

var white = "#fff"
var black = "#000"

var teamColors = map[string]lipgloss.Style{
	"ATL": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#e03a3e")).Bold(true).Padding(0, 1),
	"BOS": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#007A33")).Bold(true).Padding(0, 1),
	"BKN": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color(black)).Bold(true).Padding(0, 1),
	"CHA": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#1d1160")).Bold(true).Padding(0, 1),
	"CHI": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#CE1141")).Bold(true).Padding(0, 1),
	"CLE": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#860038")).Bold(true).Padding(0, 1),
	"DAL": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#00538C")).Bold(true).Padding(0, 1),
	"DEN": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#0E2240")).Bold(true).Padding(0, 1),
	"DET": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#C8102E")).Bold(true).Padding(0, 1),
	"GSW": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#1D428A")).Bold(true).Padding(0, 1),
	"HOU": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#CE1141")).Bold(true).Padding(0, 1),
	"IND": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#002D62")).Bold(true).Padding(0, 1),
	"LAC": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#c8102E")).Bold(true).Padding(0, 1),
	"LAL": lipgloss.NewStyle().Foreground(lipgloss.Color("#FDB927")).Background(lipgloss.Color("#552583")).Bold(true).Padding(0, 1),
	"MEM": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#5D76A9")).Bold(true).Padding(0, 1),
	"MIA": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#98002E")).Bold(true).Padding(0, 1),
	"MIL": lipgloss.NewStyle().Foreground(lipgloss.Color("#EEE1C6")).Background(lipgloss.Color("#00471B")).Bold(true).Padding(0, 1),
	"MIN": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#0C2340")).Bold(true).Padding(0, 1),
	"NOP": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#0C2340")).Bold(true).Padding(0, 1),
	"NYK": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#006BB6")).Bold(true).Padding(0, 1),
	"OKC": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#007ac1")).Bold(true).Padding(0, 1),
	"ORL": lipgloss.NewStyle().Foreground(lipgloss.Color("#C4ced4")).Background(lipgloss.Color("#0077c0")).Bold(true).Padding(0, 1),
	"PHI": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#006bb6")).Bold(true).Padding(0, 1),
	"PHX": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#1d1160")).Bold(true).Padding(0, 1),
	"POR": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#E03A3E")).Bold(true).Padding(0, 1),
	"SAC": lipgloss.NewStyle().Foreground(lipgloss.Color("#63727A")).Background(lipgloss.Color("#5a2d81")).Bold(true).Padding(0, 1),
	"SAS": lipgloss.NewStyle().Foreground(lipgloss.Color("#c4ced4")).Background(lipgloss.Color(black)).Bold(true).Padding(0, 1),
	"TOR": lipgloss.NewStyle().Foreground(lipgloss.Color("#A1A1A4")).Background(lipgloss.Color("#ce1141")).Bold(true).Padding(0, 1),
	"UTA": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#002B5C")).Bold(true).Padding(0, 1),
	"WAS": lipgloss.NewStyle().Foreground(lipgloss.Color(white)).Background(lipgloss.Color("#002B5C")).Bold(true).Padding(0, 1),
}
