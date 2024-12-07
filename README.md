# Overview
This project can be seen as a proof of concept, trying to show the functionality of SSI in combination with everyday email communication. It contains two programs:
- Janus: Janus is a mail agent written in C#. Capable of communicating via IMAP/POP with your email provider. With it, you can send and read emails using SSI security.

- Monsi: Monsi is a SSI wallet implementation written in Go using [Wails](https://wails.io/). It includes endpoints for the communication with Janus and manages encryption and decryption of mail contents.

# Setup
There are binaries for immediate testing for Windows systems here:
- .\Janus\JanusWeb\bin\Debug\net6.0
- .\monsi\build\bin

Linux and Mac will have to build there own solutions.

# Future work
- Encryption via password of DIDs (and especially its passwords) on disk.

# Remarks 
- The logos for Janus and Monsi are AI generated using DALL-E