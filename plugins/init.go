package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.2" )

func o7iN94WkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2TDtmtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mg7GzUH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x29gLqDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sY25eNMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xhvxp2rLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgY4ecXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRpPxOXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPmqd4e7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n18FjbIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0evarYckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUgSqphiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8hoI9DZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnjoePpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBYNG8beWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func icfcB5CrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PuPS7ESvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YwQaNheVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddud6hrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYTCxmd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63mSt8DeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EVUfnLoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmPrlg2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g21DMM1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f7lFt9dHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMA1o7gGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDjaQdeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDPiDufwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CW97Lck9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KDrKGuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23bH0xrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LneUnu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMtJL9QfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SoM7KhRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRFzrvaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F90btwNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8Io6ZafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ov3lcFUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86cgaGLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ytp0IrO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDkR17BHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjyTEd3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zPNtrj87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUmOubJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAsR9PjRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUWyHrJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bg97IEMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3jTrucJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3C72SiPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbdLan2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqxvgaW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmUfnCPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65IN9RCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WeozIguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQqSnNL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJqOeueWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xBAZcYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJqXGHCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UGmp0CcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RyiCa7HKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DP6KkqOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6Y8WRc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TqpF6j6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tk1X3m5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Of8Pbsy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLEsaITJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XA2IhpiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OE3O5xQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0K2f7mmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n082en7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uk05O67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsrOrptWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uq4Fe5FhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R04ui3NYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msPn5skAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRKPMcrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dF1k4d4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NQolismyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfFCYtxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQxFh1cNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKdEyj79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3L1pyXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5uFNEyPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EcqDuWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUxbC1MaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHtxCq2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5LyvNHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmK97Y7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sL9jt6K8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z08vpwToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSLlYm4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QT6ioahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfwHbtSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qwMkP2m0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXu8O5iRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3e2SGSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXD8chneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68o72mmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1iMfIw45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g95PoGMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFxh4XaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 407XaT6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8BP8Up2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcNRtIYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDE8UGR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNdO1TTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFe5rZHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XS0rET7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbEgEKtVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NN7miRX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgkrvnC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J832z10SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4RVfyfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kY41PVC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4BFUD1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yIFhUKBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQrYn5swWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpd1RZfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQnK961sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D90cLJEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1EUX0aUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gd9ylN57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QpwdLTKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Swlv8ojkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUuUp9ASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2rzS9zmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCkcrUrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CcAt8hFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSnanz88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7hIEsk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 01i5W1kBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NXPRq67pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sp4xdEv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohkDFAdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NThJRDq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gj9thFymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHbXjVfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WghsvQHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvzwg7c8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bC3U0nxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func getaDAB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLGjqblOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lUS5TLy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IOHN3w8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dUAPJReYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func En779ZOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCx2sxVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afzwjKzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8xZTPQwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hmm7jqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cyt2geFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gphBM1NzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXbQs4TNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOr85EC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcCr3scfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8TBc1DnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rshg0TDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbcZFb3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVmj1VD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dydK94IoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8xYUI8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLlwD0a5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hiSHtV44Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbH5roySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZmoYDcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLOQJXPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WTDrC8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8Ygkq7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diQHvQwZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahpRiBsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3IfadWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QwPLfX1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3PodRpM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8GtwlXM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrVA6jYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XalnnE2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYNqEr57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ic1krjJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSmjcq2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20DITnrzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdRnpeWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lt8rEoUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzQhfhebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jyOa2DwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEBHUx4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QU9KiNtLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IVUdHRhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZvlUt0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5X6eOGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Ay2Mm85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xt0J4smAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJFlwV46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cM4usemVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgQEYVU0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NirXMSn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vclseVAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSRobkRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmHT1sI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzppQfG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 17B2YbD6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4O4qj8XlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sl0rmhrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lk094sEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHdG42IVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func judfkhoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pp5I0MCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ol6r9cVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x0DaAyxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KQqCPQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpF1vvAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvLwVsUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aIVce6utWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHl3sO2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a44boWBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vvPXKoQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxXpNyKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrkKhvNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TsqXOuHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SabPMW53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nAXYBGi7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hoRIfC0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5D1Q5ijPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rocDaWjMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrJVht4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xE46uiSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QK3c2hAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DuajUIPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXR3ThC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAS84PzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6UoznEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HF7XW2J1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PU7IMeRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 49v2PieoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UDgywHFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCd0wpzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpkUWMuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nuOR6OI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hda8gkkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMUt0AAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsiWQiwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3u3ztzIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35DQivw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tOnzsS0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQJ9o82NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAtbQLlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1wRUsGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blbTYALLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONW8XkqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmdZEPERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cqnr94JHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N92kK53XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAWp4s40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a3Cx8mfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 11cSiUWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKzinMAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93OBDPLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VZ9Sb9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gjoq6a36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3IwJjUO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yc4jpEX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTUIisL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhjXKGJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func doTwWEYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TrClnl9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYz2fklqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iizNYNheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulmKtUKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UytrZPAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ss0SHmqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I6IpW1T6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRxQt26mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7D7kjW69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptL0t5r2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLBBVfVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4PysMZ1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8vugoVJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsbCwgT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uP2MZK1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHEgqJrlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taEPaBFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwLvYlk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func spcLBborWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z1O52JS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSqI5B1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvqKzVR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFoV5qjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjtkCSS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3QT2bwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8sdrTa8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7G61IvdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ajdf9254Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5I5c4UsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgOxV4TKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6tgzAUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZSqGvKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ehfJNhWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WimujOeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNORWzf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfEuxmc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fT1gUsmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03wzwuurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izxzAd0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DB9w3qLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGJookmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bd5I0fP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucBgvVLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func veuIsAMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GF359eH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p06Qfse5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3CJRGcChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Hu5etqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfycMfYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEvfEojMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1t8BHYfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSQhaNIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TK1idL39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJWYUpvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TaISLSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7rEfMyobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aIlfXKVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O84HKNGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYxDAz71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NysReqokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B52BRPjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6y9twzHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBKfynNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDdqiLjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zbj1PattWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcCXNv8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjI35wH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxoxLsXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cF1IGpJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cU5uVqaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6CDfDEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXXC6WF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEBoTglZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQ44GvjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFwn2gwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aB3O43onWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ELipXx3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhzrNS5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKXzAXzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Y98dH62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7YNPqqCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNfxziyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkMsW7KgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wuXsYdUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyo6hxOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItetimsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pnbVUsxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3ujgihRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3og0by7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 24DRPJbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODwZ7brZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PtrmbbzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iztYyZZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUJMBmcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFmDCZELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUcDagGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SjyjhLr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxsIPzB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKZfa8oeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzXC4MYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWdTp6t5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otqdNCl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XBWfVcjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0iTNKCGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7iNCROa2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlwM0lR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNJNs1UcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nkpm41YdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AkVFWjmZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zj1jg3IzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROhdugAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cv7sDxBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Audk3hiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vxw2uNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4OTPAtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMCb1xczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CeFh14yoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLKAenTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwJckw91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmGEfXWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLVjL39iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXcUH5wfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PylbbLHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQx2ITPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDkHKd4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlOcRXGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUArbZWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pS1AxFs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1L9139xLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Jv3QOZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhPGk9Z5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8yFGmTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNEqjlsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yp0tmcZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wCUVWOvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPLq0ehKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZSuY9wiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJyZcYTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cyBfrafRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAK9ujmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJM4XNJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JW4WmEewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YYnSWyyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOmHUom6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVBzcKNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgxzVVfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jLRS4AyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsM4LUn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xE6jhKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HeZtl0qyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ll521XrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWqOu5DmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxnFqqGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqWZ9Fx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZXHALXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXyBcS3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVR2dVKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADLtIuHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASmRmPgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDrgLOi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqnbY1mDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNXmhJI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFd0uD1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ujmIRqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixOUW2AEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4PZ8TwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgMjdOD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZPTI4brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fN6f2IRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxEMxvJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2Mr4GYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vAry41QpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2OqU1jghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HSZI0pQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 61WUv7JuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gtw6TlQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3gbJYMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkgXw3Q8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJkpdDiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPeRTRspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZLrm8LAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FR2HDzYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSZscLecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVXVdc3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUBS6B9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5QlpiYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHxeGMKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAKSkX1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uStPbor1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gblspk7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7e8T9aSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYigCVBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxxzcY4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhMimq77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPD4RxgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jKJzkdy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AD4X98OFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 27zwECxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64LVbo25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MeVaiGNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhHyZS5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VeQyqze6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5BgWwDn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZhrqlfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcjXv8KdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uJYzSfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpOiso26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8TSHZHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XR0C4mvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXw2QbYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSNI6tyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3PZl51DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tj9ezILvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEOROlKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wg8kynEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZI6IaokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yM6lSgcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CnTe0I9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAZh90nNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTUI1VbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ccLCegKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cVVm08dwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPqG96VFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gndjmQhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func syKHUTnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4kg4gtGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTPWkpWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9lDsHO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aRk6UIIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHKUiBJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SoeAHkbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLiKN0ddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aj14mB8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClxkqgNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uc1WuTTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvXxgvr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSFNKKwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJOJHnSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1VErSS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pK1ZVOsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSTINCXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6J9UKumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBtUt246Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PmbPrTp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elClnPAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGqktWA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZXJ0Q5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqSIoMopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vsfC5E0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aO41zoz3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6cOlV3OCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dqfYukVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHFNYPmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func loCVFcrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Degs5uxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K40VwMDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBC7IIBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SCuviBceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jc4doVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEUkuxDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nSvzUOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rEE4SPLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZdZ82K1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func biL2sw8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6WZw2EdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5iDNSBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8io4BGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4KLCDc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4hoqWmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujbnclXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYmUMWTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwYnYk69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3gFWp1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqCHPYfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWyEv0gQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIboJnB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpxuYf9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvxigwouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXt204hqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u25mTGoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KL9hYtvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeBn9tHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBSLlKzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Fzxb3ddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XhFshyR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q6eEZXKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8m7xGxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ziZMGIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8L6PnQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXVTG9cYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VIguABg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pr3c454TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nuudN73QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXL8IBV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJk3xPXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5osr1tGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xidLWz8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GZRKWWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8z75G7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFXWLMpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6sGeDE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W02FmG66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XWe4DFn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mMECR6XxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7PGcAQ33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 97NDaGInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxNxQdmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkt1Dre8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCxCbhAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AMDToh3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 821vDyyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LR5k5u3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojGT7skPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ck4XoPNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OzEdbRsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xkbtS8KPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9cNjUcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qG0iex3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxmJlYG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xu3VZKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 58mrGPZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJc1ksdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPwrW3g0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0jom17GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6InPrwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g002gMCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JlvZ9AgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func re6Yq3AOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gZXzYJKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tv3Y8tRnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9yDmJldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIj3AgZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWs10YUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RmUpgBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGMpFg68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n06Ow4yEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkTzVGaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMaihBAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqiR3LLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vc57st0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONhlShWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNomDTMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3v2x7ZXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqOaw8ZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nC8dI0YsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qy0a5WxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nsAI4AGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8q6V5tHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fsrhbg4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8J1Tk6hKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njalGF3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpwDuVt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APdnZMbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fx5bah6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juKvsfknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4TViBTLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhWLPKidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljJpvCGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LK5fMzDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRCsnfaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teZ9hEO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjT7MgLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S8mhG5zJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16o2N7VfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1VUKKlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsK1acT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4jAB17RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XbwYtt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CYpJsW9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hO8DvsYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axKN0B7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xm0AFk7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sL54If8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OiXJy4a9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eMBUhFURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CltaYOwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OlrlXgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HYvVbNsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aidAp4mDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9xmpC0QPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMNPs6OHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJTYrXoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9XdVAQwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzGwV0VlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fdpekpw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XviMz5lUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgjYFegyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0JbkL0WEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9HPiJ2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzG37j3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dKETO6B0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6rCywAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTCx6TXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XuC3qAOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbPnnGFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AMdxCCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hhLKm6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTyJN52zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJUCxc43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wTu7iapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yvtMxMQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9oujXHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRwSTWIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T2D47OfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ni2pcmmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1IcGOn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Io1otvK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsV7b8LkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9hy0VDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZlekXY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vcvqk50IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0gBgPnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qnR0N8zMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZK1Ew1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUCZzWmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JuE5JFDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TbSCp9yhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bRkfWwjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9H1Qyos7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nAsyh5GsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RwZqsUTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRr8m5gSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZOhx4fBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEsWxulTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GjpvNJBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p6HUY9ikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VB4IAgHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZQ3wUJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSkYJGzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GiIVMoB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tC31FDzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54b5F5fOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxKWoxFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LxVE8jpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mO0tcJYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Joe6xYeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5AWGriFBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JH9JByxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGDRANmcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxMBJ8kZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x62UtNC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DzMjcYDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GY96oq7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J0dfI9s2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OeGvSkwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8J4PInPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0a6IAavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sAbGNltbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wm2rH3Y2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EjJgzQplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkLPgxfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chgjs1zgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXdmvtuAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALBgCxjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btEVwptpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ja8L8Oi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEz4MBiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rl5P4vr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nY4McskZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bxmshDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SlMeJlX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FMc7riB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfeA1jESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OM6kZ9BoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPm5RjQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3TKL4ucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFFW4xw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIyDHnNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gsrf231Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QmGhPgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func asoJkXnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20v957Q0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIIVapH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nuwOWrR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RZFdQpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gR8WiJiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdAeWRShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4JQ8fe4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLJ7u58YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrMdsCybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KdN0f2dwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fyva5oseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bz776TJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpKwGbDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1VCOEpRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNombgAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jmMH0rBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z9t9VQbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQkJwSY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9UIkazfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWklFFWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0rKgFDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orqfZBc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTnLjjQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdz1yZ5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4I1wEKYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYtem7fYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xs8w9urcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VArYQHZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPjvBMPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3SBEggb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4RbzSO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERvdvYQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkZwwBEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QcqXpX8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYGhhALAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ws3TQkwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fBbtXYNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 01Ka4Ps2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWdEEc6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dGjfyzbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qrl436hQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B10WzJxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X7nDBbQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oK80JAhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFmamuN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3EvcA7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lRdwSsWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6lHhlRvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMQByXY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i43y7Yz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mw5JzHanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VeImk5wqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SySqPBKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxMg3iBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7duhhvWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55qhkqtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TM4zSEaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mtuKvYyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AuU1Q1FdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lpVt3kAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0C5wMzmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSRXffWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IiibgV17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6Mp6r33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0HevTg4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DiSOYYQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LilPv1xXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W24Q7U8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToQOuBmZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZ4UMMfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUXye1PBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjuJd8qZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4dgWDSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNXxi6IEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cudGdIUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulCGWGPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K4PwYqmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWcLbNpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g93se6Q2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHXVd5uYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VIUAZU3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkCMLEelWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jsB3N7y9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GtiVAvFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mn4Muw79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tw6mYZQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tpmi7HgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x0THtEJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMZAjaL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R28HZTz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3ggwmfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7O6HrQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enIyDFuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6udVAtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ohuQXDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdPeeO4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hig1OKqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8beiKnrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JgdqweAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkyvYEu7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcKvyL1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l52mpeS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRHrZLhIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdnCXr5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miYFrSiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6B6xKkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fH3QEjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PU90YMBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

