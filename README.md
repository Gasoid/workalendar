[![Go Reference](https://pkg.go.dev/badge/github.com/Gasoid/workalendar.svg)](https://pkg.go.dev/github.com/Gasoid/workalendar)
# Overview
This is currently a work-in-progress.
I've used the repo: https://github.com/peopledoc/workalendar as source of holidays.

Workcalendar is a go library that is inteded to handle holidays.

# Status
This library is currently being developed.

If you spot any bug or wish to add a calendar, please fill Pull Request.

## Usage Examples
### Example 1
```go
package main

import (
    "fmt"
	"time"
    //import a country
	"github.com/Gasoid/workalendar/europe/germany"
)

func main() {
    now := time.Now()
    // check whether a day is holiday
    if germany.IsHoliday(now) {
        h, _ := germany.GetHoliday(now)
        fmt.Printf("Holiday is %s", h)
    } else {
        h, _ := germany.NextHoliday(now)
        fmt.Print("No holiday today")
        fmt.Printf("Next holiday is %s", h)
    }

    // get holiday name
    may1 := time.Date(2021, time.May, 1, 0, 0, 0, 0, time.UTC)
    h, err := germany.GetHoliday(march8)
    if err != nil {
        fmt.Print("Hm, it is weird")
    } else {
        fmt.Printf("Holiday is %s", h)
    }
}

```

### Example 2
```go
package main

import (
    "fmt"
	"time"
    //import a region
	"github.com/Gasoid/workalendar/europe/spain/catalonia"
)

func main() {
    now := time.Now()
    // check whether a day is holiday
    if catalonia.IsHoliday(now) {
        h, _ := catalonia.GetHoliday(now)
        fmt.Printf("Holiday is %s", h)
    } else {
        fmt.Print("No holiday today")
    }

    // get the holiday name
    march8 := time.Date(2020, time.March, 8, 0, 0, 0, 0, time.UTC)
    h, err := catalonia.GetHoliday(march8)
    if err != nil {
        fmt.Print("Hm, it is weird")
    } else {
        fmt.Printf("Holiday is %s", h)
    }
}
```
## Available Calendars
### Europe:
- Austria
- Belarus
- Belgium
- Bulgaria
- Croatia
- Cyprus
- Czech Republic
- Denmark
- Estonia
- Finland
- France
- Germany
- Greece
- Hungary
- Iceland
- Ireland
- Italy
- Latvia
- Lithuania
- Luxembourg
- Malta
- Monaco
- Netherlands
- Norway
- Poland
- Portugal
- Romania
- Russia
- Serbia
- Slovakia
- Slovenia
- Spain
- Sweden
- Switzerland
- Ukraine
- United Kingdom


## License
This library is published under the terms of the MIT License. Please check the LICENSE file for more details.
