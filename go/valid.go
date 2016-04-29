/*
Valids rules and tags.

Usage:

/tmp/aa.txt file is like:

| 2542620639240 | {"tags": ["2199023256384","2199023264258","2199023264680","2199023264682","2199023264683"]}  | 2199023264683,0;2199023264682,0;2199023264680,0;2199023264258,2;2199023264240,3;2199023264238,2;2199023256385,1;2199023256384,2                                                                                 |
| 2542620639241 | {"tags": ["2199023262512", "2199023262446"]}                                                 | 2199023262512,0;2199023262446,2                                                                                                                                                                                 |
| 2542620639242 | {"tags": ["2199023264703", "2199023264704", "2199023260349", "2199023260350"]}               | 2199023264704,2;2199023264703,0;2199023260350,2;2199023260349,0                                                                                                                                                 |

go run valid.go --input /tmp/aa.txt
*/

package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
	"strings"
)

var inputFlag = flag.String("input", "", "The input text file.")

func Process(filepath string) []string {
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = fi.Close(); err != nil {
			panic(err)
		}
	}()

	lines := []string{}
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := strings.TrimSpace(strings.Trim(scanner.Text(), "|"))
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func ExtractRule(rule string) map[string]bool {
	reRule, _ := regexp.Compile(`\[(.*)\]`)
	result := reRule.FindStringSubmatch(rule)

	if len(result) != 2 {
		log.Fatalf("%v", result)
	}

	ret := map[string]bool{}
	segs := strings.Split(result[1], ",")
	for _, seg := range segs {
		ret[strings.Trim(strings.TrimSpace(seg), "\"")] = true
	}
	return ret
}

func ExtractTag(tag string) map[string]bool {
	ret := map[string]bool{}
	segs := strings.Split(strings.TrimSpace(tag), ";")
	for _, seg := range segs {
		parts := strings.Split(seg, ",")
		ret[parts[0]] = true
	}
	return ret
}

func M1InM2(m1 map[string]bool, m2 map[string]bool) bool {
	if len(m1) > len(m2) {
		log.Println("m1 is bigger than m2")
		return false
	}
	for k, _ := range m1 {
		if _, ok := m2[k]; !ok {
			log.Printf("can not find %s in m2\n", k)
			return false
		}
	}
	return true
}

func main() {
	flag.Parse()

	lines := Process(*inputFlag)
	for _, line := range lines {
		//		log.Printf("%s\n", line)
		segs := strings.Split(line, "|")
		rule := segs[1]
		tag := segs[2]
		ruleMap := ExtractRule(rule)
		tagMap := ExtractTag(tag)
		log.Printf("ruleMap = %v\n", ruleMap)
		log.Printf("tagMap = %v\n", tagMap)
		log.Println(M1InM2(ruleMap, tagMap))
	}
}
