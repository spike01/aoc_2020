package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//byr (Birth Year)
//iyr (Issue Year)
//eyr (Expiration Year)
//hgt (Height)
//hcl (Hair Color)
//ecl (Eye Color)
//pid (Passport ID)
//cid (Country ID

var byrRegExp = regexp.MustCompile(`byr:(\d+)`)
var iyrRegExp = regexp.MustCompile(`iyr:(\d+)`)
var eyrRegExp = regexp.MustCompile(`eyr:(\d+)`)
var hgtRegExp = regexp.MustCompile(`hgt:(\d+)`)
var hclRegExp = regexp.MustCompile(`hcl:#?([[:alnum:]]+)`)
var eclRegExp = regexp.MustCompile(`ecl:#?(\w+)`)
var pidRegExp = regexp.MustCompile(`pid:#?([[:alnum:]]+)`)
var cidRegExp = regexp.MustCompile(`cid:(\d+)`)

type passport struct {
	birthYear      int
	issueYear      int
	expirationYear int
	height         int
	hairColor      string
	eyeColor       string
	passportId     string
	countryId      string
}

func newPassportFromBuf(buf []string) *passport {
	record := strings.Join(buf, " ")
	var birthYear int
	var issueYear int
	var expirationYear int
	var height int
	var hairColor string
	var eyeColor string
	var passportId string
	var countryId string

	var err error

	byr := byrRegExp.FindStringSubmatch(record)
	if len(byr) > 1 {
		birthYear, err = strconv.Atoi(byr[1])
		if err != nil {
			birthYear = 0
		}
	}

	iyr := iyrRegExp.FindStringSubmatch(record)
	if len(iyr) > 1 {
		issueYear, err = strconv.Atoi(iyr[1])
		if err != nil {
			issueYear = 0
		}
	}

	eyr := eyrRegExp.FindStringSubmatch(record)
	if len(eyr) > 1 {
		expirationYear, err = strconv.Atoi(eyr[1])
		if err != nil {
			expirationYear = 0
		}
	}

	hgt := hgtRegExp.FindStringSubmatch(record)
	if len(hgt) > 1 {
		height, err = strconv.Atoi(hgt[1])
		if err != nil {
			height = 0
		}
	}

	hcl := hclRegExp.FindStringSubmatch(record)
	if len(hcl) > 1 {
		hairColor = hcl[1]
	}

	ecl := eclRegExp.FindStringSubmatch(record)
	if len(ecl) > 1 {
		eyeColor = ecl[1]
	}

	pid := pidRegExp.FindStringSubmatch(record)
	if len(pid) > 1 {
		passportId = pid[1]
	}

	cid := cidRegExp.FindStringSubmatch(record)
	if len(cid) > 1 {
		countryId = cid[1]
	}

	return &passport{
		birthYear:      birthYear,
		issueYear:      issueYear,
		expirationYear: expirationYear,
		height:         height,
		hairColor:      hairColor,
		eyeColor:       eyeColor,
		passportId:     passportId,
		countryId:      countryId,
	}
}

func (p *passport) isValid() bool {
	return p.birthYear != 0 &&
		p.issueYear != 0 &&
		p.expirationYear != 0 &&
		p.height != 0 &&
		p.hairColor != "" &&
		p.eyeColor != "" &&
		p.passportId != ""
}

func (p *passport) String() string {
	return fmt.Sprintf("byr: %d iyr: %d eyr: %d hgt: %d hcl: %s ecl: %s pid: %s cid: %s valid: %v", p.birthYear, p.issueYear, p.expirationYear, p.height, p.hairColor, p.eyeColor, p.passportId, p.countryId, p.isValid())
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %s", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var buf []string
	var count int

	for sc.Scan() {
		next := sc.Text()
		buf = append(buf, next)
		if next == "" {
			p := newPassportFromBuf(buf)
			if p.isValid() {
				count++
			}
			buf = []string{}
		}
	}

	// Final element, buf isn't flushed on last entry
	p := newPassportFromBuf(buf)
	if p.isValid() {
		count++
	}
	buf = []string{}

	fmt.Println("Valid:", count)
	os.Exit(0)
}
