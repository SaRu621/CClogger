package logger

import (
    "os"
    "fmt"
    "time"
    "sort"
    "math"
    "bufio"
    "strconv"
    "strings"
 )

type Client struct {
    place []int          
    inClub bool
    timeStart []time.Time
    timeEnd []time.Time
}

func NewClient(Place []int, InClub bool, TimeStart, TimeEnd []time.Time) Client {
    return Client {
        place: Place,
        inClub: InClub,
        timeStart: TimeStart,
        timeEnd: TimeEnd,
    }
}

func formatDuration(duration time.Duration) string {
    hours   := int(duration.Hours())
    minutes := int(duration.Minutes()) % 60
 
    formattedDuration := fmt.Sprintf("%02d:%02d", hours, minutes)
    
    return formattedDuration
}

func findPlace(v []bool) int {
    for i := 0; i < len(v); i++ {
        if !v[i] {
            return i
        }
    }

    return -1
}

func timeRound(start, end time.Time) int {
    difference := end.Sub(start)
    return int(math.Ceil(difference.Hours()))
}

func earlierThan(time1, time2 time.Time) bool {
    if time1.Hour() < time2.Hour() {
        return true

    } else if time1.Hour() > time2.Hour() {
        return false

    } else {
        if time1.Minute() <= time2.Minute() {
            return true

        } else {
            return false

        }
    }
}

func moneyCounter(clients map[string]Client, cost , numOfPlaces int) ([]int, []time.Duration) {
    v := make([]int, numOfPlaces)
    t := make([]time.Duration, numOfPlaces)

    for _, val := range clients {      
        for i, _ := range val.place {
            t[val.place[i] - 1] += val.timeEnd[i].Sub(val.timeStart[i])
            v[val.place[i] - 1] += timeRound(val.timeStart[i], val.timeEnd[i]) * cost      
        }
    }

    return v, t
}

func Logger(pathIn, pathOut string) error {
    inFile, err := os.Open(pathIn)
    defer inFile.Close()  
    scanner := bufio.NewScanner(inFile)
    
    count := 0
    info := []string{}
    layout := "15:04"

    var numOfPlaces int
    var openingTime time.Time
    var closingTime time.Time
    var priceOfHour int
    forErrorString := ""

    outFile, err := os.Create(pathOut)
    defer outFile.Close()

    defer func() {
        if err := recover(); err != nil {
                fmt.Println(forErrorString)
                fmt.Fprintln(outFile, forErrorString)
        }
    }()

    for scanner.Scan() {
        if count == 0 {
            forErrorString = scanner.Text()
            info = append(info, forErrorString)
            numOfPlaces, err = strconv.Atoi(forErrorString)
        }

        if count == 1 {
            forErrorString = scanner.Text()
            workTime := strings.Split(forErrorString, " ")
            
            info = append(info, workTime[0])
            info = append(info, workTime[1])
        
            openingTime, err = time.Parse(layout, workTime[0]) 
            closingTime, err = time.Parse(layout, workTime[1])
        }

        if count == 2 {
            forErrorString = scanner.Text()
            info = append(info, forErrorString)
            priceOfHour, err = strconv.Atoi(forErrorString)
        }

        count++

        if count == 3 {
            break
        }
    }

    places := make([]bool, numOfPlaces)    
    allClients := make(map[string]Client)
    clientsQueue := make(chan string, numOfPlaces)

    forWritingString := ""

    forWritingString += info[1] + "\n"

    for scanner.Scan() {   

        forErrorString = scanner.Text()
        event := strings.Split(forErrorString, " ")
        forWritingString += strings.Join(event, " ") + "\n"
       
        switch event[1] {
            case "1":
                eventTime, _ := time.Parse(layout, event[0])

                if client, exists := allClients[event[2]]; exists && client.inClub {      
                    forWritingString += "YouShallNotPass\n"

                } else if earlierThan(openingTime, eventTime) && earlierThan(eventTime, closingTime) {
                     
                    if !exists {  
                        allClients[event[2]] = NewClient([]int{}, true, []time.Time{}, []time.Time{})  

                    } else {                        
                        client.inClub = true
                        allClients[event[2]] = client
                    
                    }
                } else {
                    forWritingString += event[0] + " 13 NotOpenYet\n"
                }

            case "2":
                if len(event) != 4 {
                    panic("")
                }

                if !allClients[event[2]].inClub {
                    forWritingString += event[0] + " 13 ClientUnknown\n"
                    break
                }

                if num, _ := strconv.Atoi(event[3]); places[num - 1] {
                    forWritingString += event[0] + " 13 PlaceIsBusy\n"
      
                } else {                                    
                    client := allClients[event[2]]
                    
                    if len(client.place) != 0 {
                        places[client.place[len(client.place) - 1] - 1] = false
                    }

                    client.place = append(client.place, num)
                    eventTime, _ := time.Parse(layout, event[0])
                    
                    if len(client.timeStart) != 0 {
                        client.timeEnd = append(client.timeEnd, eventTime)
                    }

                    client.timeStart = append(client.timeStart, eventTime)
                    allClients[event[2]] = client
                    places[num - 1] = true
                }

            case "3":
                if !allClients[event[2]].inClub {
                    forWritingString += event[0] + " 13 ClientUnknown\n"
                    break
                }

                if findPlace(places) != -1 {
                    forWritingString += event[0] + " 13 ICanWaitNoLonger!\n"
                    break

                } else {
                    select {
                        case clientsQueue <- event[2]: 
                
                        default:
                            client := allClients[event[2]]
                            client.inClub = false
                            allClients[event[2]] = client    
                            forWritingString += event[0] + " 11 " + event[2] + "\n"
                    }
                }

            case "4":  
                if allClients[event[2]].inClub {
                    client := allClients[event[2]]
                    eventTime, _ := time.Parse(layout, event[0])
                    client.timeEnd = append(client.timeEnd, eventTime)
                    client.inClub = false
                    allClients[event[2]] = client
                    
                    if len(client.place) != 0 {
                        places[client.place[len(client.place) - 1] - 1] = false
                    }

                    select {
                        case name := <- clientsQueue:
                            client = allClients[name]
                            freePlace := findPlace(places)
                            client.place = append(client.place, freePlace + 1)
                            client.timeStart = append(client.timeStart, eventTime)
                            allClients[name] = client
                            places[freePlace] = true      
                            forWritingString += event[0] + " 12 " + name + " " + strconv.Itoa(freePlace + 1) + "\n"
                        
                        default:

                    }
                } else {
                    forWritingString += event[0] + " 13 ClientUnknown\n"
                    break
                }
            default:
                panic("")
        }
    }

    names := []string{}
    for key, _ := range allClients {
        names = append(names, key)
    }

    sort.Strings(names)

    for _, name := range names {
        if allClients[name].inClub {
            client := allClients[name]
            client.timeEnd = append(client.timeEnd, closingTime)
            client.inClub = false
            allClients[name] = client
            forWritingString += info[2] + " 11 " + name + "\n"
        }
    }
    
    forWritingString += info[2]

    money, spendTime := moneyCounter(allClients, priceOfHour, numOfPlaces)
    
    fmt.Fprintln(outFile, forWritingString)

    for i, _ := range money {
        fmt.Fprintln(outFile, strconv.Itoa(i + 1), strconv.Itoa(money[i]), formatDuration(spendTime[i]))
    }

    return err
}
