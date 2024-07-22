# The First Descendant Go Client

## Overview

`go-tfd-client` is a Go client library for interacting with "The First Descendant" game API. This client wraps all the API endpoints listed in the [NEXON Open API](https://openapi.nexon.com/game/tfd/?id=21) documentation, providing an easy-to-use interface for Go developers to integrate with the game's services.

## Features

- Full coverage of all API endpoints.
- Easy-to-use and idiomatic Go interface.
- Support for authentication and authorization.

## Installation

To install the library, use the following command:

```shell
go get github.com/DevZoneLabs/go-tfd-client
```

## Usage

Here is a basic example of how to use the client:

```go
package main

import (
    "fmt"
    "log"

    tfd "github.com/DevZoneLabs/go-tfd-client"
)

func main() {

    client := tfd.NewClient("YOUR_API_KEY", nil)

    // Example: Get AccountIdentifier
    accIdentifier, err := client.GetAccountIdentifier("PLAYER_ID")
    if err != nil {
        log.Fatal("Error fetching player profile: %v", err)
    }

    fmt.Printf("AccountIdentifier: %s\v", accIdentifier.OOUID)
}
```

## Documentation

Full documentation is available on [GoDocs](https://pkg.go.dev/github.com/DevZoneLabs/go-tfd-api)

## API Coverage

The client covers the following API endpoints:

#### Users

- GetAccountIdentifier(userName string) (\*models.AccountIdentifier, error)

- GetBasicInformation(ouid string) (\*models.UserBasic, error)

- GetUserDescendant(ouid string) (\*models.UserDescendant, error)

- GetUserWeapons(ouid string, language LanguageCode) (\*models.UserWeapon, error)

- GetUserReactor(ouid string, language LanguageCode) (\*models.UserReactor, error)

- GetExternalComponent(ouid string, language LanguageCode) (\*models.UserExternalComponent, error)

#### Metadata

- GetDescendantsMetadata(language LanguageCode) ([]models.Descendant, error)

- GetWeaponsMetadata(language LanguageCode) ([]models.Weapon, error)

- GetModulesMetadata(language LanguageCode) ([]models.Module, error)

- GetReactorsMetadata(language LanguageCode) ([]models.Reactor, error)

- GetExternalComponentsMetadata(language LanguageCode) ([]models.ExternalComponent, error)

- GetRewardsMetadata(language LanguageCode) ([]models.Reward, error)

- GetStatsMetadata(language LanguageCode) ([]models.StatMeta, error)

- GetVoidBattlesMetadata(language LanguageCode) (\*models.VoidBattleResponse, error)

- GetVoidBattlesMetadata(language LanguageCode) (\*models.VoidBattleResponse, error)

- GetTitlesMetadata(language LanguageCode) (\*models.TitleResponse, error)

#### Module Recommendation

- GetModuleRecommendation(opts ModuleRecommendationOpts) (\*models.ModuleRecommendationResponse, error)
