package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/cespare/xxhash"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
    username := usernames[rand.Intn(len(usernames))]
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://192.168.1.25:1883")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("chat", 0, func(_ mqtt.Client, m mqtt.Message) {
        var message Message
        if err := json.Unmarshal(m.Payload(), &message); err != nil {
            fmt.Println("Error unmarshalling JSON:", err)
            return
        }

        if username != message.User {
            fmt.Printf("\r%s%s: %s%s\n", getColorForUser(message.User), message.User, message.Text, colorReset)
            fmt.Print(">> ")
        }
    }); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print(">> ")
        scanner.Scan() // Read a line of text from the user

        if err := scanner.Err(); err != nil {
            fmt.Println("Error reading input:", err)
            break
        }

        userInput := scanner.Text()

        // If the user typed "quit", then exit the program
        if userInput == "quit" {
            break
        }

        // If the user just pressed enter, do nothing
        if strings.TrimSpace(userInput) == "" {
            fmt.Println("\033[A\033[A")
            continue
        }

        payload, err := json.Marshal(Message{Text: userInput, User: username})
        if err != nil {
            fmt.Println("Error marshalling JSON:", err)
            continue
        }

        c.Publish("chat", 0, false, payload)

        fmt.Printf("\033[FYou: %s\n", userInput)
    }

	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}

type Message struct {
    Text string `json:"text"`
    User string `json:"user"`
}

func getColorForUser(username string) string {
    hash := xxhash.Sum64String(username)
    for i := 0; i < len(colors); i++ {
        if colorsInUse[int(hash % uint64(len(colors)))] == "" {
            colorsInUse[int(hash % uint64(len(colors)))] = username
            break
        }

        if colorsInUse[int(hash % uint64(len(colors)))] == username {
            break
        }

        hash++
    }
    return colors[int(hash % uint64(len(colors)))]
}

var colorsInUse = make(map[int]string)

var usernames = []string{
        "CyberX",
        "NanoBot",
        "Bitstream",
        "Cipher",
        "ByteFreak",
        "NeonHacker",
        "DataPirate",
        "TechnoWiz",
        "CodeNinja",
        "ZeroCool",
        "BinaryGuru",
        "BitHacker",
        "CryptoKid",
        "PixelPirate",
        "SynthWave",
        "CyberPunk",
        "Hacktivist",
        "NeonGhost",
        "VirusViper",
        "FireWall",
        "DarkNet",
        "StealthByte",
        "RoboHacker",
        "MatrixMind",
        "CyberSphinx",
        "CircuitBreaker",
        "NanoByte",
        "InfoBorg",
        "WebWeaver",
        "GhostCode",
        "BitBender",
        "SiliconShadow",
        "HoloByte",
        "RetroCyber",
        "QuantumByte",
        "MetaHacker",
        "Enigma",
        "NetNinja",
        "HyperLink",
        "CyberShade",
        "TechNomad",
        "DeepDive",
        "CyberVortex",
        "NeonCode",
        "CyberGlider",
        "VoxelViper",
        "DataSeeker",
        "FireWire",
        "StealthBlade",
        "OmegaByte",
        "BitSlinger",
        "WebSorcerer",
        "BinaryDragon",
        "VirusHunter",
        "SynthRider",
        "CipherMaster",
        "InfoPioneer",
        "NeuroCipher",
        "CryoByte",
        "PhotonPirate",
        "CodeMancer",
        "BitWizard",
        "HoloHacker",
        "CyberPhantom",
        "StealthSmith",
        "QuantumGhost",
        "NeonChaser",
        "TechnoViking",
        "DarkProphet",
        "HackWhisperer",
        "DigiNaut",
        "MatrixMystic",
        "CyberNomad",
        "BitWanderer",
        "NeonStrider",
        "ByteRaider",
        "CryptoCyborg",
        "NanoScribe",
        "FireCipher",
        "WebSlinger",
        "ByteVoyager",
        "DataPhreak",
        "NeonTrail",
        "CircuitSorcerer",
        "SynthRogue",
        "CodeFusionist",
        "HoloNomad",
        "CyberShadow",
        "QuantumCrafter",
        "BitPioneer",
        "InfoSculptor",
        "NeuroRider",
        "VirusVoyager",
        "StealthGlider",
        "OmegaCode",
        "PixelPhantom",
        "TechRogue",
    }

var colors = []string{
    "\033[31m",
    "\033[32m",
    "\033[33m",
    "\033[34m",
    "\033[35m",
    "\033[36m",
}

var colorReset  = "\033[0m"
