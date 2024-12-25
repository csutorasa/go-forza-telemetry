# go-forza-telemetry

This library implements the Forza "Data Out" protocol. It is compatible with:

- Forza Motorsport 7
- Forza Motorsport (2023)
- Forza Horizon 4
- Forza Horizon 5

## Install

Go 1.18 is the minimum supported version.

```sh
go get github.com/csutorasa/go-forza-telemetry
```

## Setup

- Start the game and go to the Options/HUD
- Set Data Out to `ON`
- Set Data Out IP Address to the host of your application
- Set Data Out IP Port to the port your application

### Loopback

If you have the [Steam](https://store.steampowered.com/) version of the game, you can skip this step.

If you have the [Windows App](https://learn.microsoft.com/en-us/windows/apps/windows-app-sdk/) version of the game,
and the host is the same as where the game runs, then you need to allow the local UDP packets to be sent.
You can use [Windows 8 AppContainer Loopback Utility from Telerik](https://www.telerik.com/fiddler/add-ons) which automates the process.

## Example

Examples can be found in the [example](example) directory.
