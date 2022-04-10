# Coins-Info

[![Build Status](https://app.travis-ci.com/samettunay/coins-info.svg?branch=main)](https://travis-ci.com/samettunay/coins-info)
[![GoDoc](https://godoc.org/github.com/anaskhan96/soup?status.svg)](https://pkg.go.dev/github.com/anaskhan96/soup)
[![Go Report Card](https://goreportcard.com/badge/github.com/samettunay/coins-info)](https://goreportcard.com/report/github.com/samettunay/coins-info)

Crypto coin live information with [Golang](https://www.golang.org/)


```
 go get github.com/samettunay/coins-info
```

### Functions

```go

func ShowLiveTable(int) // Displays cryptocurrencies in live tabular form

func Info(string) // Shows the information of the cryptocurrency you are looking for

```

### Output

```go
coins.ShowLiveTable(10)
```

![Screenshot_1](https://user-images.githubusercontent.com/79511355/162594936-dfb17f6b-3650-493c-808f-407b7a3dad8c.png)

```go
coins.Info("BTC")
```

![Screenshot_4](https://user-images.githubusercontent.com/79511355/162624537-f2f44f79-c658-4a37-8df4-042d6becafc4.png)

