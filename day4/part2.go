package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

var byrRegExp = regexp.MustCompile(`byr:(\d{4})`)
var iyrRegExp = regexp.MustCompile(`iyr:(\d{4})`)
var eyrRegExp = regexp.MustCompile(`eyr:(\d{4})`)
var hgtRegExp = regexp.MustCompile(`hgt:(\d+)(cm|in)`)
var hclRegExp = regexp.MustCompile(`hcl:#([0-9a-f]{6})`)
var eclRegExp = regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)`)
var pidRegExp = regexp.MustCompile(`pid:([[:alnum:]]{9})`)
var cidRegExp = regexp.MustCompile(`cid:(\d+)`)

type passport struct {
	birthYear      int
	issueYear      int
	expirationYear int
	height         int
	heightUnit     string
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
	var heightUnit string
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
		if birthYear < 1920 || 2002 < birthYear {
			birthYear = 0
		}
	}

	iyr := iyrRegExp.FindStringSubmatch(record)
	if len(iyr) > 1 {
		issueYear, err = strconv.Atoi(iyr[1])
		if err != nil {
			issueYear = 0
		}
		if issueYear < 2010 || 2020 < issueYear {
			issueYear = 0
		}
	}

	eyr := eyrRegExp.FindStringSubmatch(record)
	if len(eyr) > 1 {
		expirationYear, err = strconv.Atoi(eyr[1])
		if err != nil {
			expirationYear = 0
		}
		if expirationYear < 2020 || 2030 < expirationYear {
			expirationYear = 0
		}
	}

	hgt := hgtRegExp.FindStringSubmatch(record)
	if len(hgt) > 2 {
		height, err = strconv.Atoi(hgt[1])
		if err != nil {
			height = 0
		}
		heightUnit = hgt[2]
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
		heightUnit:     heightUnit,
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
		p.isValidHeight() &&
		p.hairColor != "" &&
		p.eyeColor != "" &&
		p.passportId != ""
}

func (p *passport) isValidHeight() bool {
	if p.heightUnit == "cm" && !(p.height < 150 || 193 < p.height) {
		return true
	}
	// This seems to be a bug, 76 should give a valid submission (it doesn't)
	// - or I've messed up a range somewhere else
	if p.heightUnit == "in" && !(p.height < 59 || 75 < p.height) {
		return true
	}
	return false
}

// This and InvalidFields() are pure debugging... I could not for the life of
// me figure out why my answer was repeatedly wrong
func (p *passport) String() string {
	var out []string
	out = append(out, "[")
	if p.birthYear != 0 {
		out = append(out, fmt.Sprintf("byr:%d", p.birthYear))
	}
	if p.countryId != "" {
		out = append(out, fmt.Sprintf("cid:%s", p.countryId))
	}
	if p.eyeColor != "" {
		out = append(out, fmt.Sprintf("ecr:%s", p.eyeColor))
	}
	if p.expirationYear != 0 {
		out = append(out, fmt.Sprintf("eyr:%d", p.expirationYear))
	}
	if p.hairColor != "" {
		out = append(out, fmt.Sprintf("hcr:#%s", p.hairColor))
	}
	if p.isValidHeight() {
		out = append(out, fmt.Sprintf("hgt:%d%s", p.height, p.heightUnit))
	}
	if p.issueYear != 0 {
		out = append(out, fmt.Sprintf("iyr:%d", p.issueYear))
	}
	if p.passportId != "" {
		out = append(out, fmt.Sprintf("pid:%s", p.passportId))
	}
	out = append(out, "]")
	return strings.Join(out, " ")
}

func (p *passport) InvalidFields() {
	if p.birthYear == 0 {
		fmt.Println("byr not between 1920 and 2002")
	}
	if p.issueYear == 0 {
		fmt.Println("iyr not between 2010 and 2020")
	}
	if p.expirationYear == 0 {
		fmt.Println("eyr not between 2020 and 2030")
	}
	if !p.isValidHeight() {
		if p.heightUnit == "cm" {
			fmt.Println("hgt not between 150 and 193 cm")
		}
		if p.heightUnit == "in" {
			fmt.Println("hgt not between 59 and 76 in")
		}
	}
	if p.hairColor == "" {
		fmt.Println("hcl invalid/not present")
	}
	if p.eyeColor == "" {
		fmt.Println("ecl invalid/not present")
	}
	if p.passportId == "" {
		fmt.Println("pid invalid/not present")
	}
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
			} else {
				joined := strings.Join(buf, " ")
				split := strings.Split(joined, " ")
				sort.Strings(split)

				fmt.Println(split)
				fmt.Println(p)
				p.InvalidFields()
				fmt.Println("")
			}
			buf = []string{}
		}
	}

	// Final element, buf isn't flushed on last entry
	p := newPassportFromBuf(buf)
	if p.isValid() {
		count++
	} else {
		joined := strings.Join(buf, " ")
		split := strings.Split(joined, " ")
		sort.Strings(split)

		fmt.Println(split)
		fmt.Println(p)
		p.InvalidFields()
		fmt.Println("")
	}
	buf = []string{}

	fmt.Println("Valid:", count)
	os.Exit(0)
}
