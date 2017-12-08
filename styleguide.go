package main

import (
  "fmt"
  "encoding/xml"
  "os"
  "io/ioutil"
  "strings"
  "strconv"
)

type (

  Stat struct {
    Low      string `xml:"low"`
    High     string `xml:"high"`
    Flexible bool `xml:"flexible,attr"`
  }
  Stats struct {
    Ibu   *Stat `xml:"ibu"`
    Og    *Stat `xml:"og"`
    Fg    *Stat `xml:"fg"`
    Srm   *Stat `xml:"srm"`
    Abv   *Stat `xml:"abv"`
  }
  Subcategory struct {
    ID          string `xml:"id,attr"`
    Name        string `xml:"name"`
    Aroma       string `xml:"aroma"`
    Appearance  string `xml:"appearance"`
    Flavor      string `xml:"flavor"`
    Mouthfeel   string `xml:"mouthfeel"`
    Impression  string `xml:"impression"`
    Comments    string `xml:"comments"`
    History     string `xml:"history"`
    Ingredients string `xml:"ingredients"`
    Comparison  string `xml:"comparison"`
    Examples    string `xml:"examples"`
    Tags        string `xml:"tags"`
    Stats       *Stats `xml:"stats"`
  }

  Revision struct {
    Number string `xml:"number,attr"`
    Text   string `xml:",chardata"`
  }

  Category struct {
    ID string `xml:"id,attr"`
    Revision *Revision `xml:"revision"`
    Name        string `xml:"name"`
    Notes       string `xml:"notes"`
    Subcategories []*Subcategory `xml:"subcategory"`
  }

  Class struct {
    Type         string `xml:"type,attr"`
    Categories []*Category `xml:"category"`
  }

  Styleguide struct {
    Classes []*Class `xml:"class"`
  }

)

var INSERT_CATEGORY = "insert into category (category_name, category_number) values ('%s', '%s');"
var INSERT_STYLE = "insert into style (style_name, style_letter, category_name) values ('%s', '%s', '%s');"
var INSERT_TAG_TYPE = "insert into tag_type (tag_type_name) values ('%s');"
var INSERT_TAG = "insert into tag (tag_name, tag_type, tag_description) values ('%s', '%s', '%s');"
var INSERT_STYLE_TAG = "insert into style_tag (tag_name, style_name) values ('%s', '%s');"
var INSERT_STAT_TYPE = "insert into stat_type (stat_type_name, measuring_unit) values ('%s', '%s');"
var INSERT_STATISTIC = "insert into vital_statistic (style, statistic_type, statistic_lower, statistic_upper) values ('%s', '%s', %.3f, %.3f);"
var INSERT_STATISTIC_FLEXIBLE = "insert into vital_statistic (style, statistic_type) values ('%s', '%s');"

func main() {
	xmlFile, err := os.Open("styleguide.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
  bytes, err := ioutil.ReadAll(xmlFile)
  if err != nil {
    fmt.Println("Unable to read file: ", err)
    return
  }
	var styleguide Styleguide

	xml.Unmarshal(bytes, &styleguide)

  loadTagTypes()
  loadStatisticTypes()

	for _, category := range styleguide.Classes[0].Categories {
    fmt.Printf(INSERT_CATEGORY, category.Name, category.ID)
    fmt.Print("\n\n")
    for _, style := range category.Subcategories {
      fmt.Printf(INSERT_STYLE, style.Name, style.ID, category.Name)
      fmt.Printf("\n")
      tags := strings.Split(style.Tags, ", ")
      for _, tag := range tags {
        fmt.Printf(INSERT_STYLE_TAG, tag, style.Name)
        fmt.Print("\n")
      }
      loadStat(style.Stats.Abv, style.Name, "abv")
      loadStat(style.Stats.Fg, style.Name, "fg")
      loadStat(style.Stats.Og, style.Name, "og")
      loadStat(style.Stats.Ibu, style.Name, "ibu")
      loadStat(style.Stats.Srm, style.Name, "srm")
      fmt.Print("\n")
    }
  }
}

func loadStat(stat *Stat, style, statType string) {
  if stat.Flexible {
    fmt.Printf(INSERT_STATISTIC_FLEXIBLE, style, statType)
    fmt.Print("\n")
    return
  }

  low, err := strconv.ParseFloat(stat.Low, 64)
  if err != nil {
    panic(err.Error())
  }
  high, err := strconv.ParseFloat(stat.High, 64)
  if err != nil {
    panic(err.Error())
  }
  fmt.Printf(INSERT_STATISTIC, style, statType, low, high)
  fmt.Print("\n")
}

func loadTagTypes() {
  tagTypes := map[string][]string {
    "Strength": { "session-strength", "standard-strength", "high-strength", "very-high-strength"},
    "Color": {"pale-color", "amber-color", "dark-color"},
    "Fermentation": {"top-fermented", "bottom-fermented", "any-fermented", "wild-fermented"},
    "Conditioning": {"lagered", "aged"},
    "Region of Origin": {"british-isles", "western-europe", "central-europe", "eastern-europe", "north-america", "pacific"},
    "Style Family": {"ipa-family", "brown-ale-family", "pale-ale-family", "pale-lager-family", "pilsner-family", "amber-ale-family", "amber-lager-family", "dark-lager-family", "porter-family", "stout-family", "bock-family", "strong-ale-family", "wheat-beer-family", "specialty-family"} ,
    "Era": {"craft-style", "traditional-style", "historical-style"},
    "Dominant Flavor": {"malty", "bitter", "balanced", "hoppy", "roasty", "sweet", "smoke", "sour", "wood", "fruit", "spice"},
  }

  for tagType, tags := range tagTypes {
    fmt.Printf(INSERT_TAG_TYPE, tagType)
    fmt.Printf("\n")
    for _, tagName := range tags {
      fmt.Printf(INSERT_TAG, tagName, tagType, "")
      fmt.Print("\n")
    }
    fmt.Print("\n\n")
  }
}

func loadStatisticTypes() {
  types := map[string]string {
    "ibu": "international bitter units",
    "og": "original gravity",
    "fg": "final gravity",
    "srm": "color",
    "abv": "alcohol by volume",
  }
  for name, description := range types {
    fmt.Printf(INSERT_STAT_TYPE, name, description)
    fmt.Print("\n")
  }
  fmt.Print("\n\n")
}
