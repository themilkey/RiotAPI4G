# RiotAPI4G
RiotAPI wrapper for golang

## Installation 
In your shell  
`$go get -u github.com/themilkey/RiotAPI4G/src` 

## How to import
`import riotapi "github.com/themilkey/RiotAPI4G/src"`    

## How to use
```go
client := riotapi.New("YOUR_API_KEY")
summoner, err := client.GetSummonersByName("summonerName")
if err == nil {
  Println(summoner.ID)
} else {
  Println(err)
}
```
