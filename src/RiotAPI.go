//Package riotapi is wrapper for RiotAPI
package riotapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"errors"
)

var httpclient http.Client
var client *Client
var endPoint string

//Client is for RiotAPI
type Client struct {
	EndPoint string
	Key      string
}

//SummonerDTO represents a summoner
type SummonerDTO struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	Puuid         string `json:"puuid"`
	SummonerLevel int    `json:"summonerLevel"`
	AccountID     string `json:"accountId"`
	ID            string `json:"id"`
	RevisionDate  int64  `json:"revisionDate"`
}

//CurrentGameInfo represents active-game's information
type CurrentGameInfo struct {
	GameID            int    `json:"gameId"`
	GameStartTime     int64  `json:"gameStartTime"`
	PlatformID        string `json:"platformId"`
	GameMode          string `json:"gameMode"`
	MapID             int    `json:"mapId"`
	GameType          string `json:"gameType"`
	GameQueueConfigID int    `json:"gameQueueConfigId"`
	Observers         struct {
		EncryptionKey string `json:"encryptionKey"`
	} `json:"observers"`
	Participants []struct {
		ProfileIconID            int           `json:"profileIconId"`
		ChampionID               int           `json:"championId"`
		SummonerName             string        `json:"summonerName"`
		GameCustomizationObjects []interface{} `json:"gameCustomizationObjects"`
		Bot                      bool          `json:"bot"`
		Perks                    struct {
			PerkStyle    int   `json:"perkStyle"`
			PerkIds      []int `json:"perkIds"`
			PerkSubStyle int   `json:"perkSubStyle"`
		} `json:"perks"`
		Spell2ID   int    `json:"spell2Id"`
		TeamID     int    `json:"teamId"`
		Spell1ID   int    `json:"spell1Id"`
		SummonerID string `json:"summonerId"`
	} `json:"participants"`
	GameLength      int `json:"gameLength"`
	BannedChampions []struct {
		TeamID     int `json:"teamId"`
		ChampionID int `json:"championId"`
		PickTurn   int `json:"pickTurn"`
	} `json:"bannedChampions"`
}

