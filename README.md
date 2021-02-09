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
	"github.com/Gasoid/workalendar/europe/russia"
)

func main() {
    now := time.Now()
    // check whether a day is holiday
    if russia.IsHoliday(now) {
        h, _ := russia.GetHoliday(now)
        fmt.Printf("Holiday is %s", h)
    } else {
        fmt.Print("No holiday today")
    }

    // get holiday name
    march8 := time.Date(2020, time.March, 8, 0, 0, 0, 0, time.UTC)
    h, err := russia.GetHoliday(march8)
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
