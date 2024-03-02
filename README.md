# vdl
 An OnionOS direct download utility for the Miyoo Mini Plus. Note: this utility will only work with roms hosted on Vimm's Lair, and is an unofficial tool made by me for me. This is the first project I've made in Go, and was mostly done for fun and to learn, so please take it easy on me when reading the code, and I'd be happy to take any feedback on improvements that could be made. :)

## Installation and Usage
Grab the vdl file from the releases and store it on your SD card somewhere. I put mine in the root directory, as that's what loads up first when you enter the Terminal on the console. This gives you easy access to the download utility with minimal button presses.

Navigate to the Terminal on your Miyoo (in the Apps folder, and can be installed using the Package Manager if you don't have it already). Run the download utility by typing: 
```
./vdl
```
This assumes you've stored it in the root of your SD card. You can use the SELECT button to tab, and the START button as enter. So you can type "./v", then press SELECT to autocomplete, and then press START to run the application. 

You will then get a prompt to enter the Vault ID of the rom you wish to download, you can locate the Vault ID here (https://vimm.net/):

![Screenshot showing how to identify the vault ID of your rom from the url of the rom hosting site.](https://imgkk.com/i/1iqq.png)

In this case the ID is "3454". Type this into the downloader using the onscreen keyboard, and then press enter, or START. The rom will download, identify where it should be stored, and put it in the correct folder for you.

## Console Compatibility
| Console  | Compatible | Notes  |
| ------------- | ------------- | ------------- |
| Atari 2600  | ✔️  | -  |
| Atari 5600  | ✔️  | -  |
| Nintendo  | ✔️  | -  |
| Master System  | ✔️  | -  |
| Atari 7800  | ✔️  | -  |
| Gensis  | ✔️  | -  |
| Super Nintendo  | ✔️  | -  |
| Sega 32X  | ✔️  | -  |
| Playstation  | ❌  | Currently files are downloaded as .7z, need functionality to unzip  |
| Game Boy  | ✔️  | -  |
| Lynx  | ✔️  | -  |
| Game Gear  | ✔️  | -  |
| Virtual Boy  | ✔️  | -  |
| Game Boy Colour  | ✔️  | -  |
| Game Boy Advance  | ✔️  | -  |
| Nintendo DS  | ✔️  | -  |

## Building From Source
Install Go, clone the repo, modify your environment variables to compile for ARM and Linux, this is how I've done it in Powershell:
```
$Env:GOOS = "linux"; $Env:GOARCH = "arm"
```
In the folder you cloned run:
```
go build
```
It should spit you out an executable which you can then transfer to your Miyoo and run as above. 

## Planned Features
- PS1 compatibility
- Download progress bar
- Better error handling around failed downloads
- Rom browser to make solution self-contained
- Support for multi-disc downloads