//MatchDto represents match information include teams, participants
type MatchDto struct {
	SeasonID              int `json:"seasonId"`
	QueueID               int `json:"queueId"`
	GameID                int `json:"gameId"`
	ParticipantIdentities []struct {
		Player struct {
			CurrentPlatformID string `json:"currentPlatformId"`
			SummonerName      string `json:"summonerName"`
			MatchHistoryURI   string `json:"matchHistoryUri"`
			PlatformID        string `json:"platformId"`
			CurrentAccountID  string `json:"currentAccountId"`
			ProfileIcon       int    `json:"profileIcon"`
			SummonerID        string `json:"summonerId"`
			AccountID         string `json:"accountId"`
		} `json:"player"`
		ParticipantID int `json:"participantId"`
	} `json:"participantIdentities"`
	GameVersion string `json:"gameVersion"`
	PlatformID  string `json:"platformId"`
	GameMode    string `json:"gameMode"`
	MapID       int    `json:"mapId"`
	GameType    string `json:"gameType"`
	Teams       []struct {
		FirstDragon          bool          `json:"firstDragon"`
		Bans                 []interface{} `json:"bans"`
		FirstInhibitor       bool          `json:"firstInhibitor"`
		Win                  string        `json:"win"`
		FirstRiftHerald      bool          `json:"firstRiftHerald"`
		FirstBaron           bool          `json:"firstBaron"`
		BaronKills           int           `json:"baronKills"`
		RiftHeraldKills      int           `json:"riftHeraldKills"`
		FirstBlood           bool          `json:"firstBlood"`
		TeamID               int           `json:"teamId"`
		FirstTower           bool          `json:"firstTower"`
		VilemawKills         int           `json:"vilemawKills"`
		InhibitorKills       int           `json:"inhibitorKills"`
		TowerKills           int           `json:"towerKills"`
		DominionVictoryScore int           `json:"dominionVictoryScore"`
		DragonKills          int           `json:"dragonKills"`
	} `json:"teams"`
	Participants []struct {
		Spell1ID      int `json:"spell1Id"`
		ParticipantID int `json:"participantId"`
		Timeline      struct {
			Lane             string `json:"lane"`
			ParticipantID    int    `json:"participantId"`
			GoldPerMinDeltas struct {
				Two030 float64 `json:"20-30"`
				Zero10 float64 `json:"0-10"`
				One020 float64 `json:"10-20"`
			} `json:"goldPerMinDeltas"`
			CreepsPerMinDeltas struct {
				Two030 int     `json:"20-30"`
				Zero10 int     `json:"0-10"`
				One020 float64 `json:"10-20"`
			} `json:"creepsPerMinDeltas"`
			XpPerMinDeltas struct {
				Two030 float64 `json:"20-30"`
				Zero10 float64 `json:"0-10"`
				One020 float64 `json:"10-20"`
			} `json:"xpPerMinDeltas"`
			Role                    string `json:"role"`
			DamageTakenPerMinDeltas struct {
				Two030 float64 `json:"20-30"`
				Zero10 float64 `json:"0-10"`
				One020 int     `json:"10-20"`
			} `json:"damageTakenPerMinDeltas"`
		} `json:"timeline,omitempty"`
		Spell2ID int `json:"spell2Id"`
		TeamID   int `json:"teamId"`
		Stats    struct {
			NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
			VisionScore                     int  `json:"visionScore"`
			MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
			LargestMultiKill                int  `json:"largestMultiKill"`
			TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
			LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
			Perk1Var1                       int  `json:"perk1Var1"`
			Perk1Var3                       int  `json:"perk1Var3"`
			Perk1Var2                       int  `json:"perk1Var2"`
			TripleKills                     int  `json:"tripleKills"`
			Perk5                           int  `json:"perk5"`
			Perk4                           int  `json:"perk4"`
			PlayerScore9                    int  `json:"playerScore9"`
			PlayerScore8                    int  `json:"playerScore8"`
			Kills                           int  `json:"kills"`
			PlayerScore1                    int  `json:"playerScore1"`
			PlayerScore0                    int  `json:"playerScore0"`
			PlayerScore3                    int  `json:"playerScore3"`
			PlayerScore2                    int  `json:"playerScore2"`
			PlayerScore5                    int  `json:"playerScore5"`
			PlayerScore4                    int  `json:"playerScore4"`
			PlayerScore7                    int  `json:"playerScore7"`
			PlayerScore6                    int  `json:"playerScore6"`
			Perk5Var1                       int  `json:"perk5Var1"`
			Perk5Var3                       int  `json:"perk5Var3"`
			Perk5Var2                       int  `json:"perk5Var2"`
			TotalScoreRank                  int  `json:"totalScoreRank"`
			NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
			StatPerk1                       int  `json:"statPerk1"`
			StatPerk0                       int  `json:"statPerk0"`
			DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
			PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
			DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
			Perk2Var2                       int  `json:"perk2Var2"`
			Perk2Var3                       int  `json:"perk2Var3"`
			TotalUnitsHealed                int  `json:"totalUnitsHealed"`
			Perk2Var1                       int  `json:"perk2Var1"`
			Perk4Var1                       int  `json:"perk4Var1"`
			TotalDamageTaken                int  `json:"totalDamageTaken"`
			Perk4Var3                       int  `json:"perk4Var3"`
			WardsKilled                     int  `json:"wardsKilled"`
			LargestCriticalStrike           int  `json:"largestCriticalStrike"`
			LargestKillingSpree             int  `json:"largestKillingSpree"`
			QuadraKills                     int  `json:"quadraKills"`
			MagicDamageDealt                int  `json:"magicDamageDealt"`
			FirstBloodAssist                bool `json:"firstBloodAssist"`
			Item2                           int  `json:"item2"`
			Item3                           int  `json:"item3"`
			Item0                           int  `json:"item0"`
			Item1                           int  `json:"item1"`
			Item6                           int  `json:"item6"`
			Item4                           int  `json:"item4"`
			Item5                           int  `json:"item5"`
			Perk1                           int  `json:"perk1"`
			Perk0                           int  `json:"perk0"`
			Perk3                           int  `json:"perk3"`
			Perk2                           int  `json:"perk2"`
			Perk3Var3                       int  `json:"perk3Var3"`
			Perk3Var2                       int  `json:"perk3Var2"`
			Perk3Var1                       int  `json:"perk3Var1"`
			DamageSelfMitigated             int  `json:"damageSelfMitigated"`
			MagicalDamageTaken              int  `json:"magicalDamageTaken"`
			Perk0Var2                       int  `json:"perk0Var2"`
			FirstInhibitorKill              bool `json:"firstInhibitorKill"`
			TrueDamageTaken                 int  `json:"trueDamageTaken"`
			Assists                         int  `json:"assists"`
			Perk4Var2                       int  `json:"perk4Var2"`
			GoldSpent                       int  `json:"goldSpent"`
			TrueDamageDealt                 int  `json:"trueDamageDealt"`
			ParticipantID                   int  `json:"participantId"`
			PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
			SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
			TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
			PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
			TotalPlayerScore                int  `json:"totalPlayerScore"`
			Win                             bool `json:"win"`
			ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
			TotalDamageDealt                int  `json:"totalDamageDealt"`
			NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
			Deaths                          int  `json:"deaths"`
			WardsPlaced                     int  `json:"wardsPlaced"`
			PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
			PerkSubStyle                    int  `json:"perkSubStyle"`
			TurretKills                     int  `json:"turretKills"`
			FirstBloodKill                  bool `json:"firstBloodKill"`
			TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
			GoldEarned                      int  `json:"goldEarned"`
			KillingSprees                   int  `json:"killingSprees"`
			UnrealKills                     int  `json:"unrealKills"`
			FirstTowerAssist                bool `json:"firstTowerAssist"`
			FirstTowerKill                  bool `json:"firstTowerKill"`
			ChampLevel                      int  `json:"champLevel"`
			DoubleKills                     int  `json:"doubleKills"`
			InhibitorKills                  int  `json:"inhibitorKills"`
			FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
			Perk0Var1                       int  `json:"perk0Var1"`
			CombatPlayerScore               int  `json:"combatPlayerScore"`
			Perk0Var3                       int  `json:"perk0Var3"`
			VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
			PentaKills                      int  `json:"pentaKills"`
			TotalHeal                       int  `json:"totalHeal"`
			TotalMinionsKilled              int  `json:"totalMinionsKilled"`
			TimeCCingOthers                 int  `json:"timeCCingOthers"`
			StatPerk2                       int  `json:"statPerk2"`
		} `json:"stats"`
		ChampionID                int    `json:"championId"`
		HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier,omitempty"`
	} `json:"participants"`
	GameDuration int   `json:"gameDuration"`
	GameCreation int64 `json:"gameCreation"`
}

//MatchTimelineDto represents matchTimeline per min
type MatchTimelineDto struct {
	Frames []struct {
		Timestamp         int `json:"timestamp"`
		ParticipantFrames struct {
			P1 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"1"`
			P2 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"2"`
			P3 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"3"`
			P4 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"4"`
			P5 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"5"`
			P6 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"6"`
			P7 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"7"`
			P8 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"8"`
			P9 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"9"`
			P10 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"10"`
		} `json:"participantFrames"`
		Events []interface{} `json:"events"`
	} `json:"frames"`
	FrameInterval int `json:"frameInterval"`
}

//Print is a fucntion shows ClientKEY
func (*Client) Print() {
	fmt.Print(client.Key)
}

//New is constructor
func New(key string) *Client {
	client = &Client{EndPoint: "https://jp1.api.riotgames.com/", Key: key}
	return client
}

//GetSummonersByName is a function return SummonerDTO.
func (me *Client) GetSummonersByName(SN string) (SummonerDTO, error) {
	data, networkError := getRequest(me.EndPoint+"/lol/summoner/v4/summoners/by-name/", me.Key, SN)
	var summoner SummonerDTO
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &summoner); decodeError != nil {
			return summoner, nil
		} else {
			return summoner, decodeError
		}
	} else {
		return summoner, networkError
	}
}

//GetActivegamesBySummoner is a function return CurrentGameInfo.
func (me *Client) GetActivegamesBySummoner(encryptedSummonerID string) (CurrentGameInfo, error) {
	data, networkError := getRequest(me.EndPoint + "/lol/spectator/v4/active-games/by-summoner/", me.Key, encryptedSummonerID)
	var currentGame CurrentGameInfo
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &currentGame); decodeError != nil {
			return currentGame, nil
		} else {
			return currentGame, decodeError
		}
	} else {
		return currentGame, networkError
	}
}

//GetMatches is a function return MatchDto.
func (me *Client) GetMatches(gameID int) (MatchDto, error) {
	data, networkError := getRequest(me.EndPoint + "/lol/match/v4/matches/", me.Key, strconv.Itoa(gameID))
	var matchDto MatchDto
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &matchDto); decodeError != nil {
			return matchDto, nil
		} else {
			return matchDto, decodeError
		}
	} else {
		return matchDto, networkError
	}
}

//GetTimelineByMatch is a function return MatchDto.
func (me *Client) GetTimelineByMatch(gameID int) (MatchTimelineDto, error) {
	data, networkError := getRequest(me.EndPoint + "/lol/match/v4/matches/", me.Key, strconv.Itoa(gameID))
	var timeline MatchTimelineDto
	if networkError == nil {
		if decodeError := json.Unmarshal(data, &timeline); decodeError != nil {
			return timeline, nil
		} else {
			return timeline, decodeError
		}
	} else {
		return timeline, networkError
	}
}

func getRequest(uri, key, param string) ([]byte, error) {
	req, _ := http.NewRequest("GET", uri+url.PathEscape(param), nil)
	req.Header.Set("X-Riot-Token", key)
	data, err := httpclient.Do(req)
	if err == nil {
		buf := new(bytes.Buffer)
		io.Copy(buf, data.Body)
		if data.StatusCode != 200 {
			return buf.Bytes(), errors.New(string(data.StatusCode) + data.Status)
		}
		return buf.Bytes(), err
	} else {
		return nil, err
	}
}